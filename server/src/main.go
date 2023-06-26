package main

import (
	"container/heap"
	"fmt"
	"io"
	"net/http"
	"os"

	"dawei.io/app/airfile/controller"
	"dawei.io/app/airfile/handler"
	"dawei.io/app/airfile/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Could not load env variables!")
			return
		}
	}

	// Initialize Priority Queue
	pq := make(handler.PriorityQueue, 0)
	heap.Init(&pq)

	// Start cron job
	c := cron.New()
	handler.StartFileDestroyCron(c, &pq)
	c.Start()

	// Logger
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// CORS config
	r.Use(cors.Default())

	// PriorityQueue Middleware
	r.Use(middleware.PriorityQueueMiddleware(&pq))

	// Routes
	r.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the airfile api!")
	})

	// Handle file upload
	r.POST("/file/upload", controller.Upload)
	// Handle file download
	r.GET("/file/download", controller.Download)

	r.Run(":8080")
}
