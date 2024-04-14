/*package main

import (
	"ads-cw/pkg/game"
	board "ads-cw/pkg/sudoku_board"
)


5 3 # | # 7 # | # # #      5 3 4 | 6 7 8 | 9 1 2
6 # # | 1 9 5 | # # #      6 7 2 | 1 9 5 | 3 4 8
# 9 8 | # # # | # 6 #      1 9 8 | 3 4 2 | 5 6 7
---------------------      ---------------------
8 # # | # 6 # | # # 3      8 5 9 | 7 6 1 | 4 2 3
4 # # | 8 # 3 | # # 1  ->  4 2 6 | 8 5 3 | 7 9 1
7 # # | # 2 # | # # 6      7 1 3 | 9 2 4 | 8 5 6
---------------------      ---------------------
# 6 # | # # # | 2 8 #      9 6 1 | 5 3 7 | 2 8 4
# # # | 4 1 9 | # # 5      2 8 7 | 4 1 9 | 6 3 5
# # # | # 8 # | # 7 9      3 4 5 | 2 8 6 | 1 7 9

func createTestBoard() [][]board.Tile {
	return [][]board.Tile{
		{{5, 0}, {3, 0}, {0, 0}, {0, 1}, {7, 1}, {0, 1}, {0, 2}, {0, 2}, {0, 2}},
		{{6, 0}, {0, 0}, {0, 0}, {1, 1}, {9, 1}, {5, 1}, {0, 2}, {0, 2}, {0, 2}},
		{{0, 0}, {9, 0}, {8, 0}, {0, 1}, {0, 1}, {0, 1}, {0, 2}, {6, 2}, {0, 2}},
		{{8, 3}, {0, 3}, {0, 3}, {0, 4}, {6, 4}, {0, 4}, {0, 5}, {0, 5}, {3, 5}},
		{{4, 3}, {0, 3}, {0, 3}, {8, 4}, {0, 4}, {3, 4}, {0, 5}, {0, 5}, {1, 5}},
		{{7, 3}, {0, 3}, {0, 3}, {0, 4}, {2, 4}, {0, 4}, {0, 5}, {0, 5}, {6, 5}},
		{{0, 6}, {6, 6}, {0, 6}, {0, 7}, {0, 7}, {0, 7}, {2, 8}, {8, 8}, {0, 8}},
		{{0, 6}, {0, 6}, {0, 6}, {4, 7}, {1, 7}, {9, 7}, {0, 8}, {0, 8}, {5, 8}},
		{{0, 6}, {0, 6}, {0, 6}, {0, 7}, {8, 7}, {0, 7}, {0, 8}, {7, 8}, {9, 8}},
	}
}

func main() {
	sudoku := game.NewSudoku()
	sudoku.Play()
}
*/

package main

import (
	"ads-cw/pkg/dlx"
)

func main() {
	/*oldState, err := display.TerminalRawMode()
	if err != nil {
		panic(err)
	}
	defer display.RestoreTerminal(oldState)*/

	//gridMap := &display.ComponentNode{Component: sudoku_board.GenerateBoard(9)}
	//gridMap.Left = &display.ComponentNode{Component: menu.Content, Right: gridMap}
	/*gridMap := [][]*display.ComponentNode{
		{
			&display.ComponentNode{Component: menu.Content}, &display.ComponentNode{Component: sudoku_board.GenerateBoard(9)},
		},
	}

	pointers := display.NewPointer(0, 0, display.StandardControls, 1, 0)
	gridMap[0][1].Pointer = pointers

	canvas := display.NewCanvas(gridMap, []*display.Pointer{pointers})
	canvas.Print()*/
	matrix := dlx.NewMatrix([]string{"column1", "column2", "column3"})
	matrix.AppendRow([]string{"column1", "column2", "column3"})
	matrix.AppendRow([]string{"column1", "column3"})
	matrix.AppendRow([]string{"column3"})
	dlx.PrintMatrix(matrix)

	/*
		testItem1 := "#####\n#   ##\n#   #\n#   #\n#   #\n#####"
		testItem2 := "#####\n#   #\n##   #\n#   #\n#   #\n#####"

		result := display.SideBySide([]string{testItem1, testItem2}, 4)

		fmt.Printf(result)*/
}
