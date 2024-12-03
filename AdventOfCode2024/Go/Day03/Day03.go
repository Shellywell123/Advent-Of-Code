package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	instructions := [][]int{}

	for fscanner.Scan() {
		line := fscanner.Text()

		// regex for "mul(number,number)" (this took way too long lol)
		re := regexp.MustCompile("mul\\([0-9]?[0-9]?[0-9],[0-9]?[0-9]?[0-9]\\)")
		match := re.FindAllStringSubmatch(line, -1)

		// extract number pairs from the regex result
		for _, x := range match {

			// trim brackets
			trim_x := x[0][4 : len(x[0])-1]

			// split by comma
			pair := []int{}
			for _, number := range strings.Split(trim_x, ",") {
				n, _ := strconv.Atoi(string(number))
				pair = append(pair, n)
			}

			instructions = append(instructions, pair)
		}
	}

	return instructions
}

func parse2(filename string) [][]int {

	instructions := [][]int{}

	// reading in whole file to ignore newline chars
	contentBytes, _ := os.ReadFile(filename)
	content := string(contentBytes)

	// regex for "mul(number,number)" (this took way too long lol)
	re := regexp.MustCompile("do\\(\\)|don't\\(\\)|mul\\([0-9]?[0-9]?[0-9],[0-9]?[0-9]?[0-9]\\)")
	match := re.FindAllStringSubmatch(content, -1)

	enabled := true
	// extract number pairs from the regex result
	for _, x := range match {

		if x[0] == "do()" {
			enabled = true
			continue
		}
		if x[0] == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}

		// trim brackets
		trim_x := x[0][4 : len(x[0])-1]

		// split by comma
		pair := []int{}
		for _, number := range strings.Split(trim_x, ",") {
			n, _ := strconv.Atoi(string(number))
			pair = append(pair, n)
		}

		instructions = append(instructions, pair)

	}

	return instructions
}

func Part1(filename string) int {
	instructions := parse(filename)

	ans := 0
	for _, instruction := range instructions {
		ans += (instruction[0] * instruction[1])
	}

	return ans
}

func Part2(filename string) int {
	instructions := parse2(filename)

	ans := 0
	for _, instruction := range instructions {
		ans += (instruction[0] * instruction[1])
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day03")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests2.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
