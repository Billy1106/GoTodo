package repository

import (
	"memo-go/domain"
)

type todoRepository struct {
	todo map[int]domain.Todo
}

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{
		todo: make(map[int]domain.Todo),
	}
}

func (r *todoRepository) AllGet() ([]domain.Todo, error) {
	todo := []domain.Todo{}
	for i := 0; i < len(r.todo); i++ {
		todo = append(todo, r.todo[i])
	}
	return todo, nil
}

func (r *todoRepository) StatusUpdate(todo domain.Todo) error {
	r.todo[todo.ID] = domain.Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	return nil
}

func (r *todoRepository) Store(todo domain.Todo) (domain.Todo, error) {
	todo.ID = len(r.todo)
	r.todo[todo.ID] = todo
	return r.todo[todo.ID], nil
}

func (r *todoRepository) Delete(id int) error {
	if r.todo[id] != (domain.Todo{}) {
		delete(r.todo, id)
	}
	return nil
}
