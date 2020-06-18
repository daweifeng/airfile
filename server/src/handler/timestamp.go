package handler

import (
	"strconv"
	"strings"
	"time"
)

// AddTimestamp will Add timestamp to string
func AddTimestamp(s string) string {
	now := time.Now()

	nsec := now.UnixNano()

	result := s + "_" + strconv.FormatInt(nsec, 10)

	return result
}

// RemoveTimestamp will remove timestamp from string
func RemoveTimestamp(s string) string {
	sectionArr := strings.Split(s, "_")

	sectionArr = sectionArr[:len(sectionArr)-1]

	result := strings.Join(sectionArr, "_")

	return result
}

// FetchTimestamp will get timestamp from filename
func FetchTimestamp(s string) string {
	sectionArr := strings.Split(s, "_")

	index := len(sectionArr) - 1
	return sectionArr[index]

}

// GenerateTime will generate time object of the timestamp
func GenerateTime(timestamp string) time.Time {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(0, i)
}
