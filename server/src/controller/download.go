package controller

import (
	"net/http"

	"dawei.io/app/airfile/handler"
	"github.com/gin-gonic/gin"
)

// Download route controller
func Download(c *gin.Context) {
	id := c.Request.URL.Query().Get("key")

	// Fetch filename record
	name, err := handler.GetRecord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error on fetching recording",
		})
		return
	}

	// Check timestamp

	// Download file from s3
	url, err := handler.GetPresignedUrl(name)

}
