package route

import (
	"github.com/gin-gonic/gin"
	"github.com/private-project-pp/pos-rest-api-service/handler"
)

type routeSetup struct {
	gin             *gin.Engine
	internalhandler *handler.InternalHandler
}

func SetupRoute(route *gin.Engine, internal *handler.InternalHandler) {

	routes := routeSetup{
		gin:             route,
		internalhandler: internal,
	}

	routes.InternalRoute()

}
