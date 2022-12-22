package day9

import (
	"adventofcode/day9/bridge"
	_ "embed"
	"fmt"
)

var (
	//go:embed input/sample.txt
	sampleInput string

	//go:embed input/actual.txt
	actualInput string
)

func Execute() {
	fmt.Println(countTotalNumberOfTailPositionsChanged(sampleInput))
	fmt.Println(countTotalNumberOfTailPositionsChanged(actualInput))
}

func countTotalNumberOfTailPositionsChanged(input string) int {
	motionGuide := bridge.NewMotionGuide(input)

	head := bridge.Position{}
	tail := head

	tailPositions := []bridge.Position{tail}

	for _, m := range motionGuide {
		switch m.Direction {
		case "L":
			tailPositions = append(tailPositions, runSteps(&head, &tail, m.Steps, func(p *bridge.Position) {
				p.Horizontal--
			})...)
		case "R":
			tailPositions = append(tailPositions, runSteps(&head, &tail, m.Steps, func(p *bridge.Position) {
				p.Horizontal++
			})...)
		case "U":
			tailPositions = append(tailPositions, runSteps(&head, &tail, m.Steps, func(p *bridge.Position) {
				p.Vertical--
			})...)
		case "D":
			tailPositions = append(tailPositions, runSteps(&head, &tail, m.Steps, func(p *bridge.Position) {
				p.Vertical++
			})...)
		}
	}

	return len(bridge.UniquePositions(tailPositions))
}

func runSteps(head, tail *bridge.Position, steps int, calculateHeadPosition func(p *bridge.Position)) []bridge.Position {
	var tailPositions []bridge.Position
	for i := 0; i < steps; i++ {
		oldTail := *tail

		calculateHeadPosition(head)
		calculateTailPosition(tail, head)

		if tail.Horizontal != oldTail.Horizontal || tail.Vertical != oldTail.Vertical {
			tailPositions = append(tailPositions, *tail)
		}
	}

	return tailPositions
}

func calculateTailPosition(tail, head *bridge.Position) {
	if tail.HasSameRowOrColumnAs(head) {
		if !tail.HasSameColumnAs(head) {
			diff := head.Vertical - tail.Vertical
			if diff > 1 {
				tail.Vertical++
			} else if diff < -1 {
				tail.Vertical--
			}
		} else {
			diff := head.Horizontal - tail.Horizontal
			if diff > 1 {
				tail.Horizontal++
			} else if diff < -1 {
				tail.Horizontal--
			}
		}
	} else {
		if !tail.IsTouching(head) {
			if head.IsOnTopLeftOf(tail) {
				tail.Vertical--
				tail.Horizontal--
			} else if head.IsOnTopRightOf(tail) {
				tail.Vertical--
				tail.Horizontal++
			} else if head.IsOnBottomLeftOf(tail) {
				tail.Vertical++
				tail.Horizontal--
			} else if head.IsOnBottomRightOf(tail) {
				tail.Vertical++
				tail.Horizontal++
			}
		}
	}
}
