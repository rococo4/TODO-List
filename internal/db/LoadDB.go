package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func LoadDatabase() (*sqlx.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	log.Println(connStr)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
