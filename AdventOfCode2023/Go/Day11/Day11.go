package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"math"
	"strconv"
	"strings"
)

type galaxy struct {
	number int
	x int
	y int
}

func Parse(filename string) [][]rune {
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
		grid = append(grid, gridLine)
		y++
		
	}

	return grid
}

func CosmicExpansion(galaxies map[int]galaxy, grid [][]rune, expansion int) map[int]galaxy {
	galaxiesExpanded := map[int]galaxy{}

	for _, galaxyy := range galaxies {

		xExpansions := 0
		yExpansions := 0

		for x := 0; x < galaxyy.x; x++ {

			xExpand := true
			for yCheck := 0; yCheck < len(grid); yCheck++ {
				if grid[yCheck][x] == '#' {
					xExpand = false
					break
				}
			}

			if xExpand {
				xExpansions ++
			}
		}

		for y := 0; y < galaxyy.y; y++ {
			yExpand := true
			for xCheck := 0; xCheck < len(grid[0]); xCheck++ {
				if grid[y][xCheck] == '#' {
					yExpand = false
					break
				}
			}

			if yExpand {
				yExpansions ++
			}
		}
		
		galaxiesExpanded[galaxyy.number] = galaxy{galaxyy.number, galaxies[galaxyy.number].x + (xExpansions * (expansion-1)),  galaxies[galaxyy.number].y + (yExpansions * (expansion-1))}
	} 
	
	return galaxiesExpanded
}

func FindGalaxies(grid [][]rune) map[int]galaxy {

	number := 1
	galaxies := map[int]galaxy{}

	for y:= 0; y < len(grid); y ++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				galaxies[number] = galaxy{number, x, y}
				number++
			}
		}
	}
	return galaxies
}

func Pairs(galaxies map[int]galaxy) []string {

	pairs := []string{}
	pair := ""

	for i := 1; i <= len(galaxies); i++ {
		for j := 1; j <= len(galaxies); j++ {
			if j != i {
				if i < j {
					pair = fmt.Sprintf("%d-%d", i, j)
				} else {
					pair = fmt.Sprintf("%d-%d", j, i)
				}
				if !slices.Contains(pairs, pair) {
					pairs = append(pairs, pair)
				}
			}
		}
	}
	return pairs
}

func ShortestPaths(grid [][]rune, pairs []string, galaxies map[int]galaxy) []int {

	distances := []int{}

	for _, pair := range pairs {

		galaxiess := strings.Split(pair, "-")
		galaxyA, _:= strconv.Atoi(galaxiess[0])
		galaxyB, _ := strconv.Atoi(galaxiess[1])

		xA := galaxies[galaxyA].x
		yA := galaxies[galaxyA].y
		xB := galaxies[galaxyB].x
		yB := galaxies[galaxyB].y

		distance := int(math.Abs(float64(xB - xA)) + math.Abs(float64(yB - yA)))
		distances = append(distances, distance)
	}
	return distances
}

func Solve(filename string, expansion int) int {
	grid := (Parse(filename))
	galaxiesExpanded := CosmicExpansion(FindGalaxies(grid), grid, expansion)

	ans := 0
	for _, path := range ShortestPaths(grid,  Pairs(galaxiesExpanded), galaxiesExpanded) {
		ans += path
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day11")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Solve("tests.txt",2))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Solve("inputs.txt",2))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Solve("tests.txt",10))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Solve("inputs.txt",1000000))
}
