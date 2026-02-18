package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := initDatabase(); err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer closeDatabase()

	server := gin.Default()
	server.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	ProductRoutes(server)

	server.Run(":8080")
}
