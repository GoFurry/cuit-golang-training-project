package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type createTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

var (
	mu = sync.Mutex{}
	db = []todo{
		{ID: 1, Title: "learn gin routing", Done: false},
		{ID: 2, Title: "build a JSON API", Done: true},
	}
	nextID = 3
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"framework": "gin", "status": "ok"})
	})

	router.GET("/todos", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		c.JSON(http.StatusOK, gin.H{"items": db, "total": len(db)})
	})

	router.POST("/todos", func(c *gin.Context) {
		var req createTodoRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		defer mu.Unlock()

		item := todo{ID: nextID, Title: req.Title, Done: false}
		nextID++
		db = append(db, item)
		c.JSON(http.StatusCreated, item)
	})

	router.Run(":8081")
}
