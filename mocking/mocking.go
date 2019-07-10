package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper is an interface that defines `Sleep()`
type Sleeper interface {
	Sleep()
}

// SpySleeper is a struct
type SpySleeper struct {
	Calls int
}

// DefaultSleeper is a struct
type DefaultSleeper struct {
}

// CountdownOperationsSpy is a struct
type CountdownOperationsSpy struct {
	Calls []string
}

func main() {
	sleeper := &ConfigurableSleeper{4 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// Countdown is a function that counts from 3, printing each number on a new line (with a 1 second
// pause) and when it reaches zero it will print "Go!" and exit
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, "Go!")
}

// Sleep is implemented with SpySleeper for conformance to the `Sleeper` interface
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Sleep implementation for DefaultSleeper's conformance to the `Sleeper` interface
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const write = "write"
const sleep = "sleep"

// Sleep implmenetation for `CountdownOperationsSpy`'s conformance to the `Sleeper` interface
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write implementation for `CountdownOperationsSpy`'s conformance to `io.Writer` interface
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// ConfigurableSleeper struct
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// SpyTime struct
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep is a function
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Sleep implmentation for `ConfigurableSleeper` for conformance to `Sleeper`
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
