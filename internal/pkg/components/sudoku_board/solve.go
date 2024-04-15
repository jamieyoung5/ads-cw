package sudoku_board

import (
	"ads-cw/pkg/dlx"
	"fmt"
	"strconv"
	"strings"
)

func parseIdentifier(id string) (num int, index int, err error) {
	parts := strings.Split(id, "_")
	if len(parts) != 3 {
		return 0, 0, fmt.Errorf("invalid format of the identifier")
	}
	num, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	index, err = strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, err
	}
	return num, index, nil
}

func solutionToBoard(solution []*dlx.Node, size int) ([][]int, error) {
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}

	for _, node := range solution {
		var row, col, num int
		var err error

		// Cycle through the constraints that this node represents
		for n := node; ; n = n.Right {
			if strings.HasPrefix(n.Column.ID, "col_") {
				num, col, err = parseIdentifier(n.Column.ID)
				if err != nil {
					return nil, err
				}
			} else if strings.HasPrefix(n.Column.ID, "row_") {
				num, row, err = parseIdentifier(n.Column.ID)
				if err != nil {
					return nil, err
				}
			}

			if n.Right == node {
				break
			}
		}

		if num == 0 || row >= size || col >= size {
			return nil, fmt.Errorf("invalid data extracted from node")
		}
		board[row][col] = num
	}

	return board, nil
}

func generateDancingLinksMatrix(size int, subGridSize int, board [][]int) *dlx.Matrix {
	constraints := generateConstraints(size, subGridSize)
	matrix := dlx.NewMatrix(constraints)

	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			squareX := column / subGridSize
			squareY := row / subGridSize
			if board[row][column] == 0 {
				for num := 1; num <= size; num++ {
					matrix.AppendRow(constructRowConstraints(row, column, squareX, squareY, num))
				}
			} else {
				num := board[row][column]
				matrix.AppendRow(constructRowConstraints(row, column, squareX, squareY, num))
			}
		}
	}
	return matrix
}

func constructRowConstraints(row, column, squareX, squareY, num int) []string {
	return []string{
		fmt.Sprintf("col_%d_%d", num, column),
		fmt.Sprintf("row_%d_%d", num, row),
		fmt.Sprintf("cell_%d_%d", row, column),
		fmt.Sprintf("block_%d_%d_%d", num, squareY, squareX),
	}
}

func generateConstraints(size, gridSize int) []string {
	var constraints []string

	for col := 0; col < size; col++ {
		for num := 1; num <= size; num++ {
			constraints = append(constraints, fmt.Sprintf("col_%d_%d", num, col))
			constraints = append(constraints, fmt.Sprintf("row_%d_%d", num, col))
		}
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			constraints = append(constraints, fmt.Sprintf("cell_%d_%d", row, col))
		}
	}

	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			for num := 1; num <= size; num++ {
				constraints = append(constraints, fmt.Sprintf("block_%d_%d_%d", num, row, col))
			}
		}
	}

	return constraints
}
