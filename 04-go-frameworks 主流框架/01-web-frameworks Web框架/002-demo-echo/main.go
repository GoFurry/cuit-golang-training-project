package main

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		{ID: 1, Title: "learn echo context", Done: false},
		{ID: 2, Title: "return structured JSON", Done: true},
	}
	nextID = 3
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"framework": "echo",
			"status":    "ok",
		})
	})

	e.GET("/todos", func(c echo.Context) error {
		mu.Lock()
		defer mu.Unlock()

		return c.JSON(http.StatusOK, map[string]any{
			"items": db,
			"total": len(db),
		})
	})

	e.POST("/todos", func(c echo.Context) error {
		var req createTodoRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		if req.Title == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "title is required"})
		}

		mu.Lock()
		defer mu.Unlock()

		item := todo{ID: nextID, Title: req.Title, Done: false}
		nextID++
		db = append(db, item)
		return c.JSON(http.StatusCreated, item)
	})

	e.Logger.Fatal(e.Start(":8082"))
}
