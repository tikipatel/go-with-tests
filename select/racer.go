package racer

import (
	"net/http"
)

// Racer is a function!
func Racer(a, b string) (winner string, error error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
