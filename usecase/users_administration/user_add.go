package users_administration

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication/response"
	"github.com/private-project-pp/pos-rest-api-service/usecase/users_administration/request"
)

func (s userAdm) AddUser(ctx echo.Context, req request.UserAddRequest) (out response.UserAddResponse, err error) {

	payload := &model.UserAddRequestPayload{}

	result, err := s.userInfra.UserAdd(ctx, payload)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = response.UserAddResponse{
		UserId:  result.GetUserId(),
		Message: result.GetMessage(),
		Status:  result.GetStatus(),
	}

	return out, nil
}
