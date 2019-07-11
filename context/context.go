package main

import (
	"context"
	"fmt"
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
		data, err := store.Fetch(r.Context())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, data)
	}
}
