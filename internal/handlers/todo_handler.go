package handlers

import (
	"net/http"

	"github.com/morshedulmunna/go-curd/internal/models"
	"github.com/morshedulmunna/go-curd/internal/repositories"
	"github.com/morshedulmunna/go-curd/pkg"
)

type TodoHandler struct {
	repo repositories.TodoRepository
}

func NewTodoHandler(repo repositories.TodoRepository) *TodoHandler {
	return &TodoHandler{repo}
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repo.GetAll()
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.JSONResponse(w, http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := pkg.ParseJSONRequest(w, r, &todo); err != nil {
		return
	}

	id, err := h.repo.Create(todo)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	todo.ID = id
	pkg.JSONResponse(w, http.StatusCreated, todo)
}
