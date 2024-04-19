package controls

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
	Exit     = "exit"
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
		BackTick: Exit,
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
		BackTick:   Exit,
	}

	MenuControls = Controls{
		W:     Up,
		S:     Down,
		Enter: Selected,
	}
)

func Encode(sequence []byte) string {
	return base64.StdEncoding.EncodeToString(sequence)
}

func Decode(based64 string) []byte {
	bytes, _ := base64.StdEncoding.DecodeString(based64)
	return bytes
}
