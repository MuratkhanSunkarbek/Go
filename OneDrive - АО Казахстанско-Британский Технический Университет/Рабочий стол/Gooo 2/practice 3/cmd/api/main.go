
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-crud-api/internal/handler"
	"go-crud-api/internal/middleware"
)

func main() {

	h := handler.NewHandler()

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/health", h.Health).Methods("GET")
	r.HandleFunc("/users", h.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
