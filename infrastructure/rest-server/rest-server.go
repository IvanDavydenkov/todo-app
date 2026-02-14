package rest_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
	server *http.Server
}

func (s *GinServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {
	ginHandler := func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}
	switch method {
	case http.MethodGet:
		s.router.GET(path, ginHandler)
	case http.MethodPost:
		s.router.POST(path, ginHandler)
	case http.MethodPut:
		s.router.PUT(path, ginHandler)
	case http.MethodPatch:
		s.router.PATCH(path, ginHandler)
	case http.MethodDelete:
		s.router.DELETE(path, ginHandler)
	default:
		s.router.Any(path, ginHandler)
	}
}

func (s *GinServer) Start(port string) error {
	s.server = &http.Server{
		Addr:    port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
