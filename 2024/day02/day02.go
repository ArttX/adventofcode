package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// Check numbers
	var safeCount int = 0
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")
		var valid bool = true
		var diffs []int = []int{}
		// Make diff list
		for i := 0; i < len(splits)-1; i++ {
			num1, _ := strconv.Atoi(splits[i])
			num2, _ := strconv.Atoi(splits[i+1])
			diff := num1 - num2
			diffs = append(diffs, diff)
		}

		// Check if not changing
		for _, n := range diffs {
			if n == 0 {
				valid = false
				break
			}
		}

		// Check if all same dir
		isDecreasing := diffs[0] > 0
		for i := 1; i < len(diffs); i++ {
			if isDecreasing != (diffs[i] > 0) {
				valid = false
				break
			}
		}

		// Check diff amount
		for _, d := range diffs {
			if math.Abs(float64(d)) > 3 {
				valid = false
				break
			}
		}

		if valid {
			safeCount++
			// fmt.Println(diffs, "Safe")
		} else {
			// fmt.Println(diffs, "Unsafe")
		}
	}
	fmt.Println(safeCount)
}
