package handler

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

// StartFileDestroyCron will remove file every 10 minutes
func StartFileDestroyCron(c *cron.Cron, pq *PriorityQueue) {
	fmt.Println("Starting FileDestroyCron")
	// Run every 1 minutes
	c.AddFunc("@every 1m", func() {
		for shouldFileDestroy(pq) {
			// Pop from Queue
			item := heap.Pop(pq).(*Item)
			// Remove file from S3
			err := RemoveFromS3(item.Value)
			if err != nil {
				log.Println("Fail to remove file:")
				log.Println(err)

				// Push item back to queue with a new timestamp
				// with a new timestamp 10 minutes later
				item.Timestamp = time.Now().UnixNano() + 600000
				heap.Push(pq, item)
			}
		}
	})
}

func shouldFileDestroy(pq *PriorityQueue) bool {
	if pq.Len() == 0 {
		return false
	}
	ts := (*pq)[0].Timestamp
	uploadTime := GenerateTime(strconv.FormatInt(ts, 10))
	timestampMax := uploadTime.Add(10 * time.Minute)
	return timestampMax.Before(time.Now())
}
