package model

import "fmt"

type Role struct {
	Pos pos
	Direction
}

type pos [2]int

func (pos *pos) move(b pos) {
	pos[0] += b[0]
	pos[1] += b[1]
}

func (pos pos) EqualTo(b pos) bool {
	return pos[0] == b[0] && pos[1] == b[1]
}

func (role Role) String() string {
	return fmt.Sprintf("role.pos: %v role.direction: %v", role.Pos, role.Direction)
}

type Direction int

const (
	East Direction = iota
	South
	West
	North
)

func (d Direction) get() [2]int {
	switch d {
	case North:
		return [2]int{-1, 0}
	case East:
		return [2]int{0, 1}
	case South:
		return [2]int{1, 0}
	case West:
		return [2]int{0, -1}
	default:
		return [2]int{0, 0}
	}
}


func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Stable"
	}
}

func (role *Role) Move() {
	dire := role.Direction.get()
	role.Pos.move(dire)
}

func (role *Role) Rotate(clockwise bool) {
	if clockwise {
		role.Direction++
	} else {
		role.Direction--
	}
	role.Direction = (role.Direction + 4) % 4
}

func (role *Role) CurState()  {

}