package mocks

import (
	"testing"
	"slices"
	"bytes"
	"time"
)

const (
	write 	= "write"
	sleep	= "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept	time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept	= duration
}

func TestCountdown(t *testing.T) {
	t.Run("should sleep between writes", func(t *testing.T) {
		ops := &SpyCountdownOperations{}

		Countdown(ops, ops)

		got := ops.Calls 
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(got, want) { 
			t.Errorf("got %q want %q", got, want)
		}
	})

		
	t.Run("should print proper message", func(t *testing.T) {
		ops := &SpyCountdownOperations{}
		buffer := &bytes.Buffer{} 

		Countdown(buffer, ops)
		
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("configure sleep time", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
