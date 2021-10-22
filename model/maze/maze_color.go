package maze

type Color int

const (
	White Color = iota
	Black
	Silver
	Red
	Green
)

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	case Silver:
		return "Silver"
	case Red:
		return "Red"
	case Green:
		return "Green"
	default:
		return "White"
	}
}
