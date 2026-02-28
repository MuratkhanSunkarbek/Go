package main

import (
	"log"
	"net/http"

	"go-tasks-api/handlers"
	"go-tasks-api/middleware"
)

func main() {
	handler := middleware.Logging(
		middleware.Auth(
			http.HandlerFunc(handlers.TasksHandler),
		),
	)

	http.Handle("/tasks", handler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
