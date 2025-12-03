package day04

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() {
	file, _ := os.Open("day04/input.txt")
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lines = append(lines, line)
	}
	var count int
	for i, line := range lines {
		for j, c := range line {
			if c == 'X' {
				// forward
				if j < (len(lines)-3) && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
					count++
				}

				// reverse
				if j > 2 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S' {
					count++
				}

				// down
				if i < (len(lines)-3) && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					count++
				}
				// up
				if i > 2 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					count++
				}

				// forward - down
				if i < (len(lines)-3) && j < (len(lines)-3) && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					count++
				}

				// forward - up
				if i > 2 && j < (len(lines)-3) && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					count++
				}

				// reverse - down
				if i < (len(lines)-3) && j > 2 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					count++
				}

				// reverse - up
				if i > 2 && j > 2 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					count++
				}

			}
		}
	}

	fmt.Printf("count: %v\n", count)
}

func Part2() {
	file, _ := os.Open("day04/input.txt")
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lines = append(lines, line)
	}
	var count int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				diag0 := lines[i-1][j-1]|lines[i+1][j+1] == 'S'|'M'
				diag1 := lines[i+1][j-1]|lines[i-1][j+1] == 'S'|'M'
				if diag0 && diag1 {
					count++
				}
			}
		}
	}

	fmt.Printf("count: %v\n", count)
}
