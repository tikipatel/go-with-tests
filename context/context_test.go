package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const data = "hello, world"

func TestHandler(t *testing.T) {

	store := &SpyStore{response: data}
	svr := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)
	if response.Body.String() != data {
		t.Errorf("got %s want %s", response.Body.String(), data)
	}
}

func TestStore(t *testing.T) {

	t.Run("reutnrs data from store", func(t *testing.T) {
		store := SpyStore{response: data}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s' want '%s'", response.Body.String(), data)
		}

		if store.cancelled {
			t.Error("store should not have canceled")
		}
	})

	t.Run("tell store to cancel work if request is canceled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}

	})
}
