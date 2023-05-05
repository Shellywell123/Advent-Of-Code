package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) []string {

	// read file into string
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return strings.Split(string(b), "\n")
}

func Part1(filename string) int {

	data := Parse(filename)

	increaseCount := 0

	for i := 0; i < len(data)-1; i++ {

		curLine, err := strconv.Atoi(string(data[i]))
		if err != nil {
			panic(err)
		}

		nextLine, err := strconv.Atoi(string(data[i+1]))
		if err != nil {
			panic(err)
		}

		if nextLine > curLine {
			increaseCount++
		}
	}

	return increaseCount
}

func Part2(filename string) int {

	data := Parse(filename)

	increaseCount := 0

	for i := 0; i < len(data)-3; i++ {

		line1, err := strconv.Atoi(string(data[i]))
		if err != nil {
			panic(err)
		}

		line2, err := strconv.Atoi(string(data[i+1]))
		if err != nil {
			panic(err)
		}

		line3, err := strconv.Atoi(string(data[i+2]))
		if err != nil {
			panic(err)
		}

		line4, err := strconv.Atoi(string(data[i+3]))
		if err != nil {
			panic(err)
		}

		if (line2 + line3 + line4) > (line1 + line2 + line3) {
			increaseCount++
		}
	}

	return increaseCount
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day01")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))
}
