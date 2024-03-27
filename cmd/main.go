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
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	menuItems := []string{"Item 1", "Item 2", "Item 3"}
	currentIndex := 0

	// Set terminal to raw mode to read characters as they are typed without waiting for ENTER.
	oldState, err := terminalRawMode()
	if err != nil {
		panic(err)
	}
	defer restoreTerminal(oldState) // Ensure the terminal is restored on exit.

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Use arrow keys to navigate the menu and press ENTER to select an item.")
	printMenu(menuItems, currentIndex)

	for {
		input, _ := reader.ReadByte()

		if input == 27 { // Escape character
			next, _ := reader.ReadByte()
			if next == 91 { // ANSI escape sequence for arrow keys starts with Esc[.
				arrowKey, _ := reader.ReadByte()
				switch arrowKey {
				case 65: // Up
					if currentIndex > 0 {
						currentIndex--
					}
				case 66: // Down
					if currentIndex < len(menuItems)-1 {
						currentIndex++
					}
				}
				printMenu(menuItems, currentIndex)
			}
		} else if input == 10 { // Enter key
			fmt.Printf("\nYou selected: %s\n", menuItems[currentIndex])
			break
		}
	}
}

func terminalRawMode() (*syscall.Termios, error) {
	fd := int(os.Stdin.Fd())
	var oldState syscall.Termios
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), syscall.TCGETS, uintptr(unsafe.Pointer(&oldState)), 0, 0, 0); err != 0 {
		return nil, err
	}
	newState := oldState
	newState.Lflag &^= syscall.ECHO | syscall.ICANON // Disable echo and canonical mode
	syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&newState)), 0, 0, 0)

	return &oldState, nil
}

func restoreTerminal(state *syscall.Termios) {
	fd := int(os.Stdin.Fd())
	syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(state)), 0, 0, 0)
}

func printMenu(items []string, currentIndex int) {
	fmt.Print("\033[H\033[2J") // Clear screen
	for i, item := range items {
		if i == currentIndex {
			fmt.Printf("> %s <\n", item) // Highlight the current item
		} else {
			fmt.Println(item)
		}
	}
}
