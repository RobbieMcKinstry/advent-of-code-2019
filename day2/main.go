package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

	puzzle1(instructions)
	puzzle2(instructions)
}

func puzzle1(instructions []int) {
	instructions[1] = 12
	instructions[2] = 2
	var finalState = executeInstructions(instructions)
	fmt.Println(finalState[0])
}

func executeInstructions(instructions []int) []int {
	var pc = 0
	var ins = instructions[pc]
Loop:
	for {
		fmt.Printf("PC: %v, ins: %v\n", pc, ins)
		switch ins {
		case exit:
			fmt.Printf("Exiting\n")
			break Loop
		case addition:
			var index1 = instructions[pc+1]
			var index2 = instructions[pc+2]
			var index3 = instructions[pc+3]
			var register1 = instructions[index1]
			var register2 = instructions[index2]
			var register3 = register1 + register2
			fmt.Printf("Adding %v + %v = %v\n", register1, register2, register3)
			instructions[index3] = register3
		case multiplication:
			var index1 = instructions[pc+1]
			var index2 = instructions[pc+2]
			var index3 = instructions[pc+3]
			var register1 = instructions[index1]
			var register2 = instructions[index2]
			var register3 = register1 * register2
			fmt.Printf("Multiplying %v * %v = %v\n", register1, register2, register3)
			instructions[index3] = register3
		default:
			panic("Invalid op code")
		}
		pc += 4
		fmt.Printf("Instructions: %v\n", instructions)
		fmt.Printf("PC: %v\n", pc)
		ins = instructions[pc]
	}
	return instructions
}

func puzzle2(instructions []int) {

}
