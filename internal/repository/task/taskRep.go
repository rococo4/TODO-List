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
	var createdTask model.Task
	err := r.db.QueryRowx(
		"INSERT INTO tasks(name, description,expired_at, user_id) VALUES ($1, $2, $3, $4) RETURNING id, created_at, expired_at, name, description",
		task.Name, task.Description, task.CreatedAt, task.UserId).StructScan(&createdTask)
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
