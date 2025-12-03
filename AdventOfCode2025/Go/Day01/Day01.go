package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Parse(filename string) []string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	list1 := []string{}

	for fscanner.Scan() {
		line := fscanner.Text()
		list1 = append(list1, line)
	}

	return list1
}

func Part1(filename string, debug bool) int {
	list1 := Parse(filename)

	dial := 50
	ans := 0

	for _, v := range list1 {
		if debug {
			fmt.Printf("Instruction: %v\n", v)
		}

		rotation, _ := strconv.Atoi(v[1:])

		if v[0] == 'L' {
			dial -= rotation
		}

		if v[0] == 'R' {
			dial += rotation

		}

		for dial < 0 {
			dial += 100
		}
		for dial > 99 {
			dial -= 100
		}
		if dial == 0 {
			ans++
		}

		if debug {
			fmt.Printf("Dial now at: %v\n", dial)
		}
	}

	return ans
}

func Part2(filename string, debug bool) int {
	list1 := Parse(filename)

	dial := 50
	ans := 0

	for _, v := range list1 {
		if debug {
			fmt.Printf("Instruction: %v\n", v)
		}

		rotation, _ := strconv.Atoi(v[1:])

		if v[0] == 'L' {
			dial -= rotation
		}

		if v[0] == 'R' {
			dial += rotation
		}

		minusOne := false

		for dial < 0 {
			if dial != -rotation {
				ans++
				if debug {
					fmt.Println(" Through Zero!")
				}
			}

			dial += 100
		}

		for dial > 99 {
			ans++
			if debug {
				fmt.Println("Through Zero!")
			}
			dial -= 100
			minusOne = true
		}

		if dial == 0 {
			ans++
		}

		if minusOne && dial == 0 {
			ans--
		}

		if debug {
			fmt.Printf(" Dial now at: %v\n", dial)
		}
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day01")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
