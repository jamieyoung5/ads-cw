package display

const (
	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"
	enter = "enter"
)

type Controls map[string][]byte

var StandardControls = Controls{
	up:    []byte{27, 91, 65}, // ESC [ A
	down:  []byte{27, 91, 66}, // ESC [ B
	right: []byte{27, 91, 67}, // ESC [ C
	left:  []byte{27, 91, 68}, // ESC [ D
	enter: []byte{13},         // Carriage return (might need to be adjusted based on your terminal)
}
