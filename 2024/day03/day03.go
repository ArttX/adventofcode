package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	dr := regexp.MustCompile("\\d{1,3}")

	scanner := bufio.NewScanner(file)
	var result int = 0
	for scanner.Scan() {
		matches := r.FindAllString(scanner.Text(), -1)
		for _, m := range matches {
			nums := dr.FindAllString(m, -1)
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			result += num1 * num2
		}
	}
	fmt.Println(result)
}
