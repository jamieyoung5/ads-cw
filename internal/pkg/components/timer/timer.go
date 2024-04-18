package timer

import (
	"ads-cw/pkg/display"
	"fmt"
	"sync"
	"time"
)

type Timer struct {
	Duration    time.Duration
	Start       time.Time
	Finished    bool
	mutex       sync.Mutex
	updateTimer *time.Timer
}

func NewTimer(duration time.Duration) *Timer {
	t := &Timer{
		Duration: duration,
		Start:    time.Now(),
		Finished: false,
	}
	go t.startCountdown()
	return t
}

func (t *Timer) startCountdown() {
	t.updateTimer = time.NewTimer(t.Duration)
	<-t.updateTimer.C
	t.mutex.Lock()
	t.Finished = true
	t.mutex.Unlock()
}

func (t *Timer) Print(pointer *display.Pointer) {
	fmt.Printf(t.Serialize(pointer))
}

func (t *Timer) Serialize(pointer *display.Pointer) string {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.Finished {
		return "Times up!"
	}

	timeLeft := time.Until(t.Start.Add(t.Duration))
	return fmt.Sprintf("Time left: %v", formatDuration(timeLeft))
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)

	hours := d.Hours()
	minutes := d.Minutes()
	seconds := d.Seconds()

	if hours >= 1 {

		min := int(minutes) % 60
		return fmt.Sprintf("%dh%dm", int(hours), min)
	} else if minutes >= 1 {

		sec := int(seconds) % 60
		return fmt.Sprintf("%dm%ds", int(minutes), sec)
	}

	return fmt.Sprintf("%ds", int(seconds))
}

func (t *Timer) Select(pointer *display.Pointer, macro string) (state *display.State, exit bool) {
	return nil, false
}

func (t *Timer) GetDimensions() (height int, width int) {
	return 1, 10
}
