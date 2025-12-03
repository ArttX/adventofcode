package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}
type VisitPos struct {
	x, y int
	dir  byte
}

func Part2() {
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

	// Obsticle
	var loopObsCount int = 0

	var count = 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if i == guard.x && j == guard.y {
				continue // Skip guard pos
			}
			duplicate := make(Matrix, len(field))
			for d := range field {
				duplicate[d] = make([]string, len(field[d]))
				copy(duplicate[d], field[d])
			}
			duplicate[i][j] = "O"
			visitedCount := moveGuard(*guard, duplicate)
			if visitedCount == -99 {
				// printMatrix(duplicate)
				loopObsCount++
			}
			count++
			log.Printf("%d visitedCount: %v\n", count, visitedCount)
		}
	}

	fmt.Printf("loopObsCount: %v\n", loopObsCount)
}

func moveGuard(guard Guard, field Matrix) int {
	// Move guard
	var visited []VisitPos = []VisitPos{}
	for {
		// Check borders
		if guard.y == 0 && guard.dir == '^' {
			visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
			// fmt.Println("Guard at top", len(visited))
			return len(visited)
		}
		if guard.y == len(field)-1 && guard.dir == 'v' {
			visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
			// fmt.Println("Guard at bottom", len(visited))
			return len(visited)
		}
		if guard.x == 0 && guard.dir == '<' {
			visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
			// fmt.Println("Guard at left", len(visited))
			return len(visited)
		}
		if guard.x == len(field[0])-1 && guard.dir == '>' {
			visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
			// fmt.Println("Guard at right", len(visited))
			return len(visited)
		}

		// Move guard by direction
		switch guard.dir {
		case '^': // Up
			if field[guard.y-1][guard.x] == "#" || field[guard.y-1][guard.x] == "O" {
				guard.dir = '>'
			} else {
				visit := findVisited(visited, guard.x, guard.y-1)
				if field[guard.y-1][guard.x] == "X" && visit != nil && visit.dir == guard.dir {
					return -99
				} else {
					guard.y--
					field[guard.y+1][guard.x] = "X"
					if field[guard.y][guard.x] != "X" {
						visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
					}
					field[guard.y][guard.x] = string(guard.dir)
				}
			}
		case '>': // Right
			if field[guard.y][guard.x+1] == "#" || field[guard.y][guard.x+1] == "O" {
				guard.dir = 'v'
			} else {
				visit := findVisited(visited, guard.x+1, guard.y)
				if field[guard.y][guard.x+1] == "X" && visit != nil && visit.dir == guard.dir {
					return -99
				} else {
					guard.x++
					field[guard.y][guard.x-1] = "X"
					if field[guard.y][guard.x] != "X" {
						visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
					}
					field[guard.y][guard.x] = string(guard.dir)
				}
			}
		case 'v': // Down
			if field[guard.y+1][guard.x] == "#" || field[guard.y+1][guard.x] == "O" {
				guard.dir = '<'
			} else {
				visit := findVisited(visited, guard.x, guard.y+1)
				if field[guard.y+1][guard.x] == "X" && visit != nil && visit.dir == guard.dir {
					return -99
				} else {
					guard.y++
					field[guard.y-1][guard.x] = "X"
					if field[guard.y][guard.x] != "X" {
						visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
					}
					field[guard.y][guard.x] = string(guard.dir)
				}
			}
		case '<': // Left
			if field[guard.y][guard.x-1] == "#" || field[guard.y][guard.x-1] == "O" {
				guard.dir = '^'
			} else {
				visit := findVisited(visited, guard.x-1, guard.y)
				if field[guard.y][guard.x-1] == "X" && visit != nil && visit.dir == guard.dir {
					return -99
				} else {
					guard.x--
					field[guard.y][guard.x+1] = "X"
					if field[guard.y][guard.x] != "X" {
						visited = append(visited, VisitPos{guard.x, guard.y, guard.dir})
					}
					field[guard.y][guard.x] = string(guard.dir)
				}
			}
		}
	}
}

func printMatrix(matrix Matrix) {
	text := ""
	for _, line := range matrix {
		text += strings.Join(line, "")
		text += "\n"
	}
	fmt.Println(text)
}

func findVisited(visited []VisitPos, x, y int) *VisitPos {
	for _, visit := range visited {
		if visit.x == x && visit.y == y {
			return &visit
		}
	}
	return nil
}
