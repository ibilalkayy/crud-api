package read

import (
	"database/sql"

	"github.com/ibilalkayy/crud-api/entities"
)

func ReadAllTasks(db *sql.DB) ([]entities.TaskVariables, error) {
	rows, err := db.Query("SELECT title, body, statuss, created_at, updated_at FROM Task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.TaskVariables

	for rows.Next() {
		var task entities.TaskVariables
		if err := rows.Scan(&task.Title, &task.Body, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func ReadSpecificTask(title string, db *sql.DB) ([]entities.TaskVariables, error) {
	var rows *sql.Rows
	var err error

	if len(title) != 0 {
		query := "SELECT title, body, statuss, created_at, updated_at FROM Task WHERE title=$1"
		rows, err = db.Query(query, title)
	} else {
		query := "SELECT title, body, statuss, created_at, updated_at FROM Task"
		rows, err = db.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.TaskVariables
	for rows.Next() {
		var task entities.TaskVariables
		if err := rows.Scan(&task.Title, &task.Body, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
