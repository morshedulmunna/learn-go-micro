package repositories

import (
	"database/sql"

	"github.com/morshedulmunna/go-curd/internal/models"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetByID(id int) (models.Todo, error)
	Create(todo models.Todo) (int, error)
	Update(id int, todo models.Todo) error
	Delete(id int) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoRepository) GetByID(id int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.QueryRow("SELECT id, title, description, completed FROM todos WHERE id = $1", id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *todoRepository) Create(todo models.Todo) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id",
		todo.Title, todo.Description, todo.Completed).Scan(&id)
	return id, err
}

func (r *todoRepository) Update(id int, todo models.Todo) error {
	_, err := r.db.Exec(
		"UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4",
		todo.Title, todo.Description, todo.Completed, id)
	return err
}

func (r *todoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
