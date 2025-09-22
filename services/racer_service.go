package services

import (
	"net/http"
	"time"
)

func Race(first, second string) (winner string) {
	firstResult := duration(first)
	secondResult := duration(second)

	if firstResult >= secondResult {
		return first
	}
	return second
}

func duration(url string) int64 {
	start := time.Now()
	http.Get(url)
	return int64(time.Since(start).Seconds())
}
