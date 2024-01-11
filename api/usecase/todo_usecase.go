package usecase

import (
	"memo-go/domain"
)

// Usecase =　apiで呼ばれ、Repositoryを呼び出す関数
type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(tr domain.TodoRepository) domain.TodoService {
	return &todoUsecase{
		todoRepo: tr,
	}
}

func (u *todoUsecase) AllGet() ([]domain.Todo, error) {
	return u.todoRepo.AllGet()
}

func (u *todoUsecase) StatusUpdate(todo domain.Todo) error {
	return u.todoRepo.StatusUpdate(todo)
}

func (u *todoUsecase) Store(todo domain.Todo) (domain.Todo, error) {
	return u.todoRepo.Store(todo)
}

func (u *todoUsecase) Delete(id int) error {
	return u.todoRepo.Delete(id)
}
