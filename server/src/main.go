package main

import (
	"net/http"

	"dawei.io/app/airfile/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the airfile api!")
	})

	// Handle file upload
	r.POST("/api/upload", controller.Upload)

	r.Run(":8080")
}
