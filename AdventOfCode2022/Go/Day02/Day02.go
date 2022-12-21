package main

import (
	"bufio"
	"fmt"
	"os"
)

func Part1(filename string) int {
	var totals int
	var index int
	var line string

	outcomes := [9]string{"B X", "C Y", "A Z", "A X", "B Y", "C Z", "C X", "A Y", "B Z"}

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	totals = 0
	for fscanner.Scan() {
		line = fscanner.Text()

		// get index from list
		for i := 0; i < len(outcomes); i++ {
			if outcomes[i] == line {
				index = i + 1
				break
			}
		}
		totals += index
	}

	return totals
}

func Part2(filename string) int {
	var totals int
	var index int
	var line string

	outcomes := [9]string{"B X", "C X", "A X", "A Y", "B Y", "C Y", "C Z", "A Z", "B Z"}

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	totals = 0
	for fscanner.Scan() {
		line = fscanner.Text()

		// get index from list
		for i := 0; i < len(outcomes); i++ {
			if outcomes[i] == line {
				index = i + 1
				break
			}
		}
		totals += index
	}

	return totals
}

func main() {
	var testfile string
	var inputfile string

	testfile = "tests.txt"
	inputfile = "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day01")
	fmt.Printf("Tests : Answer to Part 1 = %d\n", Part1(testfile))
	fmt.Printf("Tests : Answer to Part 2 = %d\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %d\n", Part1(inputfile))
	fmt.Printf("Inputs: Answer to Part 2 = %d\n", Part2(inputfile))
}
