package rest_server

import (
	rest_server "todo-app/infrastructure/rest-server/interface"

	"github.com/gin-gonic/gin"
)

func NewGinServer() rest_server.Server {
	router := gin.Default()

	return &GinServer{
		router: router,
	}
}
