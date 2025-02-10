package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/private-project-pp/pos-rest-api-service/interfaces"
	"github.com/private-project-pp/pos-rest-api-service/shared/config"
)

func StartServer() {
	server := echo.New()

	err := interfaces.Container(server)
	if err != nil {
		panic(err)
	}

	for _, route := range server.Routes() {
		fmt.Printf("[%s] - %s - %s", route.Method, route.Name, route.Path)
		fmt.Println()
	}

	server.Start(config.Service.Port)
}
