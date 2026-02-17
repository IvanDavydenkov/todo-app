package dto

import (
	"time"
	"todo-app/domain/entities"
)

type GetTodoByIdResponseDTO struct {
	Title       string                  `json:"title"`
	Subtitle    string                  `json:"subtitle"`
	Id          string                  `json:"id"`
	Status      entities.TodoStatusEnum `json:"status"`
	CreatedAt   time.Time               `json:"createdAt"`
	LastChanged time.Time               `json:"lastChanged"`
}
