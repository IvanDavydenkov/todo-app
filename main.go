package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	rest_server "todo-app/infrastructure/rest-server"
)

func main() {
	port := ":8080"
	server := rest_server.NewGinServer()
	server.RegisterPublicRoute(http.MethodGet, "/api/ping", func(w http.ResponseWriter, r *http.Request) {
		ResponseByJson(w, "pong", http.StatusOK)
	})
	if err := server.Start(port); err != nil {
		panic("Ошибка запуска сервера:" + err.Error())
	}
	fmt.Println("Сервер запущен на порте:", port)
}

var contentType = "Content-Type"
var applicationJson = "application/json"

func ResponseByJson(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
