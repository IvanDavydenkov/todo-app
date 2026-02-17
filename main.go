package main

import (
	"fmt"
	"net/http"
	"todo-app/adapters/controllers/rest/requests/todo"
	rest_server "todo-app/infrastructure/rest-server"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"
	server := rest_server.NewGinServer()
	server.RegisterPublicRoute(http.MethodGet, "/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	todo.NewTodoController(server)
	if err := server.Start(port); err != nil {
		panic("Ошибка запуска сервера:" + err.Error())
	}
	fmt.Println("Сервер запущен на порте:", port)
}
