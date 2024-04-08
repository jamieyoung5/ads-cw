package display

import (
	"os"
	"syscall"
	"unsafe"
)

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
