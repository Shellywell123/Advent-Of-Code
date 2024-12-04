package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x int
	y int
}

func parse(filename string) [][]rune {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	y := 0

	grid := [][]rune{}
	for fscanner.Scan() {
		line := fscanner.Text()

		gridLine := []rune{}
		for x := 0; x < len(line); x++ {
			gridLine = append(gridLine, rune(line[x]))
		}
		y++
		grid = append(grid, gridLine)
	}

	return grid
}

func countHorizontal(grid [][]rune, ansgrid [][]rune) (int, [][]rune) {
	count := 0
	// loop down rows
	for y := 0; y < len(grid); y++ {
		// loop through columns
		for x := 3; x < len(grid[y]); x++ {

			// check for XMAS
			if grid[y][x-3] == 'X' && grid[y][x-2] == 'M' && grid[y][x-1] == 'A' && grid[y][x] == 'S' {
				count++
				ansgrid[y][x-3] = 'X'
				ansgrid[y][x-2] = 'M'
				ansgrid[y][x-1] = 'A'
				ansgrid[y][x] = 'S'
			}

			// check for SMAX
			if grid[y][x-3] == 'S' && grid[y][x-2] == 'A' && grid[y][x-1] == 'M' && grid[y][x] == 'X' {
				count++
				ansgrid[y][x-3] = 'S'
				ansgrid[y][x-2] = 'A'
				ansgrid[y][x-1] = 'M'
				ansgrid[y][x] = 'X'
			}
		}
	}
	return count, ansgrid
}

func countVertical(grid [][]rune, ansgrid [][]rune) (int, [][]rune) {
	count := 0
	// loop down rows
	for x := 0; x < len(grid[0]); x++ {
		// loop through columns
		for y := 3; y < len(grid); y++ {

			// check for XMAS
			if grid[y-3][x] == 'X' && grid[y-2][x] == 'M' && grid[y-1][x] == 'A' && grid[y][x] == 'S' {
				count++
				ansgrid[y-3][x] = 'X'
				ansgrid[y-2][x] = 'M'
				ansgrid[y-1][x] = 'A'
				ansgrid[y][x] = 'S'
			}

			// check for SMAX
			if grid[y-3][x] == 'S' && grid[y-2][x] == 'A' && grid[y-1][x] == 'M' && grid[y][x] == 'X' {
				count++
				ansgrid[y-3][x] = 'S'
				ansgrid[y-2][x] = 'A'
				ansgrid[y-1][x] = 'M'
				ansgrid[y][x] = 'X'
			}
		}
	}
	return count, ansgrid
}

func countDiagonal(grid [][]rune, ansgrid [][]rune) (int, [][]rune) {
	count := 0

	// \ diagonal
	// loop down rows
	for x := 3; x < len(grid[0]); x++ {
		// loop through columns
		for y := 3; y < len(grid); y++ {

			// check for XMAS
			if grid[y-3][x-3] == 'X' && grid[y-2][x-2] == 'M' && grid[y-1][x-1] == 'A' && grid[y][x] == 'S' {
				count++
				ansgrid[y-3][x-3] = 'X'
				ansgrid[y-2][x-2] = 'M'
				ansgrid[y-1][x-1] = 'A'
				ansgrid[y][x] = 'S'
			}

			// check for SMAX
			if grid[y-3][x-3] == 'S' && grid[y-2][x-2] == 'A' && grid[y-1][x-1] == 'M' && grid[y][x] == 'X' {
				count++
				ansgrid[y-3][x-3] = 'S'
				ansgrid[y-2][x-2] = 'A'
				ansgrid[y-1][x-1] = 'M'
				ansgrid[y][x] = 'X'
			}
		}
	}

	// / diagonal
	// loop down rows
	for x := 3; x < len(grid[0]); x++ {
		// loop through columns
		for y := 3; y < len(grid); y++ {

			// check for XMAS
			if grid[y-3][x] == 'X' && grid[y-2][x-1] == 'M' && grid[y-1][x-2] == 'A' && grid[y][x-3] == 'S' {
				count++
				ansgrid[y-3][x] = 'X'
				ansgrid[y-2][x-1] = 'M'
				ansgrid[y-1][x-2] = 'A'
				ansgrid[y][x-3] = 'S'
			}

			// check for SMAX
			if grid[y-3][x] == 'S' && grid[y-2][x-1] == 'A' && grid[y-1][x-2] == 'M' && grid[y][x-3] == 'X' {
				count++
				ansgrid[y-3][x] = 'S'
				ansgrid[y-2][x-1] = 'A'
				ansgrid[y-1][x-2] = 'M'
				ansgrid[y][x-3] = 'X'
			}
		}
	}

	return count, ansgrid
}

func countXmas(grid [][]rune, ansgrid [][]rune) (int, [][]rune) {
	count := 0

	// \ diagonal
	// loop down rows
	for x := 2; x < len(grid[0]); x++ {
		// loop through columns
		for y := 2; y < len(grid); y++ {

			// check for
			// M . S
			// . A .
			// M . S

			if grid[y-2][x-2] == 'M' && grid[y-2][x] == 'S' && grid[y-1][x-1] == 'A' && grid[y][x] == 'S' && grid[y][x-2] == 'M' {
				count++
				ansgrid[y-2][x-2] = 'M'
				ansgrid[y-2][x] = 'S'
				ansgrid[y-1][x-1] = 'A'
				ansgrid[y][x] = 'S'
				ansgrid[y][x-2] = 'M'
			}

			// check for
			// M . M
			// . A .
			// S . S
			if grid[y-2][x-2] == 'M' && grid[y-2][x] == 'M' && grid[y-1][x-1] == 'A' && grid[y][x] == 'S' && grid[y][x-2] == 'S' {
				count++
				ansgrid[y-2][x-2] = 'M'
				ansgrid[y-2][x] = 'M'
				ansgrid[y-1][x-1] = 'A'
				ansgrid[y][x] = 'S'
				ansgrid[y][x-2] = 'S'
			}

			// check for
			// S . M
			// . A .
			// S . M
			if grid[y-2][x-2] == 'S' && grid[y-2][x] == 'M' && grid[y-1][x-1] == 'A' && grid[y][x] == 'M' && grid[y][x-2] == 'S' {
				count++
				ansgrid[y-2][x-2] = 'S'
				ansgrid[y-2][x] = 'M'
				ansgrid[y-1][x-1] = 'A'
				ansgrid[y][x] = 'M'
				ansgrid[y][x-2] = 'S'
			}

			// check for
			// S . S
			// . A .
			// M . M
			if grid[y-2][x-2] == 'S' && grid[y-2][x] == 'S' && grid[y-1][x-1] == 'A' && grid[y][x] == 'M' && grid[y][x-2] == 'M' {
				count++
				ansgrid[y-2][x-2] = 'S'
				ansgrid[y-2][x] = 'S'
				ansgrid[y-1][x-1] = 'A'
				ansgrid[y][x] = 'M'
				ansgrid[y][x-2] = 'M'
			}
		}
	}

	return count, ansgrid
}

func Part1(filename string, debug bool) int {
	grid := parse(filename)

	ansgrid := [][]rune{}

	for y := 0; y < len(grid); y++ {
		gridLine := []rune{}
		for x := 0; x < len(grid[y]); x++ {
			gridLine = append(gridLine, '.')
		}
		ansgrid = append(ansgrid, gridLine)
	}

	ans1, ansgrid := countHorizontal(grid, ansgrid)
	ans2, ansgrid := countVertical(grid, ansgrid)
	ans3, ansgrid := countDiagonal(grid, ansgrid)

	if debug {
		fmt.Println("Answer Grid:")
		for y := 0; y < len(ansgrid); y++ {
			for x := 0; x < len(ansgrid[y]); x++ {
				fmt.Printf("%c", ansgrid[y][x])
			}
			fmt.Println()
		}
	}

	return ans1 + ans2 + ans3
}

func Part2(filename string, debug bool) int {
	grid := parse(filename)

	ansgrid := [][]rune{}

	for y := 0; y < len(grid); y++ {
		gridLine := []rune{}
		for x := 0; x < len(grid[y]); x++ {
			gridLine = append(gridLine, '.')
		}
		ansgrid = append(ansgrid, gridLine)
	}

	ans, ansgrid := countXmas(grid, ansgrid)

	if debug {
		fmt.Println("Answer Grid:")
		for y := 0; y < len(ansgrid); y++ {
			for x := 0; x < len(ansgrid[y]); x++ {
				fmt.Printf("%c", ansgrid[y][x])
			}
			fmt.Println()
		}
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day04")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
