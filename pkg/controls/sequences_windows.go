//go:build windows

package controls

var (
	UpArrow    = Encode([]byte{108}) // ESC [ A
	DownArrow  = Encode([]byte{46})  // ESC [ B
	LeftArrow  = Encode([]byte{44})  // ESC [ D
	RightArrow = Encode([]byte{47})  // ESC [ C

	Enter = Encode([]byte{13}) // Carriage return (might need to be adjusted based on terminal)

	SemiColon  = Encode([]byte{111})
	Apostrophe = Encode([]byte{112})
	Hashtag    = Encode([]byte{91})
	BackTick   = Encode([]byte{96})

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
	F = Encode([]byte{101})
	G = Encode([]byte{114})
	H = Encode([]byte{116})
	J = Encode([]byte{121})
	K = Encode([]byte{117})
	L = Encode([]byte{105})
	M = Encode([]byte{109})
	N = Encode([]byte{110})
	S = Encode([]byte{115})
	W = Encode([]byte{119})
	X = Encode([]byte{120})
	Z = Encode([]byte{122})
)
