package delivery

import (
	"memo-go/domain"

	"github.com/gofiber/fiber/v2"
)

type todoAllGetHandler struct {
	todoService domain.TodoService
}

func NewTodoAllGetHandler(c *fiber.App, ts domain.TodoService) {
	handler := &todoAllGetHandler{todoService: ts}
	c.Get("/todos", handler.AllGet)
}

func (h *todoAllGetHandler) AllGet(c *fiber.Ctx) error {
	todos, error := h.todoService.AllGet()
	if error != nil {
		c.Status(500).SendString("Internal Server Error")
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(todos)
}
