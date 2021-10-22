package model

import "fmt"

type MazeType int

const (
	Empty MazeType = iota	// 空
	Wall					// 墙
	WrapPoint				// 翘曲点
	Entrance				// 入口
	Export					// 出口
)

type Color int

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

const (
	White Color = iota
	Black
	Silver
	Red
	Green
)

type Maze struct {
	MazeType
	Duration int
	MazeColor Color
}

func (m Maze) String() string {
	return fmt.Sprintf("maze.type: %v maze.duration: %v maze.color: %v", m.MazeType, m.Duration, m.MazeColor)
}

func ShowMaze(maze [][]Maze) {
	for _, mazeRow := range maze {
		for _, mazeCell := range mazeRow {
			fmt.Println(mazeCell, " ")
		}
		fmt.Println()
	}
}

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