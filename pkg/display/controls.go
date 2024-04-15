package display

const (
	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"
	enter = "enter"
)

type Control struct {
	Sequence []byte
	Movement bool
}
type Controls map[string]Control

var StandardControls = Controls{
	up:    Control{Sequence: []byte{27, 91, 65}, Movement: true}, // ESC [ A
	down:  Control{Sequence: []byte{27, 91, 66}, Movement: true}, // ESC [ B
	left:  Control{Sequence: []byte{27, 91, 68}, Movement: true}, // ESC [ D
	right: Control{Sequence: []byte{27, 91, 67}, Movement: true}, // ESC [ C
	enter: Control{Sequence: []byte{10}},                         // Carriage return (might need to be adjusted based on your terminal)
	"1":   Control{Sequence: []byte{49}},
	"2":   Control{Sequence: []byte{50}},
	"3":   Control{Sequence: []byte{51}},
	"4":   Control{Sequence: []byte{52}},
	"5":   Control{Sequence: []byte{53}},
	"6":   Control{Sequence: []byte{54}},
	"7":   Control{Sequence: []byte{55}},
	"8":   Control{Sequence: []byte{56}},
	"9":   Control{Sequence: []byte{57}},
}

var MenuControls = Controls{
	up:    Control{Sequence: []byte{27, 91, 65}, Movement: true},
	down:  Control{Sequence: []byte{27, 91, 66}, Movement: true},
	enter: Control{Sequence: []byte{10}}, // Carriage return (might need to be adjusted based on your terminal)
}
