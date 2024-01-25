package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	data := [][]int{}
	for fscanner.Scan() {
		line := fscanner.Text()
		datarow := []int{}
		for _, num := range strings.Split(line, "") {
			n, _ := strconv.Atoi(num)
			datarow = append(datarow, n)
		}
		data = append(data, datarow)
	}
	return data
}

type Coord struct {
	x   int
	y   int
	dir string
}

func MinLossRoute(currentCoord Coord, endCoord Coord, grid [][]int, lossgrid [][]int, pathLoss int, straight int) {

	fmt.Print(currentCoord, "->")

	// if out of bounds
	if (currentCoord.dir == "r" && currentCoord.x+1 == len(grid[0])) ||
		(currentCoord.dir == "l" && currentCoord.x-1 < 0) ||
		(currentCoord.dir == "d" && currentCoord.y+1 == len(grid)) ||
		(currentCoord.dir == "u" && currentCoord.y-1 < 0) {
		fmt.Print("out\n")
		return
	}

	pathLoss += grid[currentCoord.y][currentCoord.x]

	// if we previously had less heat loss
	if pathLoss > lossgrid[currentCoord.y][currentCoord.x] {
		fmt.Printf("expensive\n")
		return
	}

	lossgrid[currentCoord.y][currentCoord.x] = pathLoss

	// if we reach the endCoord
	if currentCoord.x == endCoord.x && currentCoord.y == endCoord.y {
		fmt.Println("we made it", pathLoss)
		return
	}

	switch currentCoord.dir { // can optimise by making dirs int
	case "r":
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x + 1, currentCoord.y, "r"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x + 1, currentCoord.y, "u"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x + 1, currentCoord.y, "d"}, endCoord, grid, lossgrid, pathLoss, 0)

	case "l":
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x - 1, currentCoord.y, "l"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x - 1, currentCoord.y, "d"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x - 1, currentCoord.y, "u"}, endCoord, grid, lossgrid, pathLoss, 0)

	case "d":
		// forward
		if straight != 2 {
			MinLossRoute(Coord{currentCoord.x, currentCoord.y + 1, "d"}, endCoord, grid, lossgrid, pathLoss, straight+1)
		}
		// left
		MinLossRoute(Coord{currentCoord.x, currentCoord.y + 1, "r"}, endCoord, grid, lossgrid, pathLoss, 0)
		// right
		MinLossRoute(Coord{currentCoord.x, currentCoord.y + 1, "l"}, endCoord, grid, lossgrid, pathLoss, 0)

	case "u":
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
			lossgridrow = append(lossgridrow, 100000000)
		}
		lossgrid = append(lossgrid, lossgridrow)
	}

	endCoord := Coord{len(data[0]) - 1, len(data) - 1, "r"}

	// recursive populate lossgrid with min loss to each pos
	MinLossRoute(Coord{0, 0, "r"}, endCoord, data, lossgrid, 0, 0)
	MinLossRoute(Coord{0, 0, "d"}, endCoord, data, lossgrid, 0, 0)

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
