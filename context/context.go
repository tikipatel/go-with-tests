package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

}

// Store is something
type Store interface {
	Fetch() string
	Cancel()
}

// StubStore is a stub for Store
type StubStore struct {
	response string
}

// Fetch is implementation for Store
func (s *StubStore) Fetch() string {
	return s.response
}

// SpyStore is a spy Store
type SpyStore struct {
	response  string
	cancelled bool
}

// Fetch for Store
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

// Cancel for Store
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

// Server is something
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
