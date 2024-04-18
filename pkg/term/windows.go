//go:build windows

package term

import (
	"syscall"
	"unsafe"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
)

const (
	ENABLE_ECHO_INPUT uint32 = 0x0004
	ENABLE_LINE_INPUT uint32 = 0x0002
)

type windowsTerminal struct {
	handle       syscall.Handle
	originalMode uint32
}

func newTerminal() Terminal {
	handle, _ := syscall.GetStdHandle(syscall.STD_INPUT_HANDLE)
	return &windowsTerminal{handle: handle}
}

func getConsoleMode(handle syscall.Handle) (uint32, error) {
	var mode uint32
	r1, _, e1 := syscall.Syscall(procGetConsoleMode.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(&mode)), 0)
	if r1 == 0 {
		return 0, e1
	}
	return mode, nil
}

func setConsoleMode(handle syscall.Handle, mode uint32) error {
	r1, _, e1 := syscall.Syscall(procSetConsoleMode.Addr(), 2, uintptr(handle), uintptr(mode), 0)
	if r1 == 0 {
		return e1
	}
	return nil
}

func (t *windowsTerminal) EnableRawMode() error {
	mode, err := getConsoleMode(t.handle)
	if err != nil {
		return err
	}
	t.originalMode = mode

	rawMode := mode &^ (ENABLE_ECHO_INPUT | ENABLE_LINE_INPUT)
	return setConsoleMode(t.handle, rawMode)
}

func (t *windowsTerminal) Restore() error {
	return setConsoleMode(t.handle, t.originalMode)
}
