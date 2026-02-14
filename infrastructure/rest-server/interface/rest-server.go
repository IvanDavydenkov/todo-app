package rest_server

import "net/http"

type Server interface {
	RegisterPublicRoute(method, path string, handler http.HandlerFunc)
	Start(port string) error
}
