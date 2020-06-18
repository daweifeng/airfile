package controller

import (
	"net/http"
	"os"
	"time"

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

	// Check timestamp - File life: 10 mins
	timestamp := handler.FetchTimestamp(name)
	uploadTime := handler.GenerateTime(timestamp)
	timestampMax := uploadTime.Add(10 * time.Minute)
	if timestampMax.Before(time.Now()) {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("REDIRECT_URL"))
	}

	// Download file from s3
	url, err := handler.GetPresignedUrl(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error on fetching url",
		})
		return
	}

	// Response
	c.Redirect(http.StatusTemporaryRedirect, url)
}
