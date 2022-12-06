package shape

const (
	rock     = "rock"
	paper    = "paper"
	scissors = "scissors"
)

type Shape struct {
	name  string
	score int
}

func NewRock() Shape {
	return Shape{
		name:  rock,
		score: 1,
	}
}

func NewPaper() Shape {
	return Shape{
		name:  paper,
		score: 2,
	}
}

func NewScissors() Shape {
	return Shape{
		name:  scissors,
		score: 3,
	}
}

func (s Shape) IsRock() bool {
	return s.name == rock
}

func (s Shape) IsPaper() bool {
	return s.name == paper
}

func (s Shape) IsScissors() bool {
	return s.name == scissors
}

func (s Shape) Score() int {
	return s.score
}
