package domain

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Usecase
type TodoService interface {
	AllGet() ([]Todo, error)
	StatusUpdate(todo Todo) error
	Store(todo Todo) (Todo, error)
	Delete(id int) error
}

type TodoRepository interface {
	AllGet() ([]Todo, error)
	StatusUpdate(todo Todo) error
	Store(todo Todo) (Todo, error)
	Delete(id int) error
}
