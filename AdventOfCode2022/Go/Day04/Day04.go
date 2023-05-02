package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) int {

	total := 0

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	// loop through file reading one line at a time
	for fscanner.Scan() {
		line := fscanner.Text()

		// parse a line into pairs of elf section limits
		elves := strings.Split(line, ",")

		elf1Sections := strings.Split(elves[0], "-")
		elf2Sections := strings.Split(elves[1], "-")

		elf1Low, err := strconv.Atoi(elf1Sections[0])
		if err != nil {
			panic(err)
		}

		elf1High, err := strconv.Atoi(elf1Sections[1])
		if err != nil {
			panic(err)
		}

		elf2Low, err := strconv.Atoi(elf2Sections[0])
		if err != nil {
			panic(err)
		}

		elf2High, err := strconv.Atoi(elf2Sections[1])
		if err != nil {
			panic(err)
		}

		// check if limits contain one another
		if (elf1Low <= elf2Low && elf1High >= elf2High) || (elf2Low <= elf1Low && elf2High >= elf1High) {
			fmt.Println(line)
			total += 1
		}
	}
	return total
}

func Part2(filename string) int {

	total := 0

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	// loop through file reading one line at a time
	for fscanner.Scan() {
		line := fscanner.Text()

		// parse a line into pairs of elf section limits
		elves := strings.Split(line, ",")

		elf1Sections := strings.Split(elves[0], "-")
		elf2Sections := strings.Split(elves[1], "-")

		elf1Low, err := strconv.Atoi(elf1Sections[0])
		if err != nil {
			panic(err)
		}

		elf1High, err := strconv.Atoi(elf1Sections[1])
		if err != nil {
			panic(err)
		}

		elf2Low, err := strconv.Atoi(elf2Sections[0])
		if err != nil {
			panic(err)
		}

		elf2High, err := strconv.Atoi(elf2Sections[1])
		if err != nil {
			panic(err)
		}

		// check if limits contain one another
		if ((elf1Low >= elf2Low && elf1Low <= elf2High) || (elf2Low >= elf1Low && elf2Low <= elf1High) || 
			(elf1High >= elf2Low && elf1High <= elf2High) || (elf2High >= elf1Low && elf2High <= elf1High)) {
			fmt.Println(line)
			total += 1
		}
	}
	return total
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day04")
	fmt.Printf("Tests : Answer to Part 1 = %d\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %d\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %d\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %d\n", Part2(inputfile))
}
