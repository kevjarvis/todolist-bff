package domain

type Todo struct {
	Id          string
	Name        string
	Description string
	Completed   bool
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(id string) (*Todo, error)
	Create(t *Todo) error
}

type TodoUsecase interface {
	GetAll() ([]Todo, error)
	GetById(id string) (*Todo, error)
	Create(t *Todo) error
}
