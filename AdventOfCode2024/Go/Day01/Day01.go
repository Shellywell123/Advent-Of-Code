package main

import (
	"bufio"
	"math"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Parse(filename string) ([]int, []int) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}

	for fscanner.Scan() {
		line := fscanner.Text()

		line_numbers := []int{}
		for _, number := range strings.Split(line, "   ") {
			n, _ := strconv.Atoi(string(number))
			line_numbers = append(line_numbers, n)
		}
		list1 = append(list1, line_numbers[0])
		list2 = append(list2, line_numbers[1])
	}

	return list1, list2
}

func Part1(filename string) int {
	list1, list2 := Parse(filename)

	// sort list into asc order
	sort.Ints(list1)
	sort.Ints(list2)

	// total the difference between each element
	ans := 0
	for i := 0; i < len(list1); i++ {
		ans += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}
	return ans
}

func Part2(filename string) int {
	list1, list2 := Parse(filename)

	// total similarity (element in list1 * count of element in list2)
	ans := 0
	for i := 0; i < len(list1); i++ {

		count := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				count += 1
			}
		}
		ans += list1[i] * count
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day10")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
