package user

import (
	"TODO-List/internal/repository/user/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}
func (r *Repository) Create(user model.User) (*model.User, error) {
	result, err := r.db.NamedExec(
		"INSERT INTO users(username, password,first_name, last_name) VALUES (:username, :password, :first_name, :last_name)", user)
	if err != nil {
		return nil, err
	}
	var createdUser model.User
	userID, _ := result.LastInsertId()
	err = r.db.Get(&createdUser, "SELECT * FROM users WHERE id=$1", userID)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}
func (r *Repository) Find(id int) (*model.User, error) {
	var createdTask model.User
	err := r.db.Get(&createdTask, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &createdTask, nil
}
