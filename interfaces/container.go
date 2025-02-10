package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-rest-api-service/handler"
	"github.com/private-project-pp/pos-rest-api-service/infrastructures/user"
	"github.com/private-project-pp/pos-rest-api-service/interfaces/route"
	"github.com/private-project-pp/pos-rest-api-service/shared/config"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Container(driver *echo.Echo) (err error) {
	// Init config, rpc client connection setup, handler setup, router setup
	config, err := config.SetupConfig()
	if err != nil {
		return err
	}

	_ = config

	userConn, err := grpc.NewClient(config.Internal.UserRpcService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	userInfra := user.SetupUser(userConn)

	authUseCase := authentication.SetupAuthentication(userInfra)

	internalHandler := handler.SetupInternalHandler(authUseCase)

	route.SetupRoute(driver, internalHandler)

	return nil
}
