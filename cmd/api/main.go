package main

import (
	"fmt"
	"context"
	"time"

	_postgres "task-manager/internal/repository/postgres"
	"task-manager/pkg/modules"
	"task-manager/internal/handlers"
	"task-manager/internal/usecase"
	"net/http"
	"github.com/gorilla/mux"
	"task-manager/internal/middleware"
	"os"
)
/*
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
	*/

func main() {
	Run()
}

func Run() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := initPostgreConfig()

	pg := _postgres.New(dbConfig)

	userRepo := _postgres.NewUserRepository(pg)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	r := mux.NewRouter()

	//  middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.APIKeyAuth)

	// routes
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Username:    os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		SSLMode:     os.Getenv("DB_SSLMODE"),
		ExecTimeout: 5 * time.Second,
	}
}