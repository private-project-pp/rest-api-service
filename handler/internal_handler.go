package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication"
)

type InternalHandler struct {
	auth authentication.Authentication
}

func SetupInternalHandler(auth authentication.Authentication) *InternalHandler {
	return &InternalHandler{
		auth: auth,
	}
}

func (h InternalHandler) ValidateLoginUser(ctx echo.Context) error {
	out, err := h.auth.UserList()
	if err != nil {
		return err
	}
	return ctx.JSON(200, out)
}
