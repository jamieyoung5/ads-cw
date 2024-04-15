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
	Components [][]*ComponentNode
	pointers   []*Pointer
}

func NewCanvas(components [][]*ComponentNode, pointers []*Pointer) *Canvas {
	return &Canvas{
		Components: components,
		pointers:   pointers,
	}
}

func (c *Canvas) Render() {
	quit := make(chan bool)
	c.Print()
	go c.draw(quit)

	reader := bufio.NewReader(os.Stdin)

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
			height, width := c.Components[pointer.GridY][pointer.GridX].Component.GetDimensions()
			if macro, ok := pointer.controls[Encode(inputSequence)]; ok {
				switch macro {
				case up:
					if pointer.Y > 0 {
						pointer.Up()
					}
				case down:
					if pointer.Y < height-1 {
						pointer.Down()
					}
				case left:
					if pointer.X > 0 {
						pointer.Left()
					}
				case right:
					if pointer.X < width-1 {
						pointer.Right()
					}
				default:
					//Select with control!

					exit, _ := c.Components[pointer.GridY][pointer.GridX].Component.Select(pointer, inputSequence)
					quit <- exit

				}
				c.Print()
			}
		}
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

func (c *Canvas) Print() {
	fmt.Print("\033[H\033[2J\033[3J")
	canvas := c.serialize()
	fmt.Printf(canvas)
}
