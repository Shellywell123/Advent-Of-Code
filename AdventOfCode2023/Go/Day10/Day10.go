package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type coord struct {
	x int
	y int
}

func Parse(filename string) ([][]rune, coord) {
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
			if line[x] == 'S' {
				start.x = x
				start.y = y
			}
		}
		y++
		grid = append(grid, gridLine)
	}

	return grid, start
}

func StartRune(grid [][]rune, start coord) (rune, coord) {
	tops := []rune{'|', '7', 'F'}
	bottoms := []rune{'|', 'J', 'L'}
	lefts := []rune{'L', 'F', '-'}
	rights := []rune{'7', 'J', '-'}

	if start.x > 0 && slices.Contains(bottoms, grid[start.y+1][start.x]) && slices.Contains(lefts, grid[start.y][start.x-1]) {
		return '7', coord{start.x - 1, start.y}
	}
	if start.y > 0 && slices.Contains(tops, grid[start.y-1][start.x]) && slices.Contains(rights, grid[start.y][start.x+1]) {
		return 'L', coord{start.x + 1, start.y}
	}
	if start.x > 0 && start.y > 0 && slices.Contains(tops, grid[start.y-1][start.x]) && slices.Contains(lefts, grid[start.y][start.x-1]) {
		return 'J', coord{start.x - 1, start.y}
	}
	if slices.Contains(bottoms, grid[start.y+1][start.x]) && slices.Contains(rights, grid[start.y][start.x+1]) {
		return 'F', coord{start.x + 1, start.y}
	}
	if start.y > 0 && slices.Contains(tops, grid[start.y-1][start.x]) && slices.Contains(bottoms, grid[start.y+1][start.x]) {
		return '|', coord{start.x, start.y + 1}
	}
	if slices.Contains(lefts, grid[start.y][start.x-1]) && slices.Contains(rights, grid[start.y][start.x+1]) {
		return '-', coord{start.x + 1, start.y}
	}
	return 'S', coord{start.x, start.y}
}

func FindLoop(grid [][]rune, startCoord coord) ([][]rune, int) {
	pathGrid := [][]rune{}

	for i := 0; i < len(grid); i++ {
		pathGridLine := []rune{}
		for j := 0; j < len(grid[i]); j++ {

			pathGridLine = append(pathGridLine, '.')
		}
		pathGrid = append(pathGrid, pathGridLine)
	}

	startRune, currentCoord := StartRune(grid, startCoord)
	grid[startCoord.y][startCoord.x] = startRune
	pathGrid[startCoord.y][startCoord.x] = startRune
	previousCoord := startCoord
	nextCoord := coord{}
	loopLength := 1
	for nextCoord != startCoord {
		switch grid[currentCoord.y][currentCoord.x] {
		case '|':
			if previousCoord.y == currentCoord.y+1 && previousCoord.x == currentCoord.x {
				loopLength++
				nextCoord = coord{currentCoord.x, currentCoord.y - 1}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			} else if previousCoord.y == currentCoord.y-1 && previousCoord.x == currentCoord.x {
				previousCoord = currentCoord
				loopLength++
				nextCoord = coord{currentCoord.x, currentCoord.y + 1}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			}
		case '-':
			if previousCoord.x == currentCoord.x+1 && previousCoord.y == currentCoord.y {
				loopLength++
				nextCoord = coord{currentCoord.x - 1, currentCoord.y}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			} else if previousCoord.x == currentCoord.x-1 && previousCoord.y == currentCoord.y {
				loopLength++
				nextCoord = coord{currentCoord.x + 1, currentCoord.y}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			}
		case 'L':
			if previousCoord.y == currentCoord.y-1 && previousCoord.x == currentCoord.x {
				loopLength++
				nextCoord = coord{currentCoord.x + 1, currentCoord.y}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			} else if previousCoord.x == currentCoord.x+1 && previousCoord.y == currentCoord.y {
				loopLength++
				nextCoord = coord{currentCoord.x, currentCoord.y - 1}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			}
		case 'J':
			if previousCoord.x == currentCoord.x-1 && previousCoord.y == currentCoord.y {
				loopLength++
				nextCoord = coord{currentCoord.x, currentCoord.y - 1}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			} else if previousCoord.y == currentCoord.y-1 && previousCoord.x == currentCoord.x {
				loopLength++
				previousCoord = currentCoord
				nextCoord = coord{currentCoord.x - 1, currentCoord.y}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				currentCoord = nextCoord
				continue
			}
		case '7':
			if previousCoord.x == currentCoord.x-1 && previousCoord.y == currentCoord.y {
				loopLength++
				nextCoord = coord{currentCoord.x, currentCoord.y + 1}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			} else if previousCoord.y == currentCoord.y+1 && previousCoord.x == currentCoord.x {
				loopLength++
				nextCoord = coord{currentCoord.x - 1, currentCoord.y}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			}
		case 'F':
			if previousCoord.x == currentCoord.x+1 && previousCoord.y == currentCoord.y {
				loopLength++
				nextCoord = coord{currentCoord.x, currentCoord.y + 1}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			} else if previousCoord.y == currentCoord.y+1 && previousCoord.x == currentCoord.x {
				loopLength++
				nextCoord = coord{currentCoord.x + 1, currentCoord.y}
				pathGrid[currentCoord.y][currentCoord.x] = grid[currentCoord.y][currentCoord.x]
				previousCoord = currentCoord
				currentCoord = nextCoord
				continue
			}
		}
	}
	return pathGrid, loopLength
}

func MarkEscaped(grid [][]rune, start coord, path []coord, changes int) ([][]rune, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'E' {
				nextCoords := []coord{}
				if j > 0 {
					nextCoords = append(nextCoords, coord{j - 1, i})
				}
				if j < len(grid[0])-1 {
					nextCoords = append(nextCoords, coord{j + 1, i})
				}
				if i > 0 {
					nextCoords = append(nextCoords, coord{j, i - 1})
				}
				if i < len(grid)-1 {
					nextCoords = append(nextCoords, coord{j, i + 1})
				}
				if j > 0 && i < len(grid)-1 {
					nextCoords = append(nextCoords, coord{j - 1, i + 1})
				}
				if j < len(grid[0])-1 && i > 0 {
					nextCoords = append(nextCoords, coord{j + 1, i - 1})
				}
				if j > 0 && i > 0 {
					nextCoords = append(nextCoords, coord{j - 1, i - 1})
				}
				if j < len(grid[0])-1 && i < len(grid)-1 {
					nextCoords = append(nextCoords, coord{j + 1, i + 1})
				}

				for _, nextCoord := range nextCoords {
					if grid[nextCoord.y][nextCoord.x] == '.' {
						grid[nextCoord.y][nextCoord.x] = 'E'
						changes++
					}
				}
			}
		}
	}

	return grid, changes
}

func Part1(filename string) int {
	grid, startCoord := Parse(filename)
	_, loopLength := FindLoop(grid, startCoord)
	return loopLength / 2
}

func Part2(filename string) int {

	grid, startCoord := Parse(filename)
	pathGrid, _ := FindLoop(grid, startCoord)

	// pad grid with border
	paddedPathGrid := [][]rune{}
	borderPaddedPathGridRow := []rune{}
	for j := 0; j < len(pathGrid[0])*2+2; j++ {
		borderPaddedPathGridRow = append(borderPaddedPathGridRow, 'E')
	}
	paddedPathGrid = append(paddedPathGrid, borderPaddedPathGridRow)

	for i := 0; i < len(pathGrid); i++ {
		paddedPathGridRow := []rune{}
		paddedPathGridRow2 := []rune{}
		paddedPathGridRow = append(paddedPathGridRow, 'E')
		paddedPathGridRow2 = append(paddedPathGridRow2, 'E')
		for j := 0; j < len(pathGrid[i]); j++ {

			switch pathGrid[i][j] {
			case '.':
				paddedPathGridRow = append(paddedPathGridRow, '.')
				paddedPathGridRow = append(paddedPathGridRow, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			case '-':
				paddedPathGridRow = append(paddedPathGridRow, '-')
				paddedPathGridRow = append(paddedPathGridRow, '-')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			case '|':
				paddedPathGridRow = append(paddedPathGridRow, '|')
				paddedPathGridRow = append(paddedPathGridRow, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '|')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			case 'L':
				paddedPathGridRow = append(paddedPathGridRow, '┕')
				paddedPathGridRow = append(paddedPathGridRow, '-')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			case 'F':
				paddedPathGridRow = append(paddedPathGridRow, '┏')
				paddedPathGridRow = append(paddedPathGridRow, '-')
				paddedPathGridRow2 = append(paddedPathGridRow2, '|')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			case '7':
				paddedPathGridRow = append(paddedPathGridRow, '┓')
				paddedPathGridRow = append(paddedPathGridRow, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '|')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			case 'J':
				paddedPathGridRow = append(paddedPathGridRow, '┛')
				paddedPathGridRow = append(paddedPathGridRow, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
				paddedPathGridRow2 = append(paddedPathGridRow2, '.')
			}

		}
		paddedPathGridRow = append(paddedPathGridRow, 'E')
		paddedPathGridRow2 = append(paddedPathGridRow2, 'E')
		paddedPathGrid = append(paddedPathGrid, paddedPathGridRow)
		paddedPathGrid = append(paddedPathGrid, paddedPathGridRow2)
	}
	paddedPathGrid = append(paddedPathGrid, borderPaddedPathGridRow)

	changes := 1
	for changes > 0 {
		paddedPathGrid, changes = MarkEscaped(paddedPathGrid, coord{0, 0}, []coord{}, 0)
	}

	ans := 0
	for i := 0; i < len(paddedPathGrid); i++ {
		for j := 0; j < len(paddedPathGrid[i]); j++ {
			if paddedPathGrid[i][j] == '.' && j%2 != 0 && i%2 != 0 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day10")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests2.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
