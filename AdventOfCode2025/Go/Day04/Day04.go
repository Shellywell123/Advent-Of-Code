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

	grid := [][]string{}

	for fscanner.Scan() {
		line := fscanner.Text()

		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func checkNiebours(x, y int, grid [][]string) int {

	atCounter := 0

	// check top left
	if y > 0 && x > 0 {
		topLeft := grid[y-1][x-1]
		if topLeft == "@" || topLeft == "x" {
			atCounter++
		}
	}

	// check top
	if y > 0 {
		top := grid[y-1][x]
		if top == "@" || top == "x" {
			atCounter++
		}
	}

	// check top right
	if y > 0 && x < len(grid[0])-1 {
		topRight := grid[y-1][x+1]
		if topRight == "@" || topRight == "x" {
			atCounter++
		}
	}

	// check left
	if x > 0 {
		left := grid[y][x-1]
		if left == "@" || left == "x" {
			atCounter++
		}
	}

	// check right
	if x < len(grid[0])-1 {
		right := grid[y][x+1]
		if right == "@" || right == "x" {
			atCounter++
		}
	}

	// check bottom left
	if y < len(grid)-1 && x > 0 {
		bottomLeft := grid[y+1][x-1]
		if bottomLeft == "@" || bottomLeft == "x" {
			atCounter++
		}
	}

	// check bottom
	if y < len(grid)-1 {
		bottom := grid[y+1][x]
		if bottom == "@" || bottom == "x" {
			atCounter++
		}
	}

	// check bottom	right
	if y < len(grid)-1 && x < len(grid[0])-1 {
		bottomRight := grid[y+1][x+1]
		if bottomRight == "@" || bottomRight == "x" {
			atCounter++
		}
	}

	return atCounter
}

func Part1(filename string, debug bool) int {
	grid := Parse(filename)

	ans := 0
	for y, row := range grid {
		for x, _ := range row {
			if grid[y][x] == "." {
				continue
			}
			if checkNiebours(x, y, grid) < 4 {
				ans += 1
				grid[y][x] = "x"
			}
		}
		if debug {
			fmt.Println(y, row)
		}
	}

	return ans
}

func Part2(filename string, debug bool) int {
	grid := Parse(filename)

	ans := 0
	removed := 0
	for {
		removed = 0
		for y, row := range grid {
			for x, _ := range row {
				if grid[y][x] == "." {
					continue
				}
				if checkNiebours(x, y, grid) < 4 {
					removed += 1
					grid[y][x] = "."
				}
			}
			if debug {
				fmt.Println(y, row)
			}
		}
		ans += removed
		if removed == 0 {
			break
		}
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day04")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
