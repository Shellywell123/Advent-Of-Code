package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
}

func parse(filename string) ([][]int, []coord) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	y := 0
	startCoords := []coord{}

	grid := [][]int{}
	for fscanner.Scan() {
		line := fscanner.Text()

		gridLine := []int{}
		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				gridLine = append(gridLine, -1)
				continue
			}
			n, _ := strconv.Atoi(string(line[x]))

			if n == 0 {
				startCoords = append(startCoords, coord{x, y})
			}
			gridLine = append(gridLine, n)
		}
		y++
		grid = append(grid, gridLine)
	}

	return grid, startCoords
}

func trailBlazer(grid [][]int, pathGrid [][]int, startCoord coord, score int, part int) (int, [][]int) {

	if grid[startCoord.y][startCoord.x] == 9 && ((part == 1 && pathGrid[startCoord.y][startCoord.x] == 0 || part == 2 )){
		pathGrid[startCoord.y][startCoord.x] = 1
		return score + 1, pathGrid
	}

	x := startCoord.x
	y := startCoord.y

	// for each direction
	upCoord := coord{x, y - 1}
	downCoord := coord{x, y + 1}
	leftCoord := coord{x - 1, y}
	rightCoord := coord{x + 1, y}

	for _, nextCoord := range []coord{upCoord, downCoord, leftCoord, rightCoord} {

		// check if next step in bound
		if nextCoord.x < 0 || nextCoord.y < 0 || nextCoord.x >= len(grid[0]) || nextCoord.y >= len(grid) {
			continue
		}
		// check if next step is a wall
		if grid[nextCoord.y][nextCoord.x] == -1 {
			continue
		}

		// check if next step is a trail
		if grid[nextCoord.y][nextCoord.x] == grid[y][x]+1 {
			// recurse
			score, pathGrid = trailBlazer(grid, pathGrid, nextCoord, score, part)
		}

	}

	return score, pathGrid
}

func Solve(filename string, debug bool, part int) int {
	grid, startCoords := parse(filename)

	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				fmt.Printf("%v", grid[y][x])
			}
			fmt.Println()
		}

		fmt.Printf("Start: %v\n", startCoords)
	}

	ans := 0
	// for each trail head
	for _, startCoord := range startCoords {
		pathGrid := make([][]int, len(grid))
		for i := range pathGrid {
			pathGrid[i] = make([]int, len(grid[0]))
		}
		ans, pathGrid = trailBlazer(grid, pathGrid, startCoord, ans, part)
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day10")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Solve("tests.txt", false, 1))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Solve("inputs.txt", false, 1))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Solve("tests.txt", false, 2))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Solve("inputs.txt", false, 2))
}
