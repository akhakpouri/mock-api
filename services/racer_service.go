package services

import (
	"fmt"
	"net/http"
	"time"
)

func Race(first, second string) (winner string, err error) {
	timeout := (10 * time.Second)
	return configureRacer(first, second, timeout)
}

func configureRacer(first, second string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(first):
		return first, nil
	case <-ping(second):
		return second, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out while waiting for %s and %s", first, second)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
