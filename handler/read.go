package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibilalkayy/crud-api/framework/database/read"
)

func ReadTasksHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := read.ReadTask(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}
