package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-tasks-api/models"
)

var (
	tasks  = make(map[int]models.Task)
	nextID = 1
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	case http.MethodPatch:
		updateTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		var result []models.Task
		for _, t := range tasks {
			result = append(result, t)
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid id",
		})
		return
	}

	task, ok := tasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "task not found",
		})
		return
	}

	json.NewEncoder(w).Encode(task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid title",
		})
		return
	}

	task := models.Task{
		ID:    nextID,
		Title: input.Title,
		Done:  false,
	}

	tasks[nextID] = task
	nextID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid id",
		})
		return
	}

	task, ok := tasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "task not found",
		})
		return
	}

	var input struct {
		Done bool `json:"done"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task.Done = input.Done
	tasks[id] = task

	json.NewEncoder(w).Encode(map[string]bool{
		"updated": true,
	})
}
