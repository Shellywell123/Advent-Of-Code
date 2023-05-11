package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Parse(filename string) [][]int {

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	data := [][]int{}

	for fscanner.Scan() {
		line := fscanner.Text()

		lineData := []int{}
		for i := 0; i < len(line); i++ {
			char, _ := strconv.Atoi(string(line[i]))
			lineData = append(lineData, char)
		}
		data = append(data, lineData)
	}

	return data
}

func Part1(filename string) int {

	data := Parse(filename)

	numOfVisible := (2 * len(data)) + (2 * len(data[0])) - 4

	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[y])-1; x++ {

			visibleUp := true
			visibleDown := true
			visibleLeft := true
			visibleRight := true

			for lc := 0; lc < x; lc++ {
				if data[y][lc] >= data[y][x] {
					visibleLeft = false
				}
			}

			for rc := x + 1; rc < len(data[x]); rc++ {
				if data[y][rc] >= data[y][x] {
					visibleRight = false
				}
			}

			for uc := 0; uc < y; uc++ {
				if data[uc][x] >= data[y][x] {
					visibleUp = false
				}
			}

			for dc := y + 1; dc < len(data); dc++ {
				if data[dc][x] >= data[y][x] {
					visibleDown = false
				}
			}

			if visibleLeft || visibleRight || visibleUp || visibleDown {
				numOfVisible += 1
			}
		}
	}
	return numOfVisible
}

func Part2(filename string) int {

	data := Parse(filename)

	scenicScore := 0

	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[y])-1; x++ {

			distanceLeft := 0
			distanceRight := 0
			distanceUp := 0
			distanceDown := 0

			for lc := x - 1; lc >= 0; lc-- {
				distanceLeft++
				if data[y][lc] >= data[y][x] {
					break
				}
			}

			for rc := x + 1; rc < len(data[x]); rc++ {
				distanceRight++
				if data[y][rc] >= data[y][x] {
					break
				}
			}

			for uc := y - 1; uc >= 0; uc-- {
				distanceUp++
				if data[uc][x] >= data[y][x] {
					break
				}
			}

			for dc := y + 1; dc < len(data); dc++ {
				distanceDown++
				if data[dc][x] >= data[y][x] {
					break
				}
			}

			currentScenicScore := distanceLeft * distanceRight * distanceUp * distanceDown
			if currentScenicScore > scenicScore {
				scenicScore = currentScenicScore
			}
		}
	}
	return scenicScore
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day06")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))
}
