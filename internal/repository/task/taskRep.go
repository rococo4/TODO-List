package task

import (
	"TODO-List/internal/repository/task/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}
func (r *Repository) Create(task model.Task) (*model.Task, error) {
	result, err := r.db.NamedExec(
		"INSERT INTO tasks(name, description,expired_at) VALUES (:name, :description, :expired_at)", task)
	if err != nil {
		return nil, err
	}
	var createdTask model.Task
	taskId, err := result.LastInsertId()
	err = r.db.Get(&createdTask, "SELECT * FROM tasks WHERE id=$1", taskId)
	if err != nil {
		return nil, err
	}
	return &createdTask, nil
}
func (r *Repository) Find(id int) (*model.Task, error) {
	var createdTask model.Task
	err := r.db.Get(&createdTask, "SELECT * FROM tasks WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &createdTask, nil
}
