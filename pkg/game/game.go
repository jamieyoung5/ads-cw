package game

import (
	"ads-cw/pkg/menu"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Sudoku struct {
}

func NewSudoku() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) Play() {
	for {
		menu.Content.Display()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		menuItem, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		menu.Content[menuItem].Runner()
	}
}
