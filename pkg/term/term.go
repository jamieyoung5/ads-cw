package term

type Terminal interface {
	EnableRawMode() error
	Restore() error
}

// NewTerminal creates a new terminal handler based on the current operating system.
func NewTerminal() Terminal {
	return newTerminal()
}
