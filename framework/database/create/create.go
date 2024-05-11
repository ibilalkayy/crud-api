package create

import (
	"database/sql"
	"os"
	"strings"
)

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
