package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(id uint) (*Todo, error)
	Create(t *Todo) error
}

type TodoUsecase interface {
	GetAll() ([]Todo, error)
	GetById(id uint) (*Todo, error)
	Create(t *Todo) error
}
