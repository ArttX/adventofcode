package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, _ := os.Open("day07/test_input.txt")
	defer file.Close()

	// result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		result, _ := strconv.Atoi(splits[0])
		numberList := strings.Split(strings.Trim(splits[1], " "), " ")
		fmt.Println(result, numberList[:])

		numbers := []int{}
		for i := range numberList {
			num, _ := strconv.Atoi(numberList[i])
			numbers = append(numbers, num)
		}

		sum := 0
		for _, n := range numbers {
			sum += n
		}
		fmt.Println(result, sum)
	}
}
