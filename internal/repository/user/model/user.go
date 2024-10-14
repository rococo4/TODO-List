package model

import "time"

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
}
