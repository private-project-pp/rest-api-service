package interfaces

import (
	"github.com/private-project-pp/pos-general-lib/infrastructure"
	"github.com/private-project-pp/pos-rest-api-service/handler"
	"github.com/private-project-pp/pos-rest-api-service/infrastructures/user"
	"github.com/private-project-pp/pos-rest-api-service/interfaces/route"
	"github.com/private-project-pp/pos-rest-api-service/shared/config"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication"
)

func Container(driver infrastructure.RestServer) (err error) {
	// Init config, rpc client connection setup, handler setup, router setup
	config, err := config.SetupConfig()
	if err != nil {
		return err
	}

	_ = config

	userConn, err := infrastructure.InitRpcClientConnection(config.Internal.UserRpcService.Address)
	if err != nil {
		return err
	}

	userInfra := user.SetupUser(userConn)

	authUseCase := authentication.SetupAuthentication(userInfra)

	internalHandler := handler.SetupInternalHandler(authUseCase)

	route.SetupRoute(driver.RoutesDeclarations(), internalHandler)

	return nil
}
