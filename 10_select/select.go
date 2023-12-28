package slect

import (
	"net/http"
	"time"
	"errors"
)

var ErrTimeout = errors.New("Racer timeout after 10s")

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10 * time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-pingUrl(a):
		return a, nil
	case <-pingUrl(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func pingUrl(url string) chan struct{} {
	ch := make(chan struct{})

	go func(url string) {
		http.Get(url)
		close(ch)
	}(url)

	return ch
}
