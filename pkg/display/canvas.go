package display

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	asciiEscape        byte = 27
	asciiSquareBRacket byte = 91
	upKey              byte = 65
	downKey            byte = 66
	rightKey           byte = 67
	leftKey            byte = 68
)

type Component interface {
	GetInstructions() string
	GetDimensions() (height int, width int)
	Print(pointerX int, pointerY int)
	Select(pointerX int, pointerY int, keyCode byte)
}

/*
TODO: Select(...) should return *something*... we need a way of knowing if a exit condition is met for the display and if not, reporting errors or whatever.
like for example if the select action changes a tile and the game is won/its an invalid move, how do we deal with that?
*/
func Display(component Component, pointerX int, pointerY int) (exitPointerX int, exitPointerY int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(component.GetInstructions())
	component.Print(pointerX, pointerY)

	height, width := component.GetDimensions()

	for {
		input, _ := reader.ReadByte()

		if input == asciiEscape {
			next, _ := reader.ReadByte()
			if next == asciiSquareBRacket { //Arrow keys ascii sequence is ESC[ + (A/B/C/D)
				arrowKey, _ := reader.ReadByte()
				switch arrowKey {
				case upKey:
					if pointerY > 0 {
						pointerY--
					}
				case downKey:
					if pointerY < height-1 {
						pointerY++
					}
				case rightKey:
					if pointerX < width-1 {
						pointerX++
					}
				case leftKey:
					if pointerX > 0 {
						pointerX--
					}
				}

				fmt.Print("\033[H\033[2J\033[3J") // Clear screen and scrollback buffer
				fmt.Println(component.GetInstructions())
				component.Print(pointerX, pointerY)
			}
		} else {
			component.Select(pointerX, pointerY, input)
			break
		}
	}

	return pointerX, pointerY
}

func TerminalRawMode() (*syscall.Termios, error) {
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

func RestoreTerminal(state *syscall.Termios) {
	fd := int(os.Stdin.Fd())
	syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(state)), 0, 0, 0)
}
