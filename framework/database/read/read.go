package read

import (
	"database/sql"
	"fmt"

	"github.com/ibilalkayy/crud-api/entities"
)

func ReadTask(db *sql.DB) ([]entities.TaskVariables, error) {
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

	fmt.Println(tasks)

	return tasks, nil
}
