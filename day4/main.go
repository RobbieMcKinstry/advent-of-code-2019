package main

import (
	"fmt"
	"sync"
)

const (
	lowRange  = 245318
	highRange = 765747
	start     = 100000
	end       = 1000000
)

func main() {

	puzzle1()
	puzzle2()
}

func puzzle1() {
	countNumbers(isWithinRange, hasAdjacentDigits, hasMonotoneDigits)
}

func puzzle2() {
	countNumbers(isWithinRange, hasExactly2AdjacentDigits, hasMonotoneDigits)
}

func countNumbers(validators ...requirement) {
	var solutions = make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(end - start)

	go func() {
		wg.Wait()
		close(solutions)
	}()

	for i := start; i < end; i++ {
		go validate(i, solutions, &wg, validators...)
	}

	var count int
	for range solutions {
		count++
	}
	fmt.Println(count)
}

type requirement func(num int, digits []uint8) bool

func validate(num int, solution chan<- int, wg *sync.WaitGroup, validators ...requirement) {
	defer wg.Done()
	var (
		digits    = getDigits(num)
		satisfied = true
	)

	for _, v := range validators {
		satisfied = satisfied && v(num, digits)
	}

	if satisfied {
		solution <- num
	}
}

func hasAdjacentDigits(_ int, digits []uint8) bool {
	var hasAdjacentDigits bool
	for i := 1; i < len(digits); i++ {
		left, right := digits[i-1], digits[i]
		hasAdjacentDigits = hasAdjacentDigits || left == right
	}
	return hasAdjacentDigits
}

func hasExactly2AdjacentDigits(_ int, digits []uint8) bool {
	for i := 1; i < len(digits); i++ {
		two, three := digits[i-1], digits[i]

		if two != three {
			// Make sure they match
			continue
		} else if i == 1 {
			// Only check above
			four := digits[i+1]
			if four == two {
				continue
			}
		} else if i == len(digits)-1 {
			// Only check below
			one := digits[i-2]
			if one == two {
				continue
			}
		} else {
			one, four := digits[i-2], digits[i+1]
			if one == two || four == two {
				continue
			}
		}

		return true
	}
	return false
}

func hasMonotoneDigits(_ int, digits []uint8) bool {
	var hasMonotoneDigits bool = true
	for i := 1; i < len(digits); i++ {
		left, right := digits[i-1], digits[i]
		hasMonotoneDigits = hasMonotoneDigits && left <= right
	}
	return hasMonotoneDigits
}

func getDigits(num int) []uint8 {
	var digits = []uint8{}
	for num > 0 {
		var next = num % 10
		digits = append(digits, uint8(next))
		num /= 10
	}

	// Reverse the slice
	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		digits[left], digits[right] = digits[right], digits[left]
	}

	return digits
}

func isWithinRange(num int, _ []uint8) bool {
	return lowRange <= num && num <= highRange
}
