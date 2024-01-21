package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) [][]string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	data := [][]string{}
	for fscanner.Scan() {
		line := fscanner.Text()
		data = append(data, strings.Split(line, ""))
	}
	return data
}

type Coord struct {
	x   int
	y   int
	dir string
}

func MinLossRoute(currentCoord Coord, endCoord Coord, grid [][]string, lossgrid [][]int, pathLoss int, straight int) {

	// update loss to current pos
	loss, _ := strconv.Atoi(grid[currentCoord.y][currentCoord.x])
	pathLoss += loss

	// if first time here
	if lossgrid[currentCoord.y][currentCoord.x] == 0 {
		lossgrid[currentCoord.y][currentCoord.x] = pathLoss
	} else if lossgrid[currentCoord.y][currentCoord.x] != 0 && pathLoss < lossgrid[currentCoord.y][currentCoord.x] { // if cheaper path found or equal (could possibly split out to optimize later)
		lossgrid[currentCoord.y][currentCoord.x] = pathLoss
	} else { // if more loss than previously  why bother continuing
		return
	}

	// if we have finsied
	if currentCoord.x == endCoord.x && currentCoord.y == endCoord.y {
		fmt.Println("we made it")
		return
	}

	switch currentCoord.dir {
	case "r":
		// out of bounds check
		if currentCoord.x+1 == len(grid[0]) {
			break
		}
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x + 1, currentCoord.y, "r"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x + 1, currentCoord.y, "u"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x + 1, currentCoord.y, "d"}, endCoord, grid, lossgrid, pathLoss, 0)

	case "l":
		// out of bounds check
		if currentCoord.x-1 < 0 {
			break
		}
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x - 1, currentCoord.y, "l"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x - 1, currentCoord.y, "d"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x - 1, currentCoord.y, "u"}, endCoord, grid, lossgrid, pathLoss, 0)

	case "d":
		// out of bounds check
		if currentCoord.y+1 == len(grid) {
			break
		}
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x, currentCoord.y + 1, "d"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x, currentCoord.y + 1, "r"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x, currentCoord.y + 1, "l"}, endCoord, grid, lossgrid, pathLoss, 0)

	case "u":
		// out of bounds check
		if currentCoord.y-1 < 0 {
			break
		}
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x, currentCoord.y - 1, "u"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x, currentCoord.y - 1, "l"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x, currentCoord.y - 1, "r"}, endCoord, grid, lossgrid, pathLoss, 0)

	}
}

func Part1(filename string) int {
	data := Parse(filename)

	// initialize a loss grid with 0 loss for each pos
	lossgrid := [][]int{}
	for y := 0; y < len(data); y++ {
		lossgridrow := []int{}
		for x := 0; x < len(data[0]); x++ {
			lossgridrow = append(lossgridrow, 0)
		}
		lossgrid = append(lossgrid, lossgridrow)
	}

	// start params
	startCoord := Coord{0, 0, "r"}
	endCoord := Coord{len(data[0]) - 1, len(data) - 1, "r"}
	pathLoss := 0
	straightcounter := 0

	// recursive populate lossgrid with min loss to each pos
	MinLossRoute(startCoord, endCoord, data, lossgrid, pathLoss, straightcounter)

	// print end loss grid
	for _, line := range lossgrid {
		fmt.Println(line)
	}

	return lossgrid[endCoord.y][endCoord.x]
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day17")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	// fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	// fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	// fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
