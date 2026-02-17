package dto

type CreateTodoRequestDTO struct {
	Title    string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle" binding:"required"`
}
