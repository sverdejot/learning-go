package slect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	t.Run("racer return first url", func(t *testing.T) {
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowServer := makeDelayedServer(20 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		want := fastServer.URL
		got, _ := ConfigurableRacer(slowServer.URL, fastServer.URL, 21 * time.Millisecond)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("racer return error on time 10", func(t *testing.T) {
		fastServer := makeDelayedServer(2 * time.Millisecond)
		slowServer := makeDelayedServer(3 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, time.Millisecond)

		if err == nil {
			t.Errorf("expected error and got nothing")
		}
	})
}

func makeDelayedServer(sleepTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))
}
