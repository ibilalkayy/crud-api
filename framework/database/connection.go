package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

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

func CreateTable(db *sql.DB, filename string, number int) (*sql.DB, error) {
	query, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	requests := strings.Split(string(query), ";")[number]
	stmt, err := db.Prepare(requests)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	return db, nil
}
