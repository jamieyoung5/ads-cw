package display

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	asciiEscape        byte = 27
	asciiSquareBRacket byte = 91
	upKey              byte = 65
	downKey            byte = 66
	rightKey           byte = 67
	leftKey            byte = 68
)

type Canvas struct {
	components [][]*ComponentNode
	pointers   []*Pointer
}

func NewCanvas(gridMap [][]*ComponentNode, pointers []*Pointer) *Canvas {
	return &Canvas{
		components: gridMap,
		pointers:   pointers,
	}
}

func (c *Canvas) Render() {
	reader := bufio.NewReader(os.Stdin)
	c.Print()

	// Set terminal to raw mode to properly handle key presses
	if runtime.GOOS != "windows" {
		exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
		exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
	} else {
		// Windows terminal settings should be adjusted if necessary
	}

	for {
		inputSequence, err := readKeySequence(reader)
		if err != nil {
			fmt.Println("Error reading from input:", err)
			continue
		}

		fmt.Println("Debug: Key sequence received:", inputSequence) // Debugging output

		for _, pointer := range c.pointers {
			for action, controlSequence := range pointer.controls {
				if equal(inputSequence, controlSequence.Sequence) {
					if !controlSequence.Movement {
						c.components[pointer.GridY][pointer.GridX].Component.Select(pointer, controlSequence.Sequence)
					} else {
						switch action {
						case up:
							pointer.Up()
						case down:
							pointer.Down()
						case left:
							pointer.Left()
						case right:
							pointer.Right()
						}
					}
					break // Assuming one key action per loop is sufficient
				}
			}
		}

		fmt.Print("\033[H\033[2J\033[3J") // Clear screen and scrollback buffer
		c.Print()
	}
}

func readKeySequence(reader *bufio.Reader) ([]byte, error) {
	input, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	keySequence := []byte{input}
	if input == 27 { // ESC character, indicating the start of a control sequence
		// Read the full sequence (assuming fixed length for simplicity)
		for i := 0; i < 2; i++ {
			if nextByte, err := reader.ReadByte(); err == nil {
				keySequence = append(keySequence, nextByte)
			}
		}
	}

	return keySequence, nil
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/*
func (c *Canvas) Render() {
	reader := bufio.NewReader(os.Stdin)
	c.Print()

	height, width := component.GetDimensions()

	for {
		input, _ := reader.ReadByte()

		if input == asciiEscape {

		} else {
			exit, err := component.Select(pointerX, pointerY, input)
			if err != nil {
				fmt.Println(err)
			}
			if exit {
				break
			}
		}
	}

	return pointerX, pointerY
}*/

func (c *Canvas) Print() {
	canvas := draw(c.components)
	fmt.Printf(canvas)
}

/*
func (c *Canvas) registerInput() {
	next, _ := reader.ReadByte()
	if next == asciiSquareBRacket { //Arrow keys ascii sequence is ESC[ + (A/B/C/D)
		arrowKey, _ := reader.ReadByte()
		switch arrowKey {
		case upKey:
			if c.pointer.y > 0 {
				c.pointer.Up()
			}
		case downKey:
			if c.pointer.y < height-1 {
				c.pointer.Down()
			}
		case rightKey:
			if c.pointer.x < width-1 {
				c.pointer.Right()
			}
		case leftKey:
			if c.pointer.x > 0 {
				c.pointer.Left()
			}
		}

		fmt.Print("\033[H\033[2J\033[3J") // Clear screen and scrollback buffer
		component.Print(pointerX, pointerY)
	}
}*/
