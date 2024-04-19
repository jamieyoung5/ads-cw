package sudoku_board

import (
	"fmt"
	"testing"
	"time"
)

func TestBoard_Validate(t *testing.T) {
	var total int64 = 0
	for i := 0; i < 10; i++ {

		startSolve := time.Now()
		GenerateBoard(9)
		endSolve := time.Now()

		total += endSolve.Sub(startSolve).Nanoseconds()
		fmt.Printf("generated 9x9 board with a time of: %s\n", endSolve.Sub(startSolve))

	}
	fmt.Printf("generated boards with an average time of: %s\n", time.Duration(total/10))
}
