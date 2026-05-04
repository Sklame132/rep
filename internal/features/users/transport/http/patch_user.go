package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Sklame132/rep/internal/core/domain"
	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_request "github.com/Sklame132/rep/internal/core/transport/http/request"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
	core_http_types "github.com/Sklame132/rep/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	Username    core_http_types.Nullable[string] `json:"username"`
	Password    core_http_types.Nullable[string] `json:"password"`
	FirstName   core_http_types.Nullable[string] `json:"first_name"`
	LastName    core_http_types.Nullable[string] `json:"last_name"`
	Address     core_http_types.Nullable[string] `json:"address"`
	Email       core_http_types.Nullable[string] `json:"email"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
	Rating      core_http_types.Nullable[int16]  `json:"rating"`
	Role        core_http_types.Nullable[string] `json:"role"`
}

func  ValidateFieldBetween(l int, min, max int, fieldName string) error {
	if l < min || l > max {
		return fmt.Errorf("`%s` must be between %v and %v symbols", fieldName, min, max)
	}

	return nil
}

func ValidateNotNull(s *string, fieldName string) error {
	if s == nil {
		return fmt.Errorf("`%s` cat't be NULL", fieldName)
	}

	return nil
}

func (r *PatchUserRequest) Validate() error {
	if r.Username.Set {
		if notNullErr := ValidateNotNull(r.Username.Value, "Username"); notNullErr != nil {
			return notNullErr
		}
		usernameLen := len([]rune(*r.Username.Value))
		if betweenErr := ValidateFieldBetween(usernameLen, 3, 16, "Username"); betweenErr != nil {
			return betweenErr
		}
	}
	if r.Password.Set {
		if notNullErr := ValidateNotNull(r.Password.Value, "Password"); notNullErr != nil {
			return notNullErr
		}
		passwordLen := len([]rune(*r.Password.Value))
		if betweenErr := ValidateFieldBetween(passwordLen, 8, 64, "Password"); betweenErr != nil {
			return betweenErr
		}
	}
	if r.FirstName.Set {
		if notNullErr := ValidateNotNull(r.FirstName.Value, "FirstName"); notNullErr != nil {
			return notNullErr
		}
		firstNameLen := len([]rune(*r.FirstName.Value))
		if betweenErr := ValidateFieldBetween(firstNameLen, 0, 32, "FirstName"); betweenErr != nil {
			return betweenErr
		}
	}
	if r.LastName.Set {
		if notNullErr := ValidateNotNull(r.LastName.Value, "LastName"); notNullErr != nil {
			return notNullErr
		}
		lastNameLen := len([]rune(*r.LastName.Value))
		if betweenErr := ValidateFieldBetween(lastNameLen, 0, 32, "LastName"); betweenErr != nil {
			return betweenErr
		}
	}
	if r.Email.Set {
		if r.Email.Value != nil {
			emailLen := len([]rune(*r.Email.Value))
			if betweenErr := ValidateFieldBetween(emailLen, 0, 32, "Email"); betweenErr != nil {
				return betweenErr
			}
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if betweenErr := ValidateFieldBetween(phoneNumberLen, 0, 32, "PhoneNumber"); betweenErr != nil {
				return betweenErr
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must startswith `+` symbol")
			}
		}
	}
	if r.Role.Set {
		if r.Role.Value != nil {
			roleLen := len([]rune(*r.Role.Value))
			if betweenErr := ValidateFieldBetween(roleLen, 0, 16, "Role"); betweenErr != nil {
				return betweenErr
			}
		}
	}
	
	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userLogin, err := core_http_request.GetStringPathValue(r, "login")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userLogin path value",
		)

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidationRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}


	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userLogin, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.Username.ToDomain(),
		request.Password.ToDomain(),
		request.FirstName.ToDomain(),
		request.LastName.ToDomain(),
		request.Address.ToDomain(),
		request.Email.ToDomain(),
		request.PhoneNumber.ToDomain(),
		request.Rating.ToDomain(),
		request.Role.ToDomain(),
	)
}
