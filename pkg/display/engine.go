package display

import (
	"fmt"
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
	components *ComponentNode
	pointers   []*Pointer
}

func NewCanvas(gridMap *ComponentNode, pointers []*Pointer) *Canvas {
	return &Canvas{
		components: gridMap,
		pointers:   pointers,
	}
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
