package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	test   int
	inputs []int
}

func parse(filename string) []operation {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	ops := []operation{}
	for fscanner.Scan() {
		line := fscanner.Text()

		test, _ := strconv.Atoi(strings.Split(line, ": ")[0])

		inputs := []int{}
		for _, number := range strings.Split(strings.Split(line, ": ")[1], " ") {
			n, _ := strconv.Atoi(string(number))
			inputs = append(inputs, n)
		}
		ops = append(ops, operation{test, inputs})
	}

	return ops
}

func concatInt(a, b int) int {
	s := strconv.Itoa(a) + strconv.Itoa(b)
	i, _ := strconv.Atoi(s)
	return i
}

func operationValid(test, result int, inputs []int, part int) bool {

	// guard cases
	if len(inputs) == 0 && result == test {
		return true
	}
	if len(inputs) == 0 && result != test {
		return false
	}
	if result > test {
		return false
	}

	// recursive cases
	if operationValid(test, result+inputs[0], inputs[1:], part) || operationValid(test, result*inputs[0], inputs[1:], part) {
		return true
	}
	if part == 2 {
		if operationValid(test, concatInt(result, inputs[0]), inputs[1:], part) {
			return true
		}
	}
	return false

}

func Part1(filename string, debug bool) int {
	ops := parse(filename)

	ans := 0
	for _, op := range ops {

		if debug {
			fmt.Printf("op: %v\n", op)
		}

		if operationValid(op.test, 0, op.inputs, 1) {
			ans += op.test
		}
	}
	return ans
}

func Part2(filename string, debug bool) int {
	ops := parse(filename)

	ans := 0
	for _, op := range ops {

		if operationValid(op.test, 0, op.inputs, 2) {
			if debug {
				fmt.Printf("op: %v\n", op)
			}
			ans += op.test
		}
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day07")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
