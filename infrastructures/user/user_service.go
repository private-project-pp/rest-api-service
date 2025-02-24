package user

import (
	"github.com/labstack/echo/v4"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

type UserInfra interface {
	UserAdd(inCtx echo.Context, in *model.UserAddRequestPayload) (out *model.UserAddResponse, err error)
}
