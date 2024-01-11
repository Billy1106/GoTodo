package delivery

import (
	"memo-go/domain"

	"github.com/gofiber/fiber/v2"
)

type todoUpdateHandler struct {
	todoService domain.TodoService
}

func NewTodoUpdateHandler(c *fiber.App, ts domain.TodoService) {
	handler := &todoUpdateHandler{todoService: ts}
	c.Post("/update", handler.StatusUpdate)
}

func (h *todoUpdateHandler) StatusUpdate(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(400).SendString("Bad Request")
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	error := h.todoService.StatusUpdate(*todo)
	if error != nil {
		c.Status(500).SendString("Internal Server Error")
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(todo)
}
