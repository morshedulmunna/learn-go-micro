package services

import (
	"github.com/morshedulmunna/go-curd/internal/models"
	"github.com/morshedulmunna/go-curd/internal/repositories"
)

type TodoService interface {
	GetTodos() ([]models.Todo, error)
	GetTodoByID(id int) (models.Todo, error)
	CreateTodo(todo models.Todo) (int, error)
	UpdateTodo(id int, todo models.Todo) error
	DeleteTodo(id int) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (s *todoService) GetTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetTodoByID(id int) (models.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *todoService) CreateTodo(todo models.Todo) (int, error) {
	return s.repo.Create(todo)
}

func (s *todoService) UpdateTodo(id int, todo models.Todo) error {
	return s.repo.Update(id, todo)
}

func (s *todoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}
