package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const input = "input.txt"

func main() {
	var (
		file, err        = os.Open(input)
		lines            = []string{}
		lineInstructions = [][]string{}
		scanner          = bufio.NewScanner(file)
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, line := range lines {
		lineInstruction := strings.Split(strings.TrimSpace(string(line)), ",")
		lineInstructions = append(lineInstructions, lineInstruction)
	}

	puzzle1(lineInstructions)
	puzzle2(lineInstructions)
}

type Coordinate struct {
	X, Y int
	step int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func Origin() Coordinate {
	return Coordinate{0, 0, 0}
}

func (c1 Coordinate) ManhattanDistance(c2 Coordinate) int {
	var xDist = c1.X - c2.X
	var yDist = c1.Y - c2.Y
	if xDist < 0 {
		xDist = -xDist
	}
	if yDist < 0 {
		yDist = -yDist
	}
	return xDist + yDist
}

func (c Coordinate) MoveLeft(x int) Coordinate {
	return Coordinate{
		X:    c.X - x,
		Y:    c.Y,
		step: c.step + x,
	}
}

func (c Coordinate) MoveRight(x int) Coordinate {
	return Coordinate{
		X:    c.X + x,
		Y:    c.Y,
		step: c.step + x,
	}
}

func (c Coordinate) MoveUp(y int) Coordinate {
	return Coordinate{
		X:    c.X,
		Y:    c.Y + y,
		step: c.step + y,
	}
}

func (c Coordinate) MoveDown(y int) Coordinate {
	return Coordinate{
		X:    c.X,
		Y:    c.Y - y,
		step: c.step + y,
	}
}

type Set map[string]Coordinate

func (s Set) Add(c Coordinate) {
	if _, ok := s[c.String()]; !ok {
		s[c.String()] = c
	}
}

func (set Set) Intersection(other Set) Set {
	intersection := Set{}

	for key, elem1 := range other {
		if elem2, ok := set[key]; ok {
			sumElem := Coordinate{}
			sumElem.X = elem1.X
			sumElem.Y = elem1.Y
			sumElem.step = elem1.step + elem2.step
			fmt.Printf("After intersection, step = %v\n", sumElem.step)
			intersection.Add(sumElem)
		}
	}
	return intersection
}

func (set Set) Delete(co Coordinate) {
	if _, ok := set[co.String()]; ok {
		delete(set, co.String())
	}
}

// This function adds all coordinates between these two coordinates, inclusive
// These coordinates must be on the same horizontal line or vertical line.
func (s Set) AddJoiningCoordinates(first, second Coordinate) {
	fmt.Printf("first=%v, second=%v\n", first.step, second.step)
	s.Add(first)
	s.Add(second)
	step := first.step + 1
	switch {
	case first.X < second.X:
		for i := first.X; i <= second.X; i++ {
			s.Add(Coordinate{X: i, Y: first.Y, step: step})
			step++
		}
	case second.X < first.X:
		for i := second.X; i <= first.X; i++ {
			s.Add(Coordinate{X: i, Y: first.Y, step: step})
			step++
		}
	case first.Y < second.Y:
		for i := first.Y; i <= second.Y; i++ {
			s.Add(Coordinate{X: first.X, Y: i, step: step})
			step++
		}
	case second.Y < first.Y:
		for i := second.Y; i <= first.Y; i++ {
			s.Add(Coordinate{X: first.X, Y: i, step: step})
			step++
		}
	case first.X == second.X && first.Y == second.Y:
		break
	default:
		panic("Not possible")
	}
}

func puzzle1(lineInstructions [][]string) {
	fmt.Println(findOverlap(lineInstructions))
}

func findOverlap(lineInstructions [][]string) int {
	var (
		leftLine      = lineInstructions[0]
		rightLine     = lineInstructions[1]
		leftReceiver  = make(chan Set, 1)
		rightReceiver = make(chan Set, 1)
	)
	go buildCoordinates(leftLine, leftReceiver)
	go buildCoordinates(rightLine, rightReceiver)
	var leftCoordinates = <-leftReceiver
	var rightCoordinates = <-rightReceiver
	var intersection = leftCoordinates.Intersection(rightCoordinates)
	intersection.Delete(Origin())
	var smallest = math.MaxInt32
	for _, p := range intersection {
		if dist := p.ManhattanDistance(Origin()); dist < smallest {
			smallest = dist
		}
	}
	return smallest
}

func buildCoordinates(instructions []string, sender chan<- Set) {
	var visited = make(Set)
	var position = Origin()
	for _, instruction := range instructions {
		var count, err = strconv.Atoi(instruction[1:])
		var nextPosition Coordinate
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch instruction[0] {
		case 'U':
			nextPosition = position.MoveUp(count)
		case 'D':
			nextPosition = position.MoveDown(count)
		case 'L':
			nextPosition = position.MoveLeft(count)
		case 'R':
			nextPosition = position.MoveRight(count)
		}
		// fmt.Printf("Joining… pos=%v, next=%v\n", position.step, nextPosition.step)
		visited.AddJoiningCoordinates(position, nextPosition)
		position = nextPosition
	}
	sender <- visited
}

func findMinStep(lineInstructions [][]string) int {
	var (
		leftLine      = lineInstructions[0]
		rightLine     = lineInstructions[1]
		leftReceiver  = make(chan Set, 1)
		rightReceiver = make(chan Set, 1)
	)
	go buildCoordinates(leftLine, leftReceiver)
	go buildCoordinates(rightLine, rightReceiver)
	var leftCoordinates = <-leftReceiver
	var rightCoordinates = <-rightReceiver
	var intersection = leftCoordinates.Intersection(rightCoordinates)
	intersection.Delete(Origin())
	var smallest = math.MaxInt32
	for _, p := range intersection {
		if step := p.step; step < smallest {
			smallest = step
		}
	}
	return smallest
}

func puzzle2(lineInstructions [][]string) {
	fmt.Println(findMinStep(lineInstructions))
}
