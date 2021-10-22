package generate

import (
	"dimension_maze/model"
	"math/rand"
)

func Generate(height, width int, wpCount int) [][]model.Maze {
	maze := make([][]model.Maze, 0, height)
	for i := 0; i < height; i++ {
		mazeRow := make([]model.Maze, width)
		for j := 0; j < width; j++ {
			mazeType := model.MazeType(rand.Intn(5))
			duration := 0
			if mazeType == model.WrapPoint {
				duration = rand.Intn(30)
			}
			mazeRow = append(mazeRow, model.Maze{
				MazeType: mazeType,
				Duration: duration,
				MazeColor: mazeType.GetColor(),
			})
		}
		maze = append(maze, mazeRow)
	}
	return maze
}