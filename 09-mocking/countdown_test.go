package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

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

const (
	write = "write"
	sleep = "sleep"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.durationSlept = d
}

func TestCountdown(t *testing.T) {
	t.Run("countdown result", func(t *testing.T) {
		buf := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buf, spySleeper)

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 1 * time.Second

	spy := SpyTime{sleepTime}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spy.Sleep,
	}

	sleeper.Sleep()

	if spy.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spy.durationSlept)
	}
}
