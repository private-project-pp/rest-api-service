package cmd

import (
	"github.com/private-project-pp/pos-general-lib/infrastructure"
	"github.com/private-project-pp/pos-rest-api-service/interfaces"
	"github.com/private-project-pp/pos-rest-api-service/shared/config"
)

func StartServer() {
	server := infrastructure.SetupRestServer()

	err := interfaces.Container(server)
	if err != nil {
		panic(err)
	}

	server.StartServer(config.Service.Port)
}
