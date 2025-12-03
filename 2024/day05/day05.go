package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, _ := os.Open("day05/input.txt")
	scanner := bufio.NewScanner(file)

	var rules [][]string = [][]string{}
	readingRules := true
	var validReadings [][]string = [][]string{}

	// Reading data
	for scanner.Scan() {
		line := scanner.Text()
		if readingRules {
			if line == "" {
				readingRules = false
			} else {
				rule := strings.Split(line, "|")
				rules = append(rules, []string{rule[0], rule[1]})
			}
		} else {
			list := strings.Split(line, ",")
			// Do check
			allValid := true
			for i := range list {
				for j := i + 1; j < len(list); j++ {
					first := list[i]
					second := list[j]
					foundValid := false
					for _, rule := range rules {
						if rule[0] != first {
							continue
						}
						if rule[1] == second {
							foundValid = true
							break
						}
					}
					if !foundValid {
						allValid = false
					}
				}
			}
			// fmt.Println(list, allValid)
			if allValid {
				validReadings = append(validReadings, list)
			}
		}
	}

	// fmt.Println("Valid:", validReadings)

	// Sum list medians
	sum := 0
	for _, item := range validReadings {
		size := len(item)
		idx := size / 2
		num, _ := strconv.Atoi(item[idx])
		sum += num
	}
	fmt.Println("Sum:", sum)
}

func Part2() {
	file, _ := os.Open("day05/input.txt")
	scanner := bufio.NewScanner(file)

	var rules [][]string = [][]string{}
	readingRules := true
	var invalidReadings [][]string = [][]string{}

	// Reading data
	for scanner.Scan() {
		line := scanner.Text()
		if readingRules {
			if line == "" {
				readingRules = false
			} else {
				rule := strings.Split(line, "|")
				rules = append(rules, []string{rule[0], rule[1]})
			}
		} else {
			list := strings.Split(line, ",")
			// Do check
			allValid := true
			for i := range list {
				for j := i + 1; j < len(list); j++ {
					first := list[i]
					second := list[j]
					foundValid := false
					for _, rule := range rules {
						if rule[0] != first {
							continue
						}
						if rule[1] == second {
							foundValid = true
							break
						}
					}
					if !foundValid {
						allValid = false
						list[i], list[j] = list[j], list[i] // swap values
					}
				}
			}
			// fmt.Println(list, allValid)
			if !allValid {
				invalidReadings = append(invalidReadings, list)
			}
		}
	}

	// fmt.Println("Invalid:", invalidReadings)

	// Sum list medians
	sum := 0
	for _, item := range invalidReadings {
		size := len(item)
		idx := size / 2
		num, _ := strconv.Atoi(item[idx])
		sum += num
	}
	fmt.Println("Sum:", sum)
}
