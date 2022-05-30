package main

import "time"

type Timer struct {
	lastReached time.Time
	duration    time.Duration
}

func CreateTimer(duration time.Duration) Timer {
	return Timer{
		lastReached: time.Now(),
		duration:    duration,
	}
}

func (t *Timer) HasReached() bool {
	return time.Now().Sub(t.lastReached) >= t.duration
}

func (t *Timer) Reset() {
	t.lastReached = time.Now()
}
