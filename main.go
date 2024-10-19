package main

import (
	"net/http"

	dbcreator "github.com/finestgit/statement-evaluator-backend/dbCreator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}

	server.Use(cors.New(corsConfig))

	server.GET("/health", healthCheck)

	dbcreator.RegisterHeaderRoutes(server)

	server.Run(":8080")
}

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "Healthy"})
}
