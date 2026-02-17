package dto

type GetTodoByIdRequestDTO struct {
	Id string `json:"id" binding:"required"`
}
