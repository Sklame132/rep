package users_transport_http

import (
	"net/http"
	"time"

	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_request "github.com/Sklame132/rep/internal/core/transport/http/request"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const secretKey = "secret"

func (h *UsersHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	const (
		loginQueryParamKey    = "login"
		passwordQueryParamKey = "password"
	)

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request AuthUserRequest
	if err := core_http_request.DecodeAndValidationRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	user, err := h.usersService.GetUser(ctx, request.Username)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		responseHandler.ErrorResponse(err,
			"incorrect password",
		)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Username,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		responseHandler.ErrorResponse(err, "could not login")
		return
	}

	cookie := http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	responseHandler.JSONResponse(map[string]string{
		"message": "success",
	}, http.StatusOK)
}

func (h *UsersHTTPHandler) User(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w);
	
	cookie, err := r.Cookie("access_token")
	if err != nil {
		responseHandler.ErrorResponse(err, "cookies not found")
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		responseHandler.JSONResponse(map[string]string{
			"message": "unauthenticated",
		}, http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := h.usersService.GetUser(ctx, claims.Issuer)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}
	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func (h *UsersHTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w);
	
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	responseHandler.JSONResponse(map[string]string{
		"message": "success",
	}, http.StatusOK)
}