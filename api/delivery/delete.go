package delivery

import (
	"memo-go/domain"

	"github.com/gofiber/fiber/v2"
)

type todoDeleteHandler struct {
	todoService domain.TodoService
}

func NewTodoDeleteHandler(c *fiber.App, ts domain.TodoService) {
	handler := &todoDeleteHandler{todoService: ts}
	c.Delete("/delete", handler.Delete)
}

func (h *todoDeleteHandler) Delete(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(400).SendString("Bad Request")
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	error := h.todoService.Delete(todo.ID)
	if error != nil {
		c.Status(500).SendString("Internal Server Error")
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(todo)
}
