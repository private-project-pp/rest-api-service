package interfaces

import (
	"github.com/private-project-pp/pos-general-lib/infrastructure"
	"github.com/private-project-pp/pos-general-lib/logger"
	"github.com/private-project-pp/pos-general-lib/middleware"
	"github.com/private-project-pp/pos-rest-api-service/handler"
	"github.com/private-project-pp/pos-rest-api-service/infrastructures/user"
	"github.com/private-project-pp/pos-rest-api-service/interfaces/route"
	"github.com/private-project-pp/pos-rest-api-service/shared/config"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication"
)

func Container() (err error) {
	// Init config, rpc client connection setup, handler setup, router setup

	log, err := logger.SetupLogger("")

	mwConn, err := infrastructure.InitRpcClientConnection(config.Internal.UserRpcService.Address)
	if err != nil {
		// return err
	}
	defer mwConn.Close()
	middleware := middleware.SetupMiddleware(mwConn)

	server := infrastructure.SetupRestServer(log, middleware)
	config, err := config.SetupConfig()
	if err != nil {
		// return err
	}

	userConn, err := infrastructure.InitRpcClientConnection(config.Internal.UserRpcService.Address)
	if err != nil {
		// return err
	}
	defer userConn.Close()

	userInfra := user.SetupUser(userConn)

	authUseCase := authentication.SetupAuthentication(userInfra)

	internalHandler := handler.SetupInternalHandler(authUseCase)

	route.SetupRoute(server.RouteInitialization(), internalHandler)

	server.StartServer(":8083")

	return nil
}
