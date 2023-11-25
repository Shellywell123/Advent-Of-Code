package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"slices"
)

type Coord struct {
	x        int
	y        int
	altitude int
	distToE  int
}

func Parse(filename string) ([][]Coord, Coord, Coord) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	startCoord := Coord{}
	endCoord := Coord{}
	dataMatrix := [][]Coord{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		dataMatrixLine := []Coord{}
		for x, char := range line {
			if char == 'S' {
				startCoord = Coord{x: x, y: y, altitude: 0, distToE: 999}
				dataMatrixLine = append(dataMatrixLine, startCoord)
				continue
			}
			if char == 'E' {
				endCoord = Coord{x: x, y: y, altitude: 25, distToE: 0}
				dataMatrixLine = append(dataMatrixLine, endCoord)
				continue
			}
			altitude := strings.Index("abcdefghijklmnopqrstuvwxyz", string(char))
			coord := Coord{x: x, y: y, altitude: altitude, distToE: 999}
			dataMatrixLine = append(dataMatrixLine, coord)
		}
		dataMatrix = append(dataMatrix, dataMatrixLine)
		y++
	}
	return dataMatrix, startCoord, endCoord
}

func Solve(dataMatrix [][]Coord, startCoord Coord, endCoord Coord) int {
	finalDistance := 0
	for dataMatrix[startCoord.y][startCoord.x].distToE == 999 && finalDistance < 999 {

		for y := 0; y < len(dataMatrix); y++ {

			for x := 0; x < len(dataMatrix[0]); x++ {

				if dataMatrix[y][x].distToE == finalDistance {

					leftCoord := Coord{x: x - 1, y: y}
					rightCoord := Coord{x: x + 1, y: y}
					upCoord := Coord{x: x, y: y - 1}
					downCoord := Coord{x: x, y: y + 1}

					for _, coordPos := range []Coord{leftCoord, rightCoord, upCoord, downCoord} {

						if coordPos.x < 0 || coordPos.x >= len(dataMatrix[0]) || coordPos.y < 0 || coordPos.y >= len(dataMatrix) {
							continue
						}

						coord := dataMatrix[coordPos.y][coordPos.x]

						if coord.distToE != 999 && coord.distToE < dataMatrix[y][x].distToE+1 {
							continue
						}

						if coord.altitude + 2 > dataMatrix[y][x].altitude {
							dataMatrix[coord.y][coord.x].distToE = dataMatrix[y][x].distToE + 1
						}
					}
				}
			}
		}
		finalDistance++
	}
	return finalDistance
}

func Part1(filename string) int {
	dataMatrix, startCoord, endCoord := Parse(filename)
	return Solve(dataMatrix, startCoord, endCoord)
}

func Part2(filename string) int {
	pathLengths := []int{}
	dataMatrix, startCoord, endCoord := Parse(filename)
	Solve(dataMatrix, startCoord, endCoord)
	for _, line := range dataMatrix {
		for _, coord := range line {
			if coord.altitude == 0 {
				pathLengths = append(pathLengths, coord.distToE)
			}
		}
	}
	return slices.Min(pathLengths)
}

func main() {
	fmt.Println("Advent-Of-Code 2022 - Day12")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
