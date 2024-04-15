package time_trials

import "ads-cw/pkg/display"

func Play() *display.State {
	time, size := config()
	return run(size, time)
}

func config() (time int, size int) {
	return 5, 9
}

func run(size int, time int) *display.State {
	return nil
}
