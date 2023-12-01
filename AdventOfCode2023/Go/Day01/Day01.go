package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Packet struct {
	left  string
	right string
}

func isInt(char byte) bool {
	if _, err := strconv.Atoi(string(char)); err == nil {
		return true
	}
	return false
}

func wordToNumber(slice string) int {

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, number := range numbers {

		if strings.HasPrefix(slice, number) {
			return slices.Index(numbers, number) + 1
		}
	}

	return -1

}

func Part1(filename string) int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	calibrations := []int{}

	for fscanner.Scan() {
		line := fscanner.Text()

		leftCalibration := ""
		rightCalibration := ""

		for i, _ := range line {

			char := line[i]

			if isInt(char) && leftCalibration == "" {
				leftCalibration = string(char)
			}

			if isInt(char) {
				rightCalibration = string(char)
			}
		}
		calibration, _ := strconv.Atoi(leftCalibration + rightCalibration)
		calibrations = append(calibrations, calibration)
	}

	total := 0
	for _, calibration := range calibrations {
		total += calibration
	}
	return total
}

func Part2(filename string) int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	calibrations := []int{}
	for fscanner.Scan() {

		line := fscanner.Text()

		leftCalibration := -1
		rightCalibration := -1

		for i, _ := range line {
			char := line[i]

			wordValue := -1
			fix := (len(line)) - (i + 5)
			if fix > 0 {
				fix = 0
			}
			wordValue = wordToNumber(string(line[i : i+5+fix]))

			if leftCalibration == -1 {
				if isInt(char) {
					leftCalibration, _ = strconv.Atoi(string(char))
				} else if wordValue != -1 {
					leftCalibration = wordValue
				}
			}

			if isInt(char) {
				rightCalibration, _ = strconv.Atoi(string(char))
			} else if wordValue != -1 {
				rightCalibration = wordValue
			}
		}

		calibration, _ := strconv.Atoi(fmt.Sprintf("%d%d", leftCalibration, rightCalibration))
		calibrations = append(calibrations, calibration)
	}

	total := 0
	for _, calibration := range calibrations {
		total += calibration
	}
	return total
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day01")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests2.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
