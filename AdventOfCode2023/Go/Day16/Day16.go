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

func MapBeams(start bool, startCoord Coord, grid [][]string, energized map[string]string) map[string]string {

	// check if weve been here before already
	if previousDirection, alreadyEnergized := energized[fmt.Sprintf("%d-%d", startCoord.x, startCoord.y)]; alreadyEnergized && previousDirection == startCoord.dir {
		return energized
	}
	if !start {
		energized[fmt.Sprintf("%d-%d", startCoord.x, startCoord.y)] = startCoord.dir
	}

	switch startCoord.dir {
	case "r":
		if !start && startCoord.x+1 == len(grid[0]) {
			break
		}

		switch grid[startCoord.y][startCoord.x+1] {
		case ".", "-":
			energized = MapBeams(false, Coord{startCoord.x + 1, startCoord.y, "r"}, grid, energized)
		case "|":
			energized = MapBeams(false, Coord{startCoord.x + 1, startCoord.y, "u"}, grid, energized)
			energized = MapBeams(false, Coord{startCoord.x + 1, startCoord.y, "d"}, grid, energized)
		case "\\":
			energized = MapBeams(false, Coord{startCoord.x + 1, startCoord.y, "d"}, grid, energized)
		case "/":
			energized = MapBeams(false, Coord{startCoord.x + 1, startCoord.y, "u"}, grid, energized)
		}
	case "l":
		if !start && startCoord.x-1 < 0 {
			break
		}
		switch grid[startCoord.y][startCoord.x-1] {
		case ".", "-":
			energized = MapBeams(false, Coord{startCoord.x - 1, startCoord.y, "l"}, grid, energized)
		case "|":
			energized = MapBeams(false, Coord{startCoord.x - 1, startCoord.y, "u"}, grid, energized)
			energized = MapBeams(false, Coord{startCoord.x - 1, startCoord.y, "d"}, grid, energized)
		case "\\":
			energized = MapBeams(false, Coord{startCoord.x - 1, startCoord.y, "u"}, grid, energized)
		case "/":
			energized = MapBeams(false, Coord{startCoord.x - 1, startCoord.y, "d"}, grid, energized)
		}
	case "d":
		if !start && startCoord.y+1 == len(grid) {
			break
		}
		switch grid[startCoord.y+1][startCoord.x] {
		case ".", "|":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y + 1, "d"}, grid, energized)
		case "-":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y + 1, "l"}, grid, energized)
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y + 1, "r"}, grid, energized)
		case "\\":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y + 1, "r"}, grid, energized)
		case "/":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y + 1, "l"}, grid, energized)
		}
	case "u":
		if !start && startCoord.y-1 < 0 {
			break
		}
		switch grid[startCoord.y-1][startCoord.x] {
		case ".", "|":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y - 1, "u"}, grid, energized)
		case "-":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y - 1, "l"}, grid, energized)
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y - 1, "r"}, grid, energized)
		case "\\":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y - 1, "l"}, grid, energized)
		case "/":
			energized = MapBeams(false, Coord{startCoord.x, startCoord.y - 1, "r"}, grid, energized)
		}
	}
	return energized
}

func Part1(filename string) int {
	data := Parse(filename)
	energized := MapBeams(true, Coord{-1, 0, "r"}, data, map[string]string{})
	return len(energized)
}

func Part2(filename string) int {
	data := Parse(filename)
	startCoords := []Coord{}
	for i := 0; i < len(data[0]); i++ {
		startCoords = append(startCoords, Coord{i, -1, "d"})
		startCoords = append(startCoords, Coord{i, len(data), "u"})
	}
	for i := 0; i < len(data); i++ {
		startCoords = append(startCoords, Coord{-1, i, "r"})
		startCoords = append(startCoords, Coord{len(data[0]), i, "l"})
	}

	mostEnergized := 0
	for _, startCoord := range startCoords {
		energized := len(MapBeams(true, startCoord, data, map[string]string{}))
		if energized > mostEnergized {
			mostEnergized = energized
		}
	}
	return mostEnergized
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day16")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
