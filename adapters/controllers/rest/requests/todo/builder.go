package todo

import rest_server "todo-app/infrastructure/rest-server/interface"

func NewTodoController(server rest_server.Server) {
	controller := &TodoController{}
	server.RegisterPublicRoute(PostTodo.Method, PostTodo.Path, controller.CreateTodo)
	server.RegisterPublicRoute(GetTodo.Method, GetTodo.Path, controller.GetTodoById)
}
