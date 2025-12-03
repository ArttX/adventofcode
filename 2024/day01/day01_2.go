package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)

	// Convert to 2 lists
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")
		n1, _ := strconv.ParseInt(splits[0], 10, 0)
		n2, _ := strconv.ParseInt(splits[3], 10, 0)
		list1 = append(list1, int(n1))
		list2 = append(list2, int(n2))
	}

	var result int
	for _, num := range list1 {
		var count int = 0
		for _, num2 := range list2 {
			if num == num2 {
				count++
			}
		}
		result += num * count
	}

	fmt.Println(result)
}
