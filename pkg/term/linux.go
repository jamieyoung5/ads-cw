//go:build linux

package term

import (
	"os"
	"syscall"
	"unsafe"
)

type linuxTerminal struct {
	fd        int
	origState syscall.Termios
}

func newTerminal() Terminal {
	fd := int(os.Stdin.Fd())
	return &linuxTerminal{fd: fd}
}

func (t *linuxTerminal) EnableRawMode() error {
	var termios syscall.Termios
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCGETS, uintptr(unsafe.Pointer(&termios)), 0, 0, 0); err != 0 {
		return err
	}
	t.origState = termios

	newState := termios
	newState.Lflag &^= syscall.ECHO | syscall.ICANON
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCSETS, uintptr(unsafe.Pointer(&newState)), 0, 0, 0)
	if err != 0 {
		return syscall.Errno(err)
	}
	return nil
}

func (t *linuxTerminal) Restore() error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCSETS, uintptr(unsafe.Pointer(&t.origState)), 0, 0, 0)
	return err
}
