package users_administration

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication/response"
	"github.com/private-project-pp/pos-rest-api-service/usecase/users_administration/request"
)

type UsersAdministration interface {
	AddUser(ctx echo.Context, req request.UserAddRequest) (out response.UserAddResponse, err error)
}
