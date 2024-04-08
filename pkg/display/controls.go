package display

const (
	up    string = "up"
	down  string = "down"
	left  string = "left"
	right string = "right"
)

type Controls map[string]byte

var StandardControls = Controls{
	up:    65,
	down:  66,
	right: 67,
	left:  68,
}
