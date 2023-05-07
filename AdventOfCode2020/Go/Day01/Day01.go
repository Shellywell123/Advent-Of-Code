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

	for _, a := range data {
		for _, b := range data {

			aInt, err := strconv.Atoi(a)
			if err != nil {
				panic(err)
			}

			bInt, err := strconv.Atoi(b)
			if err != nil {
				panic(err)
			}

			if aInt+bInt == 2020 {
				return aInt * bInt
			}
		}
	}
	return 0
}

func Part2(filename string) int {

	data := Parse(filename)

	for _, a := range data {
		for _, b := range data {
			for _, c := range data {

				aInt, err := strconv.Atoi(a)
				if err != nil {
					panic(err)
				}

				bInt, err := strconv.Atoi(b)
				if err != nil {
					panic(err)
				}

				cInt, err := strconv.Atoi(c)
				if err != nil {
					panic(err)
				}

				if aInt+bInt+cInt == 2020 {
					return aInt * bInt * cInt
				}
			}
		}
	}
	return 0
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
