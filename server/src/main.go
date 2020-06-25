package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"dawei.io/app/airfile/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Could not load env variables!")
			return
		}
	}

	// Logger
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "OPTION"},
		AllowHeaders:    []string{"Content-Length", "Content-Type", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
	}))

	// Routes
	r.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the airfile api!")
	})

	// Handle file upload
	r.POST("/api/upload", controller.Upload)
	// Handle file download
	r.GET("/api/download", controller.Download)

	r.Run(":8080")
}
