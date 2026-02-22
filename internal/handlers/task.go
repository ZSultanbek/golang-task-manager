package handlers

import (
	"net/http"
	"strconv"
	"encoding/json"
	"strings"

	"practice3/internal/storage"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Handler logic to get tasks
	tasks, err := storage.GlobalTaskStore.GetAllTasks()
	if err != nil {
		http.Error(w, `{"error": "Failed to get tasks"}`, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, `{"error": "Failed to encode tasks"}`, http.StatusInternalServerError)
		return
	}
}

func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Handler logic to get a specific task by ID
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}
	task, exists := storage.GlobalTaskStore.GetTaskByID(id)
	if !exists {
		http.Error(w, `{"error": "task not found"}`, http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, `{"error": "Failed to encode task"}`, http.StatusInternalServerError)
		return
	}
}

func FilterTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Handler logic to filter tasks by done status
	doneStr := r.URL.Query().Get("done")
	if doneStr == "" {
		http.Error(w, `{"error": "missing done query parameter"}`, http.StatusBadRequest)
		return
	}
	done, err := strconv.ParseBool(doneStr)
	if err != nil {
		http.Error(w, `{"error": "invalid done value"}`, http.StatusBadRequest)
		return
	}

	tasks, err := storage.GlobalTaskStore.FilterTasks(done)
	if err != nil {
		http.Error(w, `{"error": "Failed to filter tasks"}`, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, `{"error": "Failed to encode tasks"}`, http.StatusInternalServerError)
		return
	}
}

// creating struct for incoming task data
type CreateTaskRequest struct {
	Title string `json:"title"`
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Handler logic to create a new task
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(req.Title) == "" {
		http.Error(w, `{"error": "invalid title"}`, http.StatusBadRequest)
		return
	}

	if len(req.Title) > 40 {
		http.Error(w, `{"error": "title too long"}`, http.StatusBadRequest)
		return
	}

	task, err := storage.GlobalTaskStore.CreateTask(req.Title)
	if err != nil {
		http.Error(w, `{"error": "Failed to create task"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, `{"error": "Failed to encode task"}`, http.StatusInternalServerError)
		return
	}
}

type UpdateTaskRequest struct {
	Done *bool `json:"done"`
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Handler logic to update a task's done status
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}
	if req.Done == nil {
		http.Error(w, `{"error": "missing done field"}`, http.StatusBadRequest)
		return
	}

	exists := storage.GlobalTaskStore.UpdateTask(id, *req.Done)
	if !exists {
		http.Error(w, `{"error": "task not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"updated": "true"})
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Handler logic to delete a task
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid id"}`, http.StatusBadRequest)
		return
	}

	exists := storage.GlobalTaskStore.DeleteTask(id)
	if !exists {
		http.Error(w, `{"error": "task not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"deleted": "true"})
}