package sudoku_board

import (
	"ads-cw/pkg/dlx"
	"math"
	"math/rand"
	"time"
)

// GenerateBoard sorry, jamie. - sincerely, past jamie
func GenerateBoard(size int) [][]int {
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	unfilled := make([][2]int, 0)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			unfilled = append(unfilled, [2]int{i, j})
		}
	}

	subGridSize := int(math.Sqrt(float64(size)))
	resultsSize := 0

	for resultsSize != 1 {
		matrix := generateDancingLinksMatrix(size, subGridSize, board)
		results := make([][]*dlx.Node, 0)
		matrix.Search([]*dlx.Node{}, &results)
		resultsSize = len(results)

		randomSolution := results[rand.Intn(resultsSize)]
		solutionBoard, _ := solutionToBoard(randomSolution, size)
		randomPosition := rand.Intn(len(unfilled))

		placement := solutionBoard[unfilled[randomPosition][0]][unfilled[randomPosition][1]]
		board[unfilled[randomPosition][0]][unfilled[randomPosition][1]] = placement
		unfilled = append(unfilled[:randomPosition], unfilled[randomPosition+1:]...)
	}

	return board
}

func addRandomPlacement(board [][]int, size int, subGridSize int) {
	matrix := generateDancingLinksMatrix(size, subGridSize, board)
	results := make([][]*dlx.Node, 0)
	matrix.Search([]*dlx.Node{}, &results)
	if len(results) == 1 {
		//convert board and exit
	}

}

func createTestBoard() *Board {
	board, _ := NewBoard([][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	return board
	/*
		return Board{
			{2, 0, 0, 4},
			{0, 0, 0, 0},
			{1, 0, 0, 2},
			{3, 0, 0, 0},
		}*/
}
