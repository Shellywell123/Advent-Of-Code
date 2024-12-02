package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func Parse(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	reports := [][]int{}

	for fscanner.Scan() {
		line := fscanner.Text()

		line_numbers := []int{}
		for _, number := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(string(number))
			line_numbers = append(line_numbers, n)
		}
		reports = append(reports, line_numbers)
	}

	return reports
}

func nonZeroGradient(report []int) bool {

	for n := 0; n < len(report)-1; n++ {

		// beyond gradient limit
		if math.Abs(float64(report[n+1] - report[n])) > 3 {
			return false
		}
			
		// no gradient
		if report[n+1] == report[n] {
			return false
		}
	}
	return true
}

func constGradient(report []int) bool {

	for n := 0; n < len(report)-2; n++ {
			
		if report[n] < report[n+1] && report[n+1] > report[n+2]{
			return false
		}

		if report[n] > report[n+1] && report[n+1] < report[n+2]{
			return false
		}
	}
	return true
}

func Part1(filename string) int {
	reports := Parse(filename)

	ans := 0
	for i := 0; i < len(reports); i++ {
		report := reports[i]

		if constGradient(report) && nonZeroGradient(report){
			ans ++
		}
	}
	return ans
}

func Part2(filename string) int {
	reports := Parse(filename)

	ans := 0
	for i := 0; i < len(reports); i++ {
		report := reports[i]

		// same as part 1
		if constGradient(report) && nonZeroGradient(report){
			ans ++
			continue
		}

		// brute force this bad boy
		for r :=0; r < len(report); r++ {

			// make a disgusting copy
			report1 := make([]int, len(report))
			copy(report1, report)

			// trial and error each combo
			temp_report := append(report1[:r], report1[r+1:]...)
		
			if constGradient(temp_report) && nonZeroGradient(temp_report){
				fmt.Println(report, temp_report)
				ans ++
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day02")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
