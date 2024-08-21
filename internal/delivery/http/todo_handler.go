package http

import (
	"net/http"

	"bff-todolist/internal/domain"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	usecase domain.TodoUsecase
}

func NewTodoHandler(useCase domain.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: useCase}
}

func (h *TodoHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/todos", h.GetAll)
	r.POST("/todos", h.Create)
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
