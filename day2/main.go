package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const target = 19690720
const input = "input.txt"
const (
	addition = iota + 1
	multiplication
	exit = 99
)

func main() {
	var (
		file, err            = os.Open(input)
		bytes                []byte
		instructionsAsString []string
		instructions         []int
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instructionsAsString = strings.Split(strings.TrimSpace(string(bytes)), ",")
	for _, inst := range instructionsAsString {
		next, err := strconv.Atoi(inst)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		instructions = append(instructions, next)
	}

	var copiedInstructions = append([]int{}, instructions...)
	puzzle1(copiedInstructions)
	puzzle2(instructions)
}

func puzzle1(instructions []int) {
	instructions[1] = 12
	instructions[2] = 2
	var finalState, err = executeInstructions(instructions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(finalState[0])
}

func executeInstructions(instructions []int) ([]int, error) {
	var pc = 0
	var ins = instructions[pc]
Loop:
	for {
		switch ins {
		case exit:
			break Loop
		case addition:
			if len(instructions) <= pc+3 {
				return nil, errors.New("Instruction out of bounds")
			}
			var index1 = instructions[pc+1]
			var index2 = instructions[pc+2]
			var index3 = instructions[pc+3]
			var register1 = instructions[index1]
			var register2 = instructions[index2]
			var register3 = register1 + register2
			instructions[index3] = register3
		case multiplication:
			if len(instructions) <= pc+3 {
				return nil, errors.New("Instruction out of bounds")
			}
			var index1 = instructions[pc+1]
			var index2 = instructions[pc+2]
			var index3 = instructions[pc+3]
			var register1 = instructions[index1]
			var register2 = instructions[index2]
			var register3 = register1 * register2
			instructions[index3] = register3
		default:
			return nil, errors.New("No such opcode")
		}
		pc += 4
		if len(instructions) <= pc {
			return nil, errors.New("Instruction out of bounds")
		}
		ins = instructions[pc]
	}
	return instructions, nil
}

// A Pair is a ...
type Pair struct {
	Noun, Verb, Output int
	Err                error
}

func puzzle2(instructions []int) {

	var (
		leftBound    = 100
		rightBound   = 100
		totalSize    = leftBound * rightBound
		resultStream = make(chan Pair, totalSize)
	)

	for i := 0; i < leftBound; i++ {
		for j := 0; j < rightBound; j++ {
			var copiedInstructions = append([]int(nil), instructions...)
			go executeProgram(copiedInstructions, i, j, resultStream)
		}
	}

	counter := 0
	for result := range resultStream {
		if result.Err == nil && result.Output == target {
			fmt.Println(result.Noun*100 + result.Verb)
			os.Exit(0)
		}

		counter++
		if counter >= totalSize {
			fmt.Println("No correct result was found")
			os.Exit(1)
		}
	}
}

func executeProgram(instructions []int, leftBound, rightBound int, resultStream chan<- Pair) {
	instructions[1] = leftBound
	instructions[2] = rightBound
	var result, err = executeInstructions(instructions)
	if err != nil {
		resultStream <- Pair{Noun: leftBound, Verb: rightBound, Output: 0, Err: err}
	} else {
		resultStream <- Pair{Noun: leftBound, Verb: rightBound, Output: result[0], Err: err}
	}
}
