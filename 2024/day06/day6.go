package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Matrix [][]string
type Guard struct {
	x, y int
	dir  byte // ^, >, v, <
}
type Dir struct {
	x, y int
}

func Part1() {
	file, _ := os.Open("day06/input.txt")
	defer file.Close()

	var field Matrix = [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, "")
		field = append(field, splits)
	}

	// Find guard
	guard := findGuard(field)
	if guard == nil {
		panic("Guard not found!")
	}
	fmt.Println("Guard:", guard)

	// Move guard
	var visitedCount = 0
	for {
		// Check borders
		if guard.y == 0 && guard.dir == '^' {
			visitedCount++
			fmt.Println("Guard at top", visitedCount)
			return
		}
		if guard.y == len(field)-1 && guard.dir == 'v' {
			visitedCount++
			fmt.Println("Guard at bottom", visitedCount)
			return
		}
		if guard.x == 0 && guard.dir == '<' {
			visitedCount++
			fmt.Println("Guard at left", visitedCount)
			return
		}
		if guard.x == len(field[0])-1 && guard.dir == '>' {
			visitedCount++
			fmt.Println("Guard at right", visitedCount)
			return
		}

		// Move guard by direction
		switch guard.dir {
		case '^': // Up
			if field[guard.y-1][guard.x] == "#" {
				guard.dir = '>'
			} else {
				guard.y--
				field[guard.y+1][guard.x] = "X"
				if field[guard.y][guard.x] != "X" {
					visitedCount++
				}
				field[guard.y][guard.x] = string(guard.dir)
			}
		case '>': // Right
			if field[guard.y][guard.x+1] == "#" {
				guard.dir = 'v'
			} else {
				guard.x++
				field[guard.y][guard.x-1] = "X"
				if field[guard.y][guard.x] != "X" {
					visitedCount++
				}
				field[guard.y][guard.x] = string(guard.dir)
			}
		case 'v': // Down
			if field[guard.y+1][guard.x] == "#" {
				guard.dir = '<'
			} else {
				guard.y++
				field[guard.y-1][guard.x] = "X"
				if field[guard.y][guard.x] != "X" {
					visitedCount++
				}
				field[guard.y][guard.x] = string(guard.dir)
			}
		case '<': // Left
			if field[guard.y][guard.x-1] == "#" {
				guard.dir = '^'
			} else {
				guard.x--
				field[guard.y][guard.x+1] = "X"
				if field[guard.y][guard.x] != "X" {
					visitedCount++
				}
				field[guard.y][guard.x] = string(guard.dir)
			}
		}
	}
}

func findGuard(field Matrix) *Guard {
	for i := range field {
		for j := range field[0] {
			if field[i][j] == "^" {
				return &Guard{x: j, y: i, dir: '^'}
			}
		}
	}
	return nil
}

func printField(field Matrix) {
	text := ""
	for _, row := range field {
		text += strings.Join(row, "")
		text += "\n"
	}

	output, _ := os.Create("test.txt")
	defer output.Close()
	output.WriteString(text)
	output.Sync()
}
