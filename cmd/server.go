package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/private-project-pp/pos-rest-api-service/interfaces"
)

func StartServer() {
	server := gin.New()

	err := interfaces.Container(server)
	if err != nil {
		panic(err)
	}

}
