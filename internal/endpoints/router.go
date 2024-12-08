package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
				"status":  http.StatusOK,
			})
		})
	}

	// return router
}