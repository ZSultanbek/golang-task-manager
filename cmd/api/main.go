package main

import (
	"log"
	"net/http"
	"fmt"

	"assignment1/internal/middleware"
	"assignment1/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("GET /tasks", handlers.GetTaskHandler)
	// mux.HandleFunc("GET /tasks/{id}", handlers.GetTaskByIDHandler)
	// mux.HandleFunc("POST /tasks", handlers.CreateTaskHandler)
	// mux.HandleFunc("PATCH /tasks", handlers.UpdateTaskHandler)

	protectedMux := middleware.APIKeyAuth(middleware.LoggingMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Has("id") {
				handlers.GetTaskByIDHandler(w, r)
			} else if r.URL.Query().Has("done") {
				handlers.FilterTasksHandler(w, r)
			} else {
				handlers.GetTaskHandler(w, r)
			}
		case http.MethodPost:
			handlers.CreateTaskHandler(w, r)
		case http.MethodPatch:
			handlers.UpdateTaskHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteTaskHandler(w, r)
		default:
			http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		}
	}),
	),)
	mux.Handle("/tasks", protectedMux)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("could not start server: %s\n", err.Error())
	}
}