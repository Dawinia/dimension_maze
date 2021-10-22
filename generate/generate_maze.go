package generate

import (
	"dimension_maze/model/maze"
	"math/rand"
)

func Generate(height, width int, wpCount int) [][]maze.Maze {
	newMaze := make([][]maze.Maze, 0, height)
	for i := 0; i < height; i++ {
		mazeRow := make([]maze.Maze, width)
		for j := 0; j < width; j++ {
			mazeType := maze.MazeType(rand.Intn(5))
			duration := 0
			if mazeType == maze.WrapPoint {
				duration = rand.Intn(30)
			}
			mazeRow = append(mazeRow, maze.Maze{
				MazeType:  mazeType,
				Duration:  duration,
				MazeColor: mazeType.GetColor(),
			})
		}
		newMaze = append(newMaze, mazeRow)
	}
	return newMaze
}
