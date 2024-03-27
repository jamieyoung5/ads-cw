package classic

import "fmt"

func Play() {
	size := config()
	run(size)
}

func config() (size int) {
	return 9
}

func run(size int) {
	fmt.Printf("running classic sudoku with board size of %d", size)

}
