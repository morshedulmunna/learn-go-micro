package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morshedulmunna/go-curd/config"
	"github.com/morshedulmunna/go-curd/internal/handlers"
	"github.com/morshedulmunna/go-curd/internal/repositories"
)

func main() {
	config.ConnectDB()

	repo := repositories.NewTodoRepository(config.DB)
	handler := handlers.NewTodoHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/todos", handler.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
