package two_player

import (
	"ads-cw/pkg/display"
	"fmt"
)

func Play() *display.State {
	size := config()
	return run(size)
}

func config() (size int) {
	return 9
}

func run(size int) *display.State {
	fmt.Printf("running time trials with board size of %d", size)
	return nil
}
