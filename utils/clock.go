package utils

import "time"

type TimeTracker struct {
	start int64
	end   int64
}

func NewTimeTracker() *TimeTracker {
	return &TimeTracker{
		start: 0,
		end:   0,
	}
}

func (t *TimeTracker) Start() {
	t.start = Now()
}

func (t *TimeTracker) End() {
	t.end = Now()
}

func (t *TimeTracker) Elapsed() int64 {
	return t.end - t.start
}

func Now() int64 {
	return int64(time.Now().UnixNano() / int64(time.Millisecond))
}
