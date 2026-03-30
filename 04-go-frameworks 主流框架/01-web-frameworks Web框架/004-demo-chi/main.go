package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		{ID: 1, Title: "learn chi router", Done: false},
		{ID: 2, Title: "use standard net/http style", Done: true},
	}
	nextID = 3
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"framework": "chi", "status": "ok"})
	})

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		writeJSON(w, http.StatusOK, map[string]any{"items": db, "total": len(db)})
	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		var req createTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if req.Title == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title is required"})
			return
		}

		mu.Lock()
		defer mu.Unlock()

		item := todo{ID: nextID, Title: req.Title, Done: false}
		nextID++
		db = append(db, item)
		writeJSON(w, http.StatusCreated, item)
	})

	http.ListenAndServe(":8084", r)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
