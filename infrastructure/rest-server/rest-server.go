package rest_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
	server *http.Server
}

func (s *GinServer) RegisterPublicRoute(method, path string, handler gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.router.GET(path, handler)
	case http.MethodPost:
		s.router.POST(path, handler)
	case http.MethodPut:
		s.router.PUT(path, handler)
	case http.MethodPatch:
		s.router.PATCH(path, handler)
	case http.MethodDelete:
		s.router.DELETE(path, handler)
	default:
		s.router.Any(path, handler)
	}
}

func (s *GinServer) Start(port string) error {
	s.server = &http.Server{
		Addr:    port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
