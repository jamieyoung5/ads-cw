//go:build linux
// +build linux

package controls

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
