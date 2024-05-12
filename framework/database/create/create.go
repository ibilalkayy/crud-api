package create

import (
	"database/sql"
	"fmt"

	"github.com/ibilalkayy/crud-api/entities"
	"github.com/ibilalkayy/crud-api/framework/database/read"
)

func CreateTask(db *sql.DB, ct *entities.TaskVariables) error {
	query := "INSERT INTO Task(title, body, statuss, created_at, updated_at) VALUES($1, $2, $3, $4, $5)"
	insert, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	value, err := read.ReadAllTasks(db)
	if err != nil {
		return err
	}

	for _, task := range value {
		if len(ct.Title) != 0 && ct.Title != task.Title {
			_, err = insert.Exec(ct.Title, ct.Body, ct.Status, ct.CreatedAt, ct.UpdatedAt)
			if err != nil {
				return err
			}
			fmt.Println("Task data is successfully inserted!")
			return nil
		}
		fmt.Println("enter a new task")
	}
	return nil
}
