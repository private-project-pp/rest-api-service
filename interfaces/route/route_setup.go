package route

import (
	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-rest-api-service/handler"
)

type routeSetup struct {
	echo            *echo.Echo
	internalhandler *handler.InternalHandler
}

func SetupRoute(driver *echo.Echo, internal *handler.InternalHandler) {

	routes := routeSetup{
		echo:            driver,
		internalhandler: internal,
	}
	routes.InternalRoute()

}
