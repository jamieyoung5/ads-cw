package display

import (
	"ads-cw/pkg/types"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type State struct {
	Components [][]*ComponentNode
	Pointers   []*Pointer
	Persist    bool
}

type Canvas struct {
	States *types.Stack[*State]
}

func NewCanvas(components [][]*ComponentNode, pointers []*Pointer, persist bool) *Canvas {
	state := &State{
		Components: components,
		Pointers:   pointers,
		Persist:    persist,
	}
	stateStack := types.NewStack[*State]()
	stateStack.Push(state)

	return &Canvas{
		States: stateStack,
	}
}

func (c *Canvas) Render() {
	quitDrawing := make(chan bool)
	go c.draw(quitDrawing)

	for !c.States.IsEmpty() {
		state := c.States.Peek()

		c.Print()

		reader := bufio.NewReader(os.Stdin)

		// Set terminal to raw mode to properly handle key presses
		if runtime.GOOS != "windows" {
			exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
			exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
		} else {
			// Windows terminal settings should be adjusted if necessary
		}

		c.ListenForInput(state, reader)
	}

	quitDrawing <- true

}

func (c *Canvas) ListenForInput(state *State, reader *bufio.Reader) {
	for {
		inputSequence, err := readKeySequence(reader)
		if err != nil {
			fmt.Println("Error reading from input:", err)
			continue
		}

		fmt.Println("Debug: Key sequence received:", inputSequence) // Debugging output

		for _, pointer := range state.Pointers {
			height, width := state.Components[pointer.GridY][pointer.GridX].Component.GetDimensions()
			if macro, ok := pointer.controls[Encode(inputSequence)]; ok {
				switch macro {
				case Up:
					if pointer.Y > 0 {
						pointer.Up()
					}
				case Down:
					if pointer.Y < height-1 {
						pointer.Down()
					}
				case Left:
					if pointer.X > 0 {
						pointer.Left()
					}
				case Right:
					if pointer.X < width-1 {
						pointer.Right()
					}
				case Undo, Redo, One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Selected, Clear:
					//If the user isnt moving around the board but they pressed a control, execute components action
					nextState, exit := state.Components[pointer.GridY][pointer.GridX].Component.Select(pointer, macro)
					c.Print()

					// current state is finished
					if exit {
						// there is a new state to add
						if nextState != nil {
							//If the current state is NOT persistent, get rid of it
							if !state.Persist {
								c.States.Pop()
							}

							//Add the new state to the top of the stack
							c.States.Push(nextState)
						} else {
							c.States.Pop()
						}

						//exit component
						return
					}
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
