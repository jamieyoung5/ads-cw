package time_trials

import "fmt"

func Play() {
	time, size := config()
	run(size, time)
}

func config() (time int, size int) {
	return 5, 9
}

func run(size int, time int) {
	fmt.Printf("running time trials with board size of %d and a time limit of %d minutes", size, time)

}
