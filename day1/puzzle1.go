package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const input1 = "input1.txt"

func main() {
	file, err := os.Open(input1)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var lines = []int{}
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
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

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []int) {
	calculateFuel(lines)
}

func calculateFuel(lines []int) {
	mapped := make([]int, 0, len(lines))
	for _, line := range lines {
		mapped = append(mapped, (line/3)-2)
	}
	total := 0
	for _, m := range mapped {
		total += m
	}
	fmt.Println(total)
}

func puzzle2(lines []int) {
	calculateFuelRec(lines)
}

func calculateFuelRec(lines []int) {
	total := 0
	fuelCounts := make(chan int, len(lines))
	for _, line := range lines {
		go calculateFuelInParallel(line, fuelCounts)
	}

	for range lines {
		total += <-fuelCounts
	}

	fmt.Println(total)
}

func calculateFuelInParallel(module int, fuelStream chan<- int) {
	fuel := fuelForModule(module)
	fuelStream <- fuel
}

func fuelForModule(module int) int {
	fuel := module/3 - 2
	if fuel <= 0 {
		return max(0, fuel)
	}
	return fuel + fuelForModule(fuel)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
