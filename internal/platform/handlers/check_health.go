package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(db *sql.DB) gin.HandlerFunc {
	var response struct {
		Status   string `json:"status"`
		Database struct {
			Status string `json:"status"`
		} `json:"database"`
		Cache struct {
			Status string `json:"status"`
		} `json:"-"`
	}

	return func(c *gin.Context) {
		response.Status = "up"
		response.Database.Status = "up"

		if err := db.Ping(); err != nil {
			response.Database.Status = "down"
			log.Printf("cannot ping db: %v", err)
		}

		// In the future we could report back on the status of our cache
		// response.Cache.Status = "down"

		c.JSON(http.StatusOK, response)
		c.Writer.Header().Add("Content-Type", "application/json")
	}
}
