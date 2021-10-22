package main

import (
	"dimension_maze/model/role"
	"fmt"
)

func main() {
	//maze := model.Maze{
	//	Duration: 10,
	//	MazeType: model.MazeType(1),
	//	MazeColor: model.Color(3),
	//}
	//fmt.Println(maze)
	//fmt.Println(maze.MazeColor == model.White)

	role := role.Role{
		Pos:       [2]int{1, 1},
		Direction: role.East,
	}
	for i := 0; i < 10; i++ {
		showForTest(role)
		role.Move()
		role.Rotate(false)
		fmt.Println("--------------------------")
	}
	//fmt.Println()
	//maze := generate.Generate(3, 3, 1)
	//model.ShowMaze(maze)
}

func showForTest(role role.Role) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if role.Pos.EqualTo([2]int{i, j}) {
				fmt.Printf("X\t")
			} else {
				fmt.Printf("O\t")
			}
		}
		fmt.Println()
	}
}
