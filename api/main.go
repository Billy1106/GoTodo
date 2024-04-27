package main

import (
	"memo-go/domain"
	"memo-go/repository"
	delivery "memo-go/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	tr := repository.NewTodoRepository()
	ts := domain.TodoService(tr)

	c := fiber.New()
	c.Use(cors.New(cors.Config{
		// https://docs.gofiber.io/api/middleware/cors#config
		AllowCredentials: true,
	}))

	delivery.NewTodoAllGetHandler(c, ts)
	delivery.NewTodoStoreHandler(c, ts)
	delivery.NewTodoUpdateHandler(c, ts)
	delivery.NewTodoDeleteHandler(c, ts)

	c.Listen("localhost:8080")
}
