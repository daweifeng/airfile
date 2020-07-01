package controller

import (
	"container/heap"
	"net/http"
	"strconv"

	"dawei.io/app/airfile/handler"
	"github.com/gin-gonic/gin"
)

// Upload controller
func Upload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")

	// add timestamp to file name
	file.Filename = handler.AddTimestamp(file.Filename)
	f, err := file.Open()
	// send file to aws s3
	_, err = handler.UploadToS3(file.Filename, file.Size, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error on uploading file to S3",
		})
		return
	}

	// save id and key to mongodb
	var record handler.Record
	record.Name = file.Filename

	objID, err := handler.CreateRecord(&record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error on recording",
		})
		return
	}

	id := objID.Hex()

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

	// Add file info to Priority Queue
	pq := c.MustGet("pq").(*handler.PriorityQueue)
	timestamp, _ := strconv.ParseInt(handler.FetchTimestamp(record.Name), 10, 64)
	item := handler.Item{
		Value:     record.Name,
		Timestamp: timestamp,
	}
	heap.Push(pq, &item)
}
