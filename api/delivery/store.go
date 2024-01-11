package delivery

import (
	"memo-go/domain"

	"github.com/gofiber/fiber/v2"
)

type todoStoreHandler struct {
	todoService domain.TodoService
}

func NewTodoStoreHandler(c *fiber.App, ts domain.TodoService) {
	handler := &todoStoreHandler{todoService: ts}
	c.Post("/store", handler.Store)
}

func (h *todoStoreHandler) Store(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(400).SendString("Bad Request")
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	newTodo, error := h.todoService.Store(*todo)
	if error != nil {
		c.Status(500).SendString("Internal Server Error")
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(newTodo)
}
