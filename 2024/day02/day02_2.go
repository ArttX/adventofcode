package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// Check numbers
	var safeCount int = 0
	var invalidSplits [][]string = [][]string{}
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")

		diffs := createDiffList(splits)
		valid := isValidReport(diffs)

		if valid {
			safeCount++
			// fmt.Println(diffs, "Safe")
		} else {
			// fmt.Println(diffs, "Unsafe")
			invalidSplits = append(invalidSplits, splits)
		}
	}

	if len(invalidSplits) > 0 {
		for _, inSplits := range invalidSplits {
			valid := false
			for i := range inSplits {
				copySplit := make([]string, len(inSplits))
				copy(copySplit, inSplits)
				newSplit := remove(copySplit, i)
				newDiff := createDiffList(newSplit)
				valid = isValidReport(newDiff)
				if valid {
					break
				}
			}
			if valid {
				safeCount++
			}
		}
	}

	fmt.Println(safeCount)
}

func createDiffList(splits []string) []int {
	var diffs []int = []int{}
	for i := 0; i < len(splits)-1; i++ {
		num1, _ := strconv.Atoi(splits[i])
		num2, _ := strconv.Atoi(splits[i+1])
		diff := num1 - num2
		diffs = append(diffs, diff)
	}
	return diffs
}

func isValidReport(diffs []int) bool {
	// Check if not changing
	for _, n := range diffs {
		if n == 0 {
			return false
		}
	}

	// Check if all same dir
	isDecreasing := diffs[0] > 0
	for i := 1; i < len(diffs); i++ {
		if isDecreasing != (diffs[i] > 0) {
			return false
		}
	}

	// Check diff amount
	for _, d := range diffs {
		if math.Abs(float64(d)) > 3 {
			return false
		}
	}
	return true
}

func remove[T any](slice []T, i int) []T {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
