package handler

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadToS3 will upload file to aws s3
func UploadToS3(fileName string, fileSize int64, r io.Reader) (string, error) {

	// Read file
	buffer := make([]byte, fileSize)
	r.Read(buffer)

	//Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	// Upload to S3
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(os.Getenv("S3_BUCKET")),
		Key:                aws.String("files/" + fileName),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(fileSize),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})

	return fileName, err
}

// GetPresignedUrl will get presigned url of the s3 object
func GetPresignedUrl(fileName string) (string, error) {
	//Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	req, _ := s3.New(sess).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String("files/" + fileName),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		return "", err
	}

	return urlStr, nil
}

// RemoveFromS3 will remove file from s3
func RemoveFromS3(fileName string) error {
	//Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String("files/" + fileName),
	}

	// RemoveFromS3
	_, err = s3.New(sess).DeleteObject(input)

	return err
}
