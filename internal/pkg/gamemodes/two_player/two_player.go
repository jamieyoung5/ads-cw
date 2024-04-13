package two_player

import "fmt"

func Play() {
	size := config()
	run(size)
}

func config() (size int) {
	return 9
}

func run(size int) {
	fmt.Printf("running time trials with board size of %d", size)

}
