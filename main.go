package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ibilalkayy/crud-api/framework/database"
	"github.com/ibilalkayy/crud-api/handler"
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

	r.GET("/insert_tasks", handler.InsertTasksHandler(db))
	r.GET("/read_tasks", handler.ReadTasksHandler(db))

	// Run the Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
