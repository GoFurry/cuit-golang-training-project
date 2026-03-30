package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type createTodoRequest struct {
	Title string `json:"title"`
}

var (
	mu = sync.Mutex{}
	db = []todo{
		{ID: 1, Title: "learn fiber handlers", Done: false},
		{ID: 2, Title: "compare with express style", Done: true},
	}
	nextID = 3
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"framework": "fiber", "status": "ok"})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		mu.Lock()
		defer mu.Unlock()
		return c.JSON(fiber.Map{"items": db, "total": len(db)})
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		var req createTodoRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if req.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
		}

		mu.Lock()
		defer mu.Unlock()

		item := todo{ID: nextID, Title: req.Title, Done: false}
		nextID++
		db = append(db, item)
		return c.Status(fiber.StatusCreated).JSON(item)
	})

	app.Listen(":8083")
}
