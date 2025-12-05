package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) ([][]int, []int) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	ranges := true
	rangesList := [][]int{}
	ingredientList := []int{}

	for fscanner.Scan() {
		line := fscanner.Text()
		if line == "" {
			ranges = false
			continue
		}

		if ranges {
			rangesStrings := strings.Split(line, "-")
			lowerLimit, _ := strconv.Atoi(rangesStrings[0])
			upperLimit, _ := strconv.Atoi(rangesStrings[1])
			rangesList = append(rangesList, []int{lowerLimit, upperLimit})
			continue
		}

		IngreidientId, _ := strconv.Atoi(line)
		ingredientList = append(ingredientList, IngreidientId)

	}

	return rangesList, ingredientList
}

func Part1(filename string, debug bool) int {
	freshRanges, ingredientIDs := Parse(filename)

	ans := 0
	for _, ingredientID := range ingredientIDs {
		for _, rangePair := range freshRanges {
			if ingredientID >= rangePair[0] && ingredientID <= rangePair[1] {
				ans++
				break
			}
		}
	}
	return ans
}

func Part2(filename string, debug bool) int {
	freshRanges, _ := Parse(filename)

	for {
		changes := 0
		breakout := false

		for i, firstranges := range freshRanges {
			if breakout {
				break
			}
			for j, secondranges := range freshRanges {
				if breakout {
					break
				}

				if i == j {
					continue
				}

				if debug {
					fmt.Println(firstranges, secondranges)
				}

				// remove all completely overlapping ranges
				if firstranges[0] >= secondranges[0] && firstranges[1] <= secondranges[1] {
					freshRanges = append(freshRanges[:i], freshRanges[i+1:]...)
					changes++
					break
				}

				// replace partial overlaps
				if firstranges[0] <= secondranges[1] && firstranges[1] >= secondranges[0] {
					if debug {
						fmt.Println("merging", firstranges, secondranges)
					}
					newRange := []int{min(firstranges[0], secondranges[0]), max(firstranges[1], secondranges[1])}
					if debug {
						fmt.Println("new range", newRange)
					}

					freshRanges[j] = newRange
					freshRanges = append(freshRanges[:i], freshRanges[i+1:]...)
					changes++
					breakout = true
					break
				}
			}
		}

		if changes == 0 {
			break
		}
	}

	ans := 0
	for _, rangePair := range freshRanges {

		ans += rangePair[1] - rangePair[0] + 1
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day05")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
