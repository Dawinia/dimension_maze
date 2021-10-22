package maze

import "fmt"

type Maze struct {
	MazeType
	Duration  int
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
