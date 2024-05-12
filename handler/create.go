package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibilalkayy/crud-api/entities"
	"github.com/ibilalkayy/crud-api/framework/database"
	"github.com/ibilalkayy/crud-api/framework/database/create"
)

func CreateTableHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := database.CreateTable(db, "framework/database/migrations/db.SQL", 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Table created successfully"})
	}
}

func InsertTasksHandler(db *sql.DB) gin.HandlerFunc {
	date := time.Now().Format("2006/01/02")
	return func(c *gin.Context) {
		tasks := []entities.TaskVariables{
			{Title: "Task 1", Body: "Task Body 1", Status: "Pending", CreatedAt: date, UpdatedAt: "0"},
		}

		// Insert each task into the database
		for _, task := range tasks {
			err := create.CreateTask(db, &task)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tasks inserted successfully"})
	}
}
