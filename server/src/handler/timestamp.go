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
