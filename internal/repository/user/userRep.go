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
	var createdUser model.User
	err := r.db.QueryRowx(
		"INSERT INTO users(username, password,first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id, username, password, first_name, last_name",
		user.Username, user.Password, user.FirstName, user.LastName).StructScan(&createdUser)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}
func (r *Repository) Find(id int) (*model.User, error) {
	var createdUser model.User
	err := r.db.Get(&createdUser, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}
