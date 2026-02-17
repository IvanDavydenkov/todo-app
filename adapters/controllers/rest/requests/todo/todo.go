package todo

import (
	"fmt"
	"net/http"
	"time"
	"todo-app/adapters/controllers/rest/requests"
	"todo-app/boundary/domain/use-case"
	"todo-app/boundary/dto"
	"todo-app/domain/entities"

	"github.com/gin-gonic/gin"
)

var (
	PostTodo = requests.Route{
		Method: "POST",
		Path:   "/api/todo/",
	}
	GetTodo = requests.Route{
		Method: "GET",
		Path:   "/api/todo/:id",
	}
)

type TodoController struct {
	todoUseCase use_case.TodoUseCaseInterface
}

func (controller *TodoController) CreateTodo(c *gin.Context) {
	var data dto.CreateTodoRequestDTO
	err := c.ShouldBindJSON(&data)
	fmt.Println(data, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

type TodoResponseDTO struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Id          string `json:"id"`
	Status      entities.TodoStatusEnum
	CreatedAt   time.Time
	LastChanged time.Time
}

func (controller *TodoController) GetTodoById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
}
