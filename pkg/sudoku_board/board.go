package sudoku_board

import "fmt"

type Board [][]Tile

func (b Board) Print(pointerX int, pointerY int) {
	fmt.Printf(SerializeBoard(b, pointerX, pointerY))
}

func (b Board) GetInstructions() string {
	return "Use arrow keys to navigate board, enter digit to edit tile"
}

func (b Board) GetDimensions() (height int, width int) {
	return len(b), len(b[0])
}

func (b Board) Select(pointerX int, pointerY int, keycode byte) {
	if !(keycode > '1' && keycode <= '9') {
		fmt.Printf("Please enter a number between 1 and 9!")
		return
	}

	b[pointerY][pointerX].Value = int(keycode - '0')
}

func GenerateBoard(size int) Board {
	return createTestBoard()
}

func createTestBoard() Board {
	return Board{
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
