package timer

import (
	"fmt"
	"sync"
	"time"
)

type Timer struct {
	TimeLeft time.Duration
	Mutex    sync.Mutex
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{TimeLeft: duration}
}

func (t *Timer) Start(done chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				t.Mutex.Lock()
				if t.TimeLeft > 0 {
					t.TimeLeft--
					fmt.Println("Time Left:", t.TimeLeft)
				} else {
					ticker.Stop()
					done <- true
					return
				}
				t.Mutex.Unlock()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
}
