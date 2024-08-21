package repository

import (
	"errors"
	"sync"

	"bff-todolist/internal/domain"

	"github.com/google/uuid"
)

type todoRepository struct {
	todos map[string]domain.Todo
	mutex sync.RWMutex
}

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{
		todos: make(map[string]domain.Todo),
	}
}

func (r *todoRepository) GetAll() ([]domain.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todos := make([]domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoRepository) GetById(id string) (*domain.Todo, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	todo, ok := r.todos[id]
	if !ok {
		return nil, errors.New("todo not found")
	}

	return &todo, nil
}

func (r *todoRepository) Create(t *domain.Todo) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	t.Id = uuid.New().String()
	r.todos[t.Id] = *t
	return nil
}
