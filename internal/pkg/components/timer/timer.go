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
	return fmt.Sprintf("Time left: %vs", timeLeft.Seconds())
}

func (t *Timer) Select(pointer *display.Pointer, macro string) (state *display.State, exit bool) {
	return nil, false
}

func (t *Timer) GetDimensions() (height int, width int) {
	return 1, 10
}
