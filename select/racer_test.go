package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare speeds of servers, return the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL //"https://www.facebook.com"
		fastURL := fastServer.URL //"https://www.quii.co.uk"

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error, but got one %v", err)
		}

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond in 20ms", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}

// func TestRacerReturnsErrorIfTimeout(t *testing.T) {
// 	// This function returns an error if a server doesn't respond within 10s

// 	serverA := makeDelayedServer(11 * time.Second)
// 	serverB := makeDelayedServer(12 * time.Second)

// 	defer serverA.Close()
// 	defer serverB.Close()

// 	got, err := Racer(serverA.URL, serverB.URL)

// 	if err == nil {
// 		t.Errorf("got this -- %s", got)
// 		t.Error("Expected an error but didn't get one")
// 	}
// }

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
