package usecase

import "bff-todolist/internal/domain"

type todoUsecase struct {
	repo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{repo: repo}
}

func (uc *todoUsecase) GetAll() ([]domain.Todo, error) {
	return uc.repo.GetAll()
}

func (uc *todoUsecase) GetById(id string) (*domain.Todo, error) {
	return uc.repo.GetById(id)
}

func (uc *todoUsecase) Create(t *domain.Todo) error {
	return uc.repo.Create(t)
}
