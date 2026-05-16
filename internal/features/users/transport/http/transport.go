package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Sklame132/rep/internal/core/domain"
	core_http_server "github.com/Sklame132/rep/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		login string,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		login string,
	) error

	PatchUser(
		ctx context.Context,
		login string,
		patch domain.UserPatch,
	) (domain.User, error)
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodGet,
			Path: "/users",
			Handler: h.GetUsers,
		},
		{
			Method: http.MethodGet,
			Path: "/users/{login}",
			Handler: h.GetUser,
		},
		{
			Method: http.MethodDelete,
			Path: "/users/{login}",
			Handler: h.DeleteUser,
		},
		{
			Method: http.MethodPatch,
			Path: "/users/{login}",
			Handler: h.PatchUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: h.CreateUser,
		},
		{
			Method: http.MethodPost,
			Path: "/login",
			Handler: h.Login,
		},
		{
			Method: http.MethodGet,
			Path: "/user",
			Handler: h.User,
		},
		{
			Method: http.MethodPost,
			Path: "/logout",
			Handler: h.Logout,
		},
	}
}
