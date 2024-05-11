package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibilalkayy/crud-api/entities"
	"github.com/ibilalkayy/crud-api/framework/database"
	"github.com/ibilalkayy/crud-api/framework/database/create"
	"github.com/ibilalkayy/crud-api/middleware"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	err := middleware.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	db, err = database.Connection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	r.GET("/create_table", createTableHandler)
	r.GET("/insert_tasks", insertTasksHandler)

	// Run the Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func createTableHandler(c *gin.Context) {
	_, err := create.CreateTable(db, "framework/database/migrations/db.SQL", 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Table created successfully"})
}

func insertTasksHandler(c *gin.Context) {
	// Sample tasks to be inserted
	tasks := []entities.TaskVariables{
		{Title: "Task 1", Body: "Task Body 1", Status: "Pending", CreatedAt: 2534, UpdatedAt: 43},
	}

	// Insert each task into the database
	for _, task := range tasks {
		err := CreateTask(&task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tasks inserted successfully"})
}

func CreateTask(ct *entities.TaskVariables) error {
	query := "INSERT INTO Task(title, body, statuss, created_at, updated_at) VALUES($1, $2, $3, $4, $5)"
	insert, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	if len(ct.Title) != 0 {
		_, err = insert.Exec(ct.Title, ct.Body, ct.Status, ct.CreatedAt, ct.UpdatedAt)
		if err != nil {
			return err
		}
		fmt.Println("Task data is successfully inserted!")
		return nil
	}
	return errors.New("enter the task")
}
