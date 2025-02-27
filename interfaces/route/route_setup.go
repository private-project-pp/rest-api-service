package route

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-rest-api-service/handler"
)

type routeSetup struct {
	*echo.Group
	internalhandler *handler.InternalHandler
}

func SetupRoute(driver *echo.Group, internal *handler.InternalHandler) {

	routes := routeSetup{
		Group:           driver,
		internalhandler: internal,
	}
	routes.InternalRoute()

}
