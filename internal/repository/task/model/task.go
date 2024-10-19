package model

import "time"

type Task struct {
	Id          int       `db:"id" `
	CreatedAt   time.Time `db:"created_at" `
	ExpiredAt   time.Time `db:"expired_at" `
	Name        string    `db:"name" `
	Description string    `db:"description" `
	UserId      int       `db:"user_id" `
}
