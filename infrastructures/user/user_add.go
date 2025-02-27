package user

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-general-lib/shared/utils"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (u userInfra) UserAdd(inCtx echo.Context, in *model.UserAddRequestPayload) (out *model.UserAddResponse, err error) {
	mdContext := utils.ClearContextWithMetadata(inCtx)

	ctx, cancel := context.WithTimeout(mdContext, 2*time.Second)
	defer cancel()
	payload := &model.UserAddRequest{
		Payload: in,
	}
	u.infra.UserAdd(ctx, payload)
	return out, nil
}
