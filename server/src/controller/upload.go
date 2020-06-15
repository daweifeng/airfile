package controller

import (
	"net/http"

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
	var reccord handler.Record
	reccord.Name = file.Filename

	objID, err := handler.CreateRecord(&reccord)
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
}
