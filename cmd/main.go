package main

import (
	"log"

	"bff-todolist/internal/delivery/http"
	"bff-todolist/internal/repository"
	"bff-todolist/internal/usecase"
	"bff-todolist/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	repo := repository.NewTodoRepository()
	uc := usecase.NewTodoUsecase(repo)
	handler := http.NewTodoHandler(uc)

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
