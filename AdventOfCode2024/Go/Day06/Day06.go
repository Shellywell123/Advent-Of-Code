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

func parse(filename string) ([][]rune, coord) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	start := coord{}
	y := 0

	grid := [][]rune{}
	for fscanner.Scan() {
		line := fscanner.Text()

		gridLine := []rune{}
		for x := 0; x < len(line); x++ {
			gridLine = append(gridLine, rune(line[x]))
			if line[x] == '^' {
				start.x = x
				start.y = y
			}
		}
		y++
		grid = append(grid, gridLine)
	}

	return grid, start
}

func isEnd(grid [][]rune, coord coord) bool {
	if (coord.x == 0 && grid[coord.y][coord.x] == '<') ||
		(coord.x >= len(grid[0])-1 && grid[coord.y][coord.x] == '>') ||
		(coord.y == 0 && grid[coord.y][coord.x] == '^') ||
		(coord.y >= len(grid)-1 && grid[coord.y][coord.x] == 'v') {
		return true
	}
	return false
}

func loopDetected(startCoord coord, grid [][]rune) bool {

	// check if at loop ^
	// . . .
	// . # .
	//.  ^ .
	// . ^ .
	if grid[startCoord.y][startCoord.x] == '^' {
		nextCoord := coord{startCoord.x, startCoord.y - 1}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == '^' {
				nextNextCoord := coord{startCoord.x, startCoord.y - 2}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						return true
					}
				}
			}
		}
	}

	// check if at loop v
	// . . .
	// . v .
	//.  v .
	// . # .
	if grid[startCoord.y][startCoord.x] == 'v' {
		nextCoord := coord{startCoord.x, startCoord.y + 1}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == 'v' {
				nextNextCoord := coord{startCoord.x, startCoord.y + 2}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						return true
					}
				}
			}
		}
	}

	// check if at loop >
	// . . . .
	// . . . .
	//.  > > #
	// . . . .
	if grid[startCoord.y][startCoord.x] == '>' {
		nextCoord := coord{startCoord.x + 1, startCoord.y}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == '>' {
				nextNextCoord := coord{startCoord.x + 2, startCoord.y}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						return true
					}
				}
			}
		}
	}

	// check if at loop <
	// . . . .
	// . . . .
	//.  # < <
	// . . . .
	if grid[startCoord.y][startCoord.x] == '<' {
		nextCoord := coord{startCoord.x - 1, startCoord.y}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == '<' {
				nextNextCoord := coord{startCoord.x - 2, startCoord.y}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						return true
					}
				}
			}
		}
	}

	// check if at loop ^>v
	// . . . .  . . . .  . . . .
	// . # . .  . # . .  . # . .
	// . ^ # .  . > # .  . v # .
	// . ^ . .  . ^ . .  . v . .
	if grid[startCoord.y][startCoord.x] == '^' {
		nextCoord := coord{startCoord.x, startCoord.y - 1}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == 'V' {
				nextNextCoord := coord{startCoord.x, startCoord.y - 2}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						nextNextNextCoord := coord{startCoord.x + 1, startCoord.y - 1}
						if isEnd(grid, nextNextNextCoord) {
							return false
						} else {
							if grid[nextNextNextCoord.y][nextNextNextCoord.x] == '#' {
								return true
							}
						}
					}
				}
			}
		}
	}

	// check if at loop >v<
	// . . . .  . . . .  . . . .
	// > > # .  > v # .  < < # .
	// . # . .  . # . .  . # . .
	if grid[startCoord.y][startCoord.x] == '>' {
		nextCoord := coord{startCoord.x + 1, startCoord.y}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == '<' {
				nextNextCoord := coord{startCoord.x + 2, startCoord.y}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						nextNextNextCoord := coord{startCoord.x + 1, startCoord.y + 1}
						if isEnd(grid, nextNextNextCoord) {
							return false
						} else {
							if grid[nextNextNextCoord.y][nextNextNextCoord.x] == '#' {
								return true
							}
						}
					}
				}
			}
		}
	}

	// check if at loop v<^
	// . v . .  . v . .  . ^ . .
	// # v . .  # < . .  # ^ . .
	// . # . .  . # . .  . # . .
	// . . . .  . v . .  . . . .
	if grid[startCoord.y][startCoord.x] == 'v' {
		nextCoord := coord{startCoord.x, startCoord.y + 1}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == '^' {
				nextNextCoord := coord{startCoord.x, startCoord.y + 2}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						nextNextNextCoord := coord{startCoord.x - 1, startCoord.y + 1}
						if isEnd(grid, nextNextNextCoord) {
							return false
						} else {
							if grid[nextNextNextCoord.y][nextNextNextCoord.x] == '#' {
								return true
							}
						}
					}
				}
			}
		}
	}

	// check if at loop <^>
	// . . # .  . . # .  . . # .
	// . # < <  . # ^ <  . # > >
	// . . . .  . . . .  . . . .
	if grid[startCoord.y][startCoord.x] == '<' {
		nextCoord := coord{startCoord.x - 1, startCoord.y}
		if isEnd(grid, nextCoord) {
			return false
		} else {
			if grid[nextCoord.y][nextCoord.x] == '>' {
				nextNextCoord := coord{startCoord.x - 2, startCoord.y}
				if isEnd(grid, nextNextCoord) {
					return false
				} else {
					if grid[nextNextCoord.y][nextNextCoord.x] == '#' {
						nextNextNextCoord := coord{startCoord.x - 1, startCoord.y - 1}
						if isEnd(grid, nextNextNextCoord) {
							return false
						} else {
							if grid[nextNextNextCoord.y][nextNextNextCoord.x] == '#' {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

func moveGuard(startCoord coord, grid [][]rune, distinctCoords int, debug bool) (int, bool, [][]rune) {

	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				fmt.Printf("%c", grid[y][x])
			}
			fmt.Println()
		}
		fmt.Println()
	}

	nextCoord := coord{}

	if isEnd(grid, startCoord) {
		return distinctCoords + 1, false, grid
	}

	// req for part 2
	if loopDetected(startCoord, grid) {
		return distinctCoords, true, grid
	}

	nextChar := rune('.')
	startChar := grid[startCoord.y][startCoord.x]

	// determine guards next natural move
	// doing it the long way cba to do matrix rotations
	if startChar == '^' && grid[startCoord.y-1][startCoord.x] != '#' {
		nextCoord.y = startCoord.y - 1
		nextCoord.x = startCoord.x
		nextChar = '^'
	}

	if startChar == '^' && grid[startCoord.y-1][startCoord.x] == '#' {
		nextCoord.y = startCoord.y
		nextCoord.x = startCoord.x + 1
		nextChar = '>'
	}

	if startChar == 'v' && grid[startCoord.y+1][startCoord.x] != '#' {
		nextCoord.y = startCoord.y + 1
		nextCoord.x = startCoord.x
		nextChar = 'v'
	}

	if startChar == 'v' && grid[startCoord.y+1][startCoord.x] == '#' {
		nextCoord.y = startCoord.y
		nextCoord.x = startCoord.x - 1
		nextChar = '<'
	}

	if startChar == '>' && grid[startCoord.y][startCoord.x+1] != '#' {
		nextCoord.y = startCoord.y
		nextCoord.x = startCoord.x + 1
		nextChar = '>'
	}

	if startChar == '>' && grid[startCoord.y][startCoord.x+1] == '#' {
		nextCoord.y = startCoord.y + 1
		nextCoord.x = startCoord.x
		nextChar = 'v'
	}

	if startChar == '<' && grid[startCoord.y][startCoord.x-1] != '#' {
		nextCoord.y = startCoord.y
		nextCoord.x = startCoord.x - 1
		nextChar = '<'
	}

	if startChar == '<' && grid[startCoord.y][startCoord.x-1] == '#' {
		nextCoord.y = startCoord.y - 1
		nextCoord.x = startCoord.x
		nextChar = '^'
	}

	// check for second block resulting in U turn
	//
	// . . . .
	// . # . .
	// . ^ # .
	// . ^ . .
	//
	// note - this also changes loop detection logic (took me a while to spot)

	if grid[nextCoord.y][nextCoord.x] == '#' {

		if startChar == '^' {
			nextChar = 'v'
		}
		if startChar == 'v' {
			nextChar = '^'
		}
		if startChar == '>' {
			nextChar = '<'
		}
		if startChar == '<' {
			nextChar = '>'
		}
		nextCoord = startCoord
	}

	// increment distinctCoords if next cell in new
	if grid[nextCoord.y][nextCoord.x] == '.' {
		distinctCoords++
	}
	grid[nextCoord.y][nextCoord.x] = nextChar

	return moveGuard(nextCoord, grid, distinctCoords, debug)
}

func Part1(filename string, debug bool) int {
	// side note - i started this on the train during https://www.bbc.co.uk/news/live/cv2g2rrjywqt
	grid, startCoord := parse(filename)

	ans, _, _ := moveGuard(startCoord, grid, 0, debug)

	return ans
}

func Part2(filename string, debug bool) int {
	grid, _ := parse(filename)

	ans := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {

			grid, startCoord := parse(filename)

			// trial obstruction in new coord
			if grid[y][x] == '.' {
				grid[y][x] = '#'
			} else {
				continue
			}

			if debug {
				for y := 0; y < len(grid); y++ {
					for x := 0; x < len(grid[y]); x++ {
						fmt.Printf("%c", grid[y][x])
					}
					fmt.Println()
				}
				fmt.Println()
			}

			_, loop, _ := moveGuard(startCoord, grid, 0, debug)
			if loop {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day06")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
