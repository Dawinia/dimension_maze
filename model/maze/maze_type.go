package maze

type MazeType int

const (
	Empty     MazeType = iota // 空
	Wall                      // 墙
	WrapPoint                 // 翘曲点
	Entrance                  // 入口
	Export                    // 出口
)

func (mType MazeType) GetColor() Color {
	switch mType {
	case Empty:
		return White
	case Wall:
		return Black
	case WrapPoint:
		return Silver
	case Entrance:
		return Red
	case Export:
		return Green
	default:
		return White
	}
}
