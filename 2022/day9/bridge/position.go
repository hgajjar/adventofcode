package bridge

type Position struct {
	Vertical   int
	Horizontal int
}

func (p *Position) HasSameRowOrColumnAs(p2 *Position) bool {
	return p.hasSameRowAs(p2) || p.HasSameColumnAs(p2)
}

func (p *Position) hasSameRowAs(p2 *Position) bool {
	return p.Horizontal == p2.Horizontal
}

func (p *Position) HasSameColumnAs(p2 *Position) bool {
	return p.Vertical == p2.Vertical
}

func (p *Position) IsTouching(p2 *Position) bool {
	return p.isTouchingOnTopLeftWith(p2) ||
		p.isTouchingOnTopRightWith(p2) ||
		p.isTouchingOnBottomLeftWith(p2) ||
		p.isTouchingOnBottomRightWith(p2)
}

func (p *Position) isTouchingOnTopLeftWith(p2 *Position) bool {
	return p.Vertical == p2.Vertical-1 && p.Horizontal == p2.Horizontal-1
}

func (p *Position) isTouchingOnTopRightWith(p2 *Position) bool {
	return p.Vertical == p2.Vertical-1 && p.Horizontal == p2.Horizontal+1
}

func (p *Position) isTouchingOnBottomLeftWith(p2 *Position) bool {
	return p.Vertical == p2.Vertical+1 && p.Horizontal == p2.Horizontal-1
}

func (p *Position) isTouchingOnBottomRightWith(p2 *Position) bool {
	return p.Vertical == p2.Vertical+1 && p.Horizontal == p2.Horizontal+1
}

func (p *Position) IsOnTopLeftOf(p2 *Position) bool {
	return p.Vertical < p2.Vertical && p.Horizontal < p2.Horizontal
}

func (p *Position) IsOnTopRightOf(p2 *Position) bool {
	return p.Vertical < p2.Vertical && p.Horizontal > p2.Horizontal
}

func (p *Position) IsOnBottomLeftOf(p2 *Position) bool {
	return p.Vertical > p2.Vertical && p.Horizontal < p2.Horizontal
}

func (p *Position) IsOnBottomRightOf(p2 *Position) bool {
	return p.Vertical > p2.Vertical && p.Horizontal > p2.Horizontal
}

func UniquePositions(positions []Position) []Position {
	var uniqPositions []Position

loop:
	for _, p1 := range positions {
		for _, p2 := range uniqPositions {
			if p1.Vertical == p2.Vertical && p1.Horizontal == p2.Horizontal {
				continue loop
			}
		}
		uniqPositions = append(uniqPositions, p1)
	}

	return uniqPositions
}
