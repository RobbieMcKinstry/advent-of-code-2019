package main

import (
	"testing"
)

func TestPuzzle1(t *testing.T) {
	instructions := [][]string{
		{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	}
	const expected = 159
	observed := findOverlap(instructions)
	if observed != expected {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}

func TestPuzzle2(t *testing.T) {
	instructions := [][]string{
		{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	}
	const expected = 610
	observed := findMinStep(instructions)
	if observed != expected {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}

func TestPuzzle3(t *testing.T) {
	instructions := [][]string{
      {"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"},
      {"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"},
	}
	const expected = 410
	observed := findMinStep(instructions)
	if observed != expected {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}
