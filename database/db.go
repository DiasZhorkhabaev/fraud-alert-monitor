package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=9123 dbname=fraud_monitor sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	DB = db

	fmt.Println("Connected to PostgreSQL")

	return db, nil
}
