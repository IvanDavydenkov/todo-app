package rest_server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	RegisterPublicRoute(method, path string, handler gin.HandlerFunc)
	Start(port string) error
}
