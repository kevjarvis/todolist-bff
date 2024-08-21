package repository

import (
	"bff-todolist/internal/domain"

	"gorm.io/gorm"
)

type gormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &gormTodoRepository{db: db}
}

func (r *gormTodoRepository) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *gormTodoRepository) GetById(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	result := r.db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (r *gormTodoRepository) Create(t *domain.Todo) error {
	return r.db.Create(t).Error
}
