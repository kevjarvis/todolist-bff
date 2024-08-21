package main

import (
	"log"

	"bff-todolist/internal/delivery/http"
	"bff-todolist/internal/repository"
	"bff-todolist/internal/usecase"
	"bff-todolist/pkg/config"
	"bff-todolist/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	db, err := database.NewGormDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.NewGormTodoRepository(db)
	uc := usecase.NewTodoUsecase(repo)
	handler := http.NewTodoHandler(uc)

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
