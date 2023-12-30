package ctx

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"context"
	"log"
	"errors"
)

type SpyStore struct {
	response  string
	cancelled bool
}

type SpyResponseWriter struct {
	written	bool
}

func (w *SpyResponseWriter) Header() http.Header {
	w.written = true
	return nil
}

func (w *SpyResponseWriter) Write([]byte) (int, error) {
	w.written = true
	return 0, errors.New("not implemented")
}

func (w *SpyResponseWriter) WriteHeader(status int) {
	w.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
				case <-ctx.Done():
					log.Println("spy store got cancelled")
					return
				default:
					time.Sleep(10 * time.Millisecond)
					result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("successful fetch", func(t *testing.T) {
		data := "hello, world!"
		store := SpyStore{response: data}
		server := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s" and want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("cancel after 5ms", func(t *testing.T) {
		data 		:= "hello, world!"
		store 	:= &SpyStore{response: data}
		server 	:= Server(store)
		request 	:= httptest.NewRequest(http.MethodGet, "/", nil)
		
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response 	:= &SpyResponseWriter{} 

		server.ServeHTTP(response, request)
		
		if response.written {
			t.Error("a response should not have been written")
		}
	})
}

func (store *SpyStore) assertCancelled(t testing.TB) {
	t.Helper()	
	if !store.cancelled { 
		t.Error("operation was not cancelled")
	}
}
