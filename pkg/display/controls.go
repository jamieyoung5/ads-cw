package display

import "encoding/base64"

// macros to associate with specific action
const (
	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"
)

var (
	UpArrow    = Encode([]byte{27, 91, 65}) // ESC [ A
	DownArrow  = Encode([]byte{27, 91, 66}) // ESC [ B
	LeftArrow  = Encode([]byte{27, 91, 68}) // ESC [ D
	RightArrow = Encode([]byte{27, 91, 67}) // ESC [ C
	Enter      = Encode([]byte{10})         // Carriage return (might need to be adjusted based on terminal)
	W          = Encode([]byte{119})
	A          = Encode([]byte{97})
	S          = Encode([]byte{115})
	D          = Encode([]byte{100})
	One        = Encode([]byte{49})
	Two        = Encode([]byte{50})
	Three      = Encode([]byte{51})
	Four       = Encode([]byte{52})
	Five       = Encode([]byte{53})
	Six        = Encode([]byte{54})
	Seven      = Encode([]byte{55})
	Eight      = Encode([]byte{56})
	Nine       = Encode([]byte{57})
)

// Controls key: base(d)64 encoded byte sequence, value: macro
type Controls map[string]string

var (
	SudokuControls = Controls{
		UpArrow:    up,
		DownArrow:  down,
		LeftArrow:  left,
		RightArrow: right,
		Enter:      "",
		One:        "",
		Two:        "",
		Three:      "",
		Four:       "",
		Five:       "",
		Six:        "",
		Seven:      "",
		Eight:      "",
		Nine:       "",
	}

	SudokuControlsAlternate = Controls{
		W:     up,
		S:     down,
		A:     left,
		D:     right,
		Enter: "",
		One:   "",
		Two:   "",
		Three: "",
		Four:  "",
		Five:  "",
		Six:   "",
		Seven: "",
		Eight: "",
		Nine:  "",
	}

	MenuControls = Controls{
		UpArrow:   up,
		DownArrow: down,
		Enter:     "",
	}
)

func Encode(sequence []byte) string {
	return base64.StdEncoding.EncodeToString(sequence)
}

func Decode(based64 string) []byte {
	bytes, _ := base64.StdEncoding.DecodeString(based64)
	return bytes
}
