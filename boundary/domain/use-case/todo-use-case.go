package use_case

import (
	"context"
	"todo-app/boundary/dto"
)

type TodoUseCaseInterface interface {
	CreateTodo(ctx context.Context, data *dto.CreateTodoRequestDTO) (dto.CreateTodoResponseDTO, error)
	GetTodoById(ctx context.Context, data *dto.GetTodoByIdRequestDTO) (dto.GetTodoByIdResponseDTO, error)
}
