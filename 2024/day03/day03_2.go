package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|don't\\(\\)|do\\(\\)")
	dr := regexp.MustCompile("\\d{1,3}")

	scanner := bufio.NewScanner(file)
	var result int = 0
	enabled := true
	for scanner.Scan() {
		tokens := r.FindAllString(scanner.Text(), -1)

		for _, t := range tokens {
			if enabled {
				if t == "don't()" {
					enabled = false
					continue
				} else if t == "do()" {
					enabled = true
					continue
				} else {
					nums := dr.FindAllString(t, -1)
					num1, _ := strconv.Atoi(nums[0])
					num2, _ := strconv.Atoi(nums[1])
					result += num1 * num2
				}
			} else {
				if t == "do()" {
					enabled = true
					continue
				}
			}
		}
	}

	fmt.Println(result)
}
