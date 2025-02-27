package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-general-lib/session"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication"
	"github.com/private-project-pp/pos-rest-api-service/usecase/users_administration"
	"github.com/private-project-pp/pos-rest-api-service/usecase/users_administration/request"
)

type InternalHandler struct {
	auth    authentication.Authentication
	userAdm users_administration.UsersAdministration
}

func SetupInternalHandler(auth authentication.Authentication) *InternalHandler {
	return &InternalHandler{
		auth: auth,
	}
}

func (h InternalHandler) AddUser(ctx echo.Context) error {
	req, err := request.NewUserAdd(ctx)
	if err != nil {
		return session.SetRestResponse(ctx, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}
	out, err := h.userAdm.AddUser(ctx, req)
	if err != nil {
		return session.SetRestResponse(ctx, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}

	return session.SetRestResponse(ctx, out, nil)
}

func (h InternalHandler) ValidateLoginUser(ctx echo.Context) error {
	out, err := h.auth.UserList()
	if err != nil {
		return err
	}
	return ctx.JSON(200, out)
}
