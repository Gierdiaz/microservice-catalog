package endpoints

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/Gierdiaz/Book/config"
	"github.com/Gierdiaz/Book/internal/setup"
	"github.com/gin-gonic/gin"
)

func InitRouter(config *config.Config, db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies([]string{"172.18.0.4"})

	bookHandler := setup.SetupBook(db)

	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Funcionando com sucesso!",
			"status":  http.StatusOK,
		})	
	})

	v1 := router.Group("/v1")
	{		
		v1.POST("/books", bookHandler.CreateBook)
		v1.GET("/books", bookHandler.GetBooks)
		v1.GET("/books/:id", bookHandler.GetBookById)
		v1.PUT("/books/:id", bookHandler.UpdateBook)
		v1.DELETE("/books/:id", bookHandler.DeleteBook)
	}

	return router
}
