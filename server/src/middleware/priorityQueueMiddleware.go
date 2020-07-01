package middleware

import (
	"dawei.io/app/airfile/handler"
	"github.com/gin-gonic/gin"
)

func PriorityQueueMiddleware(pq *handler.PriorityQueue) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pq", pq)
		c.Next()
	}
}
