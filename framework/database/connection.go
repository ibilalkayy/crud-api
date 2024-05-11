package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	url := os.Getenv("DB_URL")
	if len(url) == 0 {
		return nil, errors.New("DB_URL environment is not set")
	}
	var err error
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
