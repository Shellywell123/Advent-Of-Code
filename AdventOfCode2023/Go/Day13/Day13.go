package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(filename string) [][]string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	patterns := [][]string{}

	newPattern := []string{}
	for fscanner.Scan() {

		line := fscanner.Text()

		if line == "" {
			patterns = append(patterns, newPattern)
			newPattern = []string{}
			continue
		}
		newPattern = append(newPattern, line)

	}
	patterns = append(patterns, newPattern)

	return patterns
}

func reverseSlice(slice []string) []string {
	reversedSlice := []string{}
	for i := len(slice) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, slice[i])
	}
	return reversedSlice
}

func searchForHorizontalReflection(pattern []string, tolerance int) int {
	reflectionColumn := 0

	for col := 1; col < len(pattern[0]); col++ {
		t := 0
		for row := 0; row < len(pattern); row++ {
			rowSlice := strings.Split(pattern[row], "")
			left := fmt.Sprint(reverseSlice(rowSlice[:col]))[1:]
			right := fmt.Sprint(rowSlice[col:])[1:]
			width := min(len(left), len(right)) - 1

			for c := 0; c < width; c++ {
				if left[c] != right[c] {
					t++
				}
				if t > tolerance { // optimization
					break
				}
			}
			if t > tolerance { // optimization
				break
			}

			if row == len(pattern)-1 && t == tolerance { // if we reach the last row that means we have a mirror on this column
				return col
			}
		}
	}
	return reflectionColumn
}

func searchForVerticalReflection(pattern []string, tolerance int) int {
	reflectionRow := 0
	for row := 1; row < len(pattern); row++ {
		t := 0
		for col := 0; col < len(pattern[0]); col++ {
			colSlice := []string{}
			for _, line := range pattern {
				colSlice = append(colSlice, string(line[col]))
			}
			if len(colSlice) != len(reverseSlice(colSlice)) {
				os.Exit(1)
			}
			top := fmt.Sprint(reverseSlice(colSlice[:row]))[1:]
			bottom := fmt.Sprint(colSlice[row:])[1:]
			depth := min(len(top), len(bottom)) - 1

			for c := 0; c < depth; c++ {
				if top[c] != bottom[c] {
					t++
					if t > tolerance { // optimization
						break
					}
				}
			}
			if t > tolerance { // optimization
				break
			}

			if col == len(pattern[0])-1 && t == tolerance { // if we reach the last row that means we have a mirror on this row
				return row
			}
		}
	}

	return reflectionRow
}

func Solve(filename string, tolerance int) int {
	patterns := (Parse(filename))

	reflectionColumns := 0
	reflectionRows := 0

	for _, pattern := range patterns {
		reflectionRows += searchForVerticalReflection(pattern, tolerance)
		reflectionColumns += searchForHorizontalReflection(pattern, tolerance)
	}

	return reflectionColumns + (reflectionRows * 100)
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day13")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Solve("tests.txt", 0))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Solve("inputs.txt", 0))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Solve("tests.txt", 1))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Solve("inputs.txt", 1))
}
