package display

import "encoding/base64"

// macros to associate with specific action
const (
	Up       = "Up"
	Down     = "Down"
	Left     = "Left"
	Right    = "Right"
	Undo     = "Undo"
	Redo     = "Redo"
	Selected = "Selected"
	Clear    = "clear"
	One      = "1"
	Two      = "2"
	Three    = "3"
	Four     = "4"
	Five     = "5"
	Six      = "6"
	Seven    = "7"
	Eight    = "8"
	Nine     = "9"
)

var (
	UpArrow    = Encode([]byte{27, 91, 65}) // ESC [ A
	DownArrow  = Encode([]byte{27, 91, 66}) // ESC [ B
	LeftArrow  = Encode([]byte{27, 91, 68}) // ESC [ D
	RightArrow = Encode([]byte{27, 91, 67}) // ESC [ C

	Enter = Encode([]byte{10}) // Carriage return (might need to be adjusted based on terminal)

	SemiColon  = Encode([]byte{59})
	Apostrophe = Encode([]byte{39})
	Hashtag    = Encode([]byte{35})

	OneKey   = Encode([]byte{49})
	TwoKey   = Encode([]byte{50})
	ThreeKey = Encode([]byte{51})
	FourKey  = Encode([]byte{52})
	FiveKey  = Encode([]byte{53})
	SixKey   = Encode([]byte{54})
	SevenKey = Encode([]byte{55})
	EightKey = Encode([]byte{56})
	NineKey  = Encode([]byte{57})

	A = Encode([]byte{97})
	B = Encode([]byte{98})
	C = Encode([]byte{99})
	D = Encode([]byte{100})
	F = Encode([]byte{102})
	G = Encode([]byte{103})
	H = Encode([]byte{104})
	J = Encode([]byte{106})
	K = Encode([]byte{107})
	L = Encode([]byte{108})
	M = Encode([]byte{109})
	N = Encode([]byte{110})
	S = Encode([]byte{115})
	W = Encode([]byte{119})
	X = Encode([]byte{120})
	Z = Encode([]byte{122})
)

// Controls key: base(d)64 encoded byte sequence, value: macro
type Controls map[string]string

var (
	SudokuControls = Controls{
		W: Up,
		S: Down,
		A: Left,
		D: Right,

		C: Undo,
		X: Clear,
		Z: Redo,

		OneKey:   One,
		TwoKey:   Two,
		ThreeKey: Three,
		FourKey:  Four,
		FiveKey:  Five,
		SixKey:   Six,
		SevenKey: Seven,
		EightKey: Eight,
		NineKey:  Nine,
	}

	SudokuControlsAlternate = Controls{
		UpArrow:    Up,
		DownArrow:  Down,
		LeftArrow:  Left,
		RightArrow: Right,

		M: Undo,
		N: Clear,
		B: Redo,

		F:          One,
		G:          Two,
		H:          Three,
		J:          Four,
		K:          Five,
		L:          Six,
		SemiColon:  Seven,
		Apostrophe: Eight,
		Hashtag:    Nine,
	}

	MenuControls = Controls{
		UpArrow:   Up,
		DownArrow: Down,
		Enter:     Selected,
	}
)

func Encode(sequence []byte) string {
	return base64.StdEncoding.EncodeToString(sequence)
}

func Decode(based64 string) []byte {
	bytes, _ := base64.StdEncoding.DecodeString(based64)
	return bytes
}
