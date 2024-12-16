package puzzle

type Robot struct {
	Pos []int
	Vel []int
}

type Robots []Robot

func NewRobot(p1, p2, v1, v2 string) Robot {
	return Robot{
		Pos: []int{getNumber(p2), getNumber(p1)},
		Vel: []int{getNumber(v2), getNumber(v1)},
	}
}

func (r Robot) Move() {
	r.Pos[0] += r.Vel[0]
	r.Pos[1] += r.Vel[1]
}
