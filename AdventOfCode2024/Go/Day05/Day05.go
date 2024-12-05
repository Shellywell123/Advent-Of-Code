package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) ([][]int, [][]int) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	rules := [][]int{}
	pagesSections := [][]int{}

	rulesInput := true
	for fscanner.Scan() {
		line := fscanner.Text()

		if line == "" {
			rulesInput = false
			continue
		}

		if rulesInput {
			// split by comma
			pair := []int{}
			for _, number := range strings.Split(line, "|") {
				n, _ := strconv.Atoi(string(number))
				pair = append(pair, n)
			}
			rules = append(rules, pair)
		} else {
			// split by comma
			pageSection := []int{}
			for _, number := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(string(number))
				pageSection = append(pageSection, n)
			}
			pagesSections = append(pagesSections, pageSection)
		}
	}

	return rules, pagesSections
}

func reorder(pageSection []int, rules [][]int, changes int) []int {

	for i, number := range pageSection {
		for _, rule := range rules {
			if number == rule[0] {
				// check if rule is incorrect
				for p, previousPageNumber := range pageSection[:i] {
					if previousPageNumber == rule[1] {
						// swap
						pageSection[i], pageSection[p] = pageSection[p], pageSection[i]
						changes++
					}
				}
			}
		}
	}
	if changes > 0 {
		reorder(pageSection, rules, 0)
	}
	return pageSection
}

func Part1(filename string, debug bool) int {
	rules, pagesSections := parse(filename)

	ans := 0

	for _, pageSection := range pagesSections {
		pageSectionCorrect := true

		for i, number := range pageSection {

			// check if page number is correct position
			for _, rule := range rules {

				// rule applies to this number
				if number == rule[0] {

					// check if rule is incorrect
					for _, previousPageNumber := range pageSection[:i] {
						if previousPageNumber == rule[1] {
							pageSectionCorrect = false
							break
						}
					}
				}
			}

			if !pageSectionCorrect {
				break
			}
		}

		if pageSectionCorrect {
			// get middle num
			if debug {
				fmt.Println("Page section correct", pageSection)
			}
			middleNum := pageSection[len(pageSection)/2]
			ans += middleNum
		}
	}

	return ans
}

func Part2(filename string, debug bool) int {
	rules, pagesSections := parse(filename)

	ans := 0

	for _, pageSection := range pagesSections {
		pageSectionCorrect := true

		for i, number := range pageSection {

			// check if page number is correct position
			for _, rule := range rules {

				// rule applies to this number
				if number == rule[0] {

					// check if rule is incorrect
					for _, previousPageNumber := range pageSection[:i] {
						if previousPageNumber == rule[1] {
							pageSectionCorrect = false
							break
						}
					}
				}
			}

			if !pageSectionCorrect {
				break
			}
		}

		if !pageSectionCorrect {

			if debug {
				fmt.Println("Page section incorrect", pageSection)
			}

			// reorder
			reorderedPageSection := reorder(pageSection, rules, 0)

			if debug {
				fmt.Println("Reordered page section: ", reorderedPageSection)
			}

			// get middle num
			middleNum := pageSection[len(pageSection)/2]
			ans += middleNum
		}
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day05")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
