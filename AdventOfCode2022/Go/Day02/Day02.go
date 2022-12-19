package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) uint64 {
	var totals []uint64
	var currentBag uint64
	var ans uint64
	var line string

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	currentBag = 0
	for fscanner.Scan() {
		line = fscanner.Text()

		if strings.TrimSpace(line) == "" {
			totals = append(totals, currentBag)
			currentBag = 0
			continue
		} else {
			i, _ := strconv.ParseUint(line, 10, 64)
			currentBag += i
		}
	}

	ans = 0
	for j := 0; j < len(totals); j++ {
		if ans < totals[j] {
			ans = totals[j]
		}
	}

	return ans
}

func Part2(filename string) uint64 {
	var totals []uint64
	var currentBag uint64
	var line string

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	currentBag = 0
	for fscanner.Scan() {
		line = fscanner.Text()

		if strings.TrimSpace(line) == "" {
			totals = append(totals, currentBag)
			currentBag = 0
			continue
		} else {
			i, _ := strconv.ParseUint(line, 10, 64)
			currentBag += i
		}
	}

	var ans1 uint64
	ans1 = 0
	for j := 0; j < len(totals); j++ {
		if ans1 < totals[j] {
			ans1 = totals[j]
		}
	}

	var ans2 uint64
	ans2 = 0
	for j := 0; j < len(totals); j++ {
		if ans2 < totals[j] && totals[j] < ans1 {
			ans2 = totals[j]
		}
	}

	var ans3 uint64
	ans3 = 0
	for j := 0; j < len(totals); j++ {
		if ans3 < totals[j] && totals[j] < ans2 {
			ans3 = totals[j]
		}
	}

	return ans1 + ans2 + ans3
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
