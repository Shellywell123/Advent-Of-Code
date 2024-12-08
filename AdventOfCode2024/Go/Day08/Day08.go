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

func parse(filename string) ([][]rune, map[string][]coord) {
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

	// collect all the antenna frequencies
	antennas := map[string][]coord{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '.' {
				continue
			}
			_, ok := antennas[string(grid[y][x])]
			if ok {
				antennas[string(grid[y][x])] = append(antennas[string(grid[y][x])], coord{x, y})
			} else {
				antennas[string(grid[y][x])] = []coord{coord{x, y}}
			}
		}
	}

	return grid, antennas
}

func Part1(filename string, debug bool) int {
	grid, antennas := parse(filename)

	ans := 0

	for k, v := range antennas {

		if debug {
			fmt.Printf("Antenna %v: %v\n", k, v)
		}

		// we require at least 2 antennas
		if len(v) < 2 {
			continue
		}

		// iterate through each combo of antennas
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {

				if debug {
					fmt.Printf("Antenna %v: %v, %v\n", k, v[i], v[j])
				}

				// create antinodes
				dx := v[i].x - v[j].x
				dy := v[i].y - v[j].y

				if v[i].x+dx >= 0 && v[i].x+dx < len(grid[0]) && v[i].y+dy >= 0 && v[i].y+dy < len(grid) {
					fmt.Println(v[j].y-dy, v[j].x-dx)
					if grid[v[i].y+dy][v[i].x+dx] != '#' {
						ans++
						grid[v[i].y+dy][v[i].x+dx] = '#'
					}
				}

				if v[j].x-dx >= 0 && v[j].x-dx < len(grid[0]) && v[j].y-dy >= 0 && v[j].y-dy < len(grid) {
					fmt.Println(v[j].y-dy, v[j].x-dx)
					if string(grid[v[j].y-dy][v[j].x-dx]) != "#" {
						ans++
						grid[v[j].y-dy][v[j].x-dx] = '#'
					}
				}
			}
		}
	}

	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				fmt.Printf("%c", grid[y][x])
			}
			fmt.Println()
		}
	}

	return ans
}


func Part2(filename string, debug bool) int {
	grid, antennas := parse(filename)

	ans := 0

	for k, v := range antennas {

		if debug {
			fmt.Printf("Antenna %v: %v\n", k, v)
		}

		// we require at least 2 antennas
		if len(v) < 2 {
			continue
		}

		// iterate through each combo of antennas
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {

				if debug {
					fmt.Printf("Antenna %v: %v, %v\n", k, v[i], v[j])
				}

				// create antinodes 

				dx := v[i].x - v[j].x
				dy := v[i].y - v[j].y

				inRange := true // make shift while loop
				ir1 := 0
				for inRange {
					if v[i].x+(ir1*dx) >= 0 && v[i].x+(ir1*dx) < len(grid[0]) && v[i].y+(ir1*dy) >= 0 && v[i].y+(ir1*dy) < len(grid) {
						if grid[v[i].y+(ir1*dy)][v[i].x+(ir1*dx)] != '#' {
							ans++
							grid[v[i].y+(ir1*dy)][v[i].x+(ir1*dx)] = '#'
						}
					ir1++
					} else {
						inRange = false
					}
				}

				inRange2 := true
				ir2 := 0
				for inRange2 {
					if v[j].x-(ir2*dx) >= 0 && v[j].x-(ir2*dx) < len(grid[0]) && v[j].y-(ir2*dy) >= 0 && v[j].y-(ir2*dy) < len(grid) {
						if string(grid[v[j].y-(ir2*dy)][v[j].x-(ir2*dx)]) != "#" {
							ans++
							grid[v[j].y-(ir2*dy)][v[j].x-(ir2*dx)] = '#'
						}
					ir2++
					} else {
						inRange2 = false
					}
				}
			}
		}
	}

	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				fmt.Printf("%c", grid[y][x])
			}
			fmt.Println()
		}
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day08")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
