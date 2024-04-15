package game_history

import (
	"ads-cw/pkg/display"
	"fmt"
)

func Open() *display.State {
	return run()
}

func run() *display.State {
	fmt.Printf("displaying game history")
	return nil
}
