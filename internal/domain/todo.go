package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string
	Description string
	Completed   bool
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
