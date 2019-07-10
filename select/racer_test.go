package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()

	slowURL := slowServer.URL //"https://www.facebook.com"
	fastURL := fastServer.URL //"https://www.quii.co.uk"

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestRacerReturnsErrorIfTimeout(t *testing.T) {
	// This function returns an error if a server doesn't respond within 10s

	serverA := makeDelayedServer(11)
	defer serverA.Close()
	serverB := makeDelayedServer(12)
	defer serverB.Close()

	_, err := Racer(serverA.URL, serverB.URL)

	if err == nil {
		t.Error("Expected an error but didn't get one")
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
