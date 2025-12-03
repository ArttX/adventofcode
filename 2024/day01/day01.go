package day01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
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

	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	var result int
	for i := 0; i < len(list1); i++ {
		result += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println(result)
}
