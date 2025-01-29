package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/private-project-pp/pos-rest-api-service/handler"
	"github.com/private-project-pp/pos-rest-api-service/interfaces/route"
	"github.com/private-project-pp/pos-rest-api-service/usecase/authentication"
)

func Container(gin *gin.Engine) (err error) {
	// Init config, rpc client connection setup, handler setup, router setup

	authUseCase := authentication.SetupAuthentication()

	internalHandler := handler.SetupInternalHandler(authUseCase)

	route.SetupRoute(gin, internalHandler)

	return nil
}
