package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) ([]int, []int) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	times := []int{}
	distances := []int{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		if y == 0 {
			for _, num := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line[9:], "  ", " "), "  ", " "), "  ", " "), " ") {
				time, _ := strconv.Atoi(string(num))
				if time == 0 {
					continue
				}
				times = append(times, time)
			}
		}
		if y == 1 {
			for _, num := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line[9:], "  ", " "), "  ", " "), "  ", " "), " ") {
				distance, _ := strconv.Atoi(string(num))
				if distance == 0 {
					continue
				}
				distances = append(distances, distance)
			}
		}
		y++
	}

	return times, distances
}

func Parse2(filename string) ([]int, []int) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	times := []int{}
	distances := []int{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		if y == 0 {
			for _, num := range strings.Split(strings.ReplaceAll(line[9:], " ", ""), " ") {
				time, _ := strconv.Atoi(string(num))
				if time == 0 {
					continue
				}
				times = append(times, time)
			}
		}
		if y == 1 {
			for _, num := range strings.Split(strings.ReplaceAll(line[9:], " ", ""), " ") {
				distance, _ := strconv.Atoi(string(num))
				if distance == 0 {
					continue
				}
				distances = append(distances, distance)
			}
		}
		y++
	}

	return times, distances
}

func Part1(filename string) int {
	times, distances := Parse(filename)

	ans := 1
	for i := 0; i <= len(times)-1; i++ {
		beat := 0
		for x := 1; x <= times[i]-1; x++ {
			y := (x * times[i]) - (x * x)
			if y > distances[i] {
				beat++
			}
		}
		ans *= beat
	}

	return ans
}

func Part2(filename string) int {
	times, distances := Parse2(filename)

	ans := 1
	for i := 0; i <= len(times)-1; i++ {
		beat := 0
		for x := 1; x <= times[i]-1; x++ {
			y := (x * times[i]) - (x * x)
			if y > distances[i] {
				beat++
			}
		}
		ans *= beat
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day06")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
