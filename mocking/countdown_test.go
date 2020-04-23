package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type WriterSleeper struct {
	Calls []string
}

func (w *WriterSleeper) Sleep() {
	w.Calls = append(w.Calls, "sleep")
}

func (w *WriterSleeper) Write(p []byte) (n int, err error) {
	w.Calls = append(w.Calls, "write")
	return
}

type SpyDuration struct {
	durationSlept time.Duration
}

func (s *SpyDuration) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("should countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		sleeper := SpySleeper{}

		Countdown(buffer, &sleeper)

		actual := buffer.String()
		expected := "3\n2\n1\nGo!\n"

		if expected != actual {
			t.Errorf("Expected %q, but got %q", expected, actual)
		}
		expectedSleep := 4
		if sleeper.Calls != expectedSleep {
			t.Errorf("Expected %q Calls to Sleep, but go %q", expectedSleep, sleeper.Calls)
		}
	})

	t.Run("should order sleep and write", func(t *testing.T) {
		sleeper := WriterSleeper{}

		Countdown(&sleeper, &sleeper)

		actual := sleeper.Calls
		expected := []string{"sleep", "write", "sleep", "write", "sleep", "write", "sleep", "write"}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v, but got %v", expected, actual)
		}
	})
}

func TestSleeper(t *testing.T) {
	t.Run("should sleep for duration", func(t *testing.T) {
		duration := 2 * time.Second

		spyDuration := SpyDuration{}
		sleeper := DefaultSleeper{duration, spyDuration.Sleep}

		sleeper.Sleep()

		if spyDuration.durationSlept != duration {
			t.Errorf("Expected sleeping duration of %q but got %q", duration, spyDuration.durationSlept)
		}
	})
}
