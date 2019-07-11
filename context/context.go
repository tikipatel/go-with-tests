package main

import (
	"context"
	"net/http"
)

func main() {

}

// Store is something
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// StubStore is a stub for Store
type StubStore struct {
	response string
}

// Fetch is implementation for Store
func (s *StubStore) Fetch() string {
	return s.response
}

// Server is something
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		// data := make(chan string, 1)

		// go func() {
		// 	data <- store.Fetch()
		// }()

		// select {
		// case d := <-data:
		// 	fmt.Fprint(w, d)
		// case <-ctx.Done():
		// 	store.Cancel()
		// }
	}
}
