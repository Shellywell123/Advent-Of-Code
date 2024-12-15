package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type coord struct {
	x int
	y int
}

func parse(filename string) ([][]rune, coord, string) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	y := 0

	robotCoord := coord{}
	grid := [][]rune{}
	movements := ""
	isMovement := false
	for fscanner.Scan() {
		line := fscanner.Text()

		if line == "" {
			isMovement = true
			continue
		}

		if isMovement {
			movements += line
			continue
		}

		gridLine := []rune{}
		for x := 0; x < len(line); x++ {
			gridLine = append(gridLine, rune(line[x]))
			for x := 0; x < len(line); x++ {
				if line[x] == '@' {
					robotCoord = coord{x, y}
				}
			}
		}
		y++
		grid = append(grid, gridLine)
	}

	return grid, robotCoord, movements
}

func moveRobot(grid [][]rune, robotCoord coord, movements string, part int, debug bool) ([][]rune, coord) {

	if len(movements) == 0 {
		return grid, robotCoord
	}

	if debug {
		printGrid(grid)
		fmt.Println(robotCoord)
		fmt.Println(len(movements), string(movements[0]))
	}

	movement := movements[0]

	switch movement {
	case '<':
		if grid[robotCoord.y][robotCoord.x-1] == '#' {
			return moveRobot(grid, robotCoord, movements[1:], part, debug)
		}

		moveBox := false
		gap := coord{}
		for x := robotCoord.x; x >= 0; x-- {
			if grid[robotCoord.y][x] == '.' {
				moveBox = true
				gap = coord{x, robotCoord.y}
				break
			}
			if grid[robotCoord.y][x] == '#' {
				break
			}
		}

		if moveBox && part == 1 {
			grid[gap.y][gap.x] = 'O'
			grid[robotCoord.y][robotCoord.x] = '.'
			robotCoord.x--
			grid[robotCoord.y][robotCoord.x] = '@'
		}

		if moveBox && part == 2 {
			for x := gap.x; x <= robotCoord.x-2; x++ {
				grid[robotCoord.y][x] = grid[robotCoord.y][x+1]
			}
			grid[robotCoord.y][robotCoord.x] = '.'
			robotCoord.x--
			grid[robotCoord.y][robotCoord.x] = '@'
		}
		break

	case '>':
		if grid[robotCoord.y][robotCoord.x+1] == '#' {
			return moveRobot(grid, robotCoord, movements[1:], part, debug)
		}

		moveBox := false
		gap := coord{}
		for x := robotCoord.x; x < len(grid[0]); x++ {
			if grid[robotCoord.y][x] == '.' {
				moveBox = true
				gap = coord{x, robotCoord.y}
				break
			}
			if grid[robotCoord.y][x] == '#' {
				break
			}
		}

		if moveBox && part == 1 {
			grid[gap.y][gap.x] = 'O'
			grid[robotCoord.y][robotCoord.x] = '.'
			robotCoord.x++
			grid[robotCoord.y][robotCoord.x] = '@'
			grid[robotCoord.y][robotCoord.x-1] = '.'
		}

		if moveBox && part == 2 {
			for x := gap.x; x >= robotCoord.x+2; x-- {
				grid[robotCoord.y][x] = grid[robotCoord.y][x-1]
			}
			grid[robotCoord.y][robotCoord.x] = '.'
			robotCoord.x++
			grid[robotCoord.y][robotCoord.x] = '@'
			grid[robotCoord.y][robotCoord.x-1] = '.'
		}
		break

	case '^':
		if grid[robotCoord.y-1][robotCoord.x] == '#' {
			return moveRobot(grid, robotCoord, movements[1:], part, debug)
		}

		moveBox := false
		gap := coord{}
		for y := robotCoord.y; y >= 0; y-- {
			if grid[y][robotCoord.x] == '.' {
				moveBox = true
				gap = coord{robotCoord.x, y}
				break
			}
			if grid[y][robotCoord.x] == '#' {
				break
			}
		}

		// normal map
		if moveBox && part == 1 {
			grid[gap.y][gap.x] = 'O'
			grid[robotCoord.y][robotCoord.x] = '.'
			robotCoord.y--
			grid[robotCoord.y][robotCoord.x] = '@'
		}

		// wide map both sides of box
		if moveBox && part == 2 {
			can, boxesToMove := canItMove("up", grid, coord{robotCoord.x, robotCoord.y}, []coord{})
			if can {
				grid = moveTriangle("up", grid, boxesToMove)

				grid[robotCoord.y][robotCoord.x] = '.'
				robotCoord.y--
				grid[robotCoord.y][robotCoord.x] = '@'
			}
		}
		break

	case 'v':
		if grid[robotCoord.y+1][robotCoord.x] == '#' {
			return moveRobot(grid, robotCoord, movements[1:], part, debug)
		}

		moveBox := false
		gap := coord{}
		for y := robotCoord.y; y < len(grid); y++ {
			if grid[y][robotCoord.x] == '.' {
				moveBox = true
				gap = coord{robotCoord.x, y}
				break
			}
			if grid[y][robotCoord.x] == '#' {
				break
			}
		}

		// normal map
		if moveBox && part == 1 {
			grid[gap.y][gap.x] = 'O'
			grid[robotCoord.y][robotCoord.x] = '.'
			robotCoord.y++
			grid[robotCoord.y][robotCoord.x] = '@'
		}

		// wide map both sides of box
		if moveBox && part == 2 {
			can, boxesToMove := canItMove("down", grid, coord{robotCoord.x, robotCoord.y}, []coord{})

			if can {
				grid = moveTriangle("down", grid, boxesToMove)
				grid[robotCoord.y][robotCoord.x] = '.'
				robotCoord.y++
				grid[robotCoord.y][robotCoord.x] = '@'
			}
		}
		break
	}

	return moveRobot(grid, robotCoord, movements[1:], part, debug)
}

func canItMove(dir string, grid [][]rune, boxCoord coord, boxesToMove []coord) (bool, []coord) {

	c := 1
	if dir == "up" {
		c = -1
	}

	if grid[boxCoord.y][boxCoord.x] == '@' {
		if grid[boxCoord.y+c][boxCoord.x] == '.' {
			return true, boxesToMove
		}

		if grid[boxCoord.y+c][boxCoord.x] == '#' {
			return false, boxesToMove
		}

		if grid[boxCoord.y+c][boxCoord.x] == '[' {
			return canItMove(dir, grid, coord{boxCoord.x, boxCoord.y + c}, boxesToMove)
		}

		if grid[boxCoord.y+c][boxCoord.x] == ']' {
			return canItMove(dir, grid, coord{boxCoord.x - 1, boxCoord.y + c}, boxesToMove)
		}
	}

	if grid[boxCoord.y][boxCoord.x] == '[' {

		newBox := true
		for _, b := range boxesToMove {
			if b == boxCoord {
				newBox = false
			}
		}
		if newBox && grid[boxCoord.y][boxCoord.x] == '[' {
			boxesToMove = append(boxesToMove, boxCoord)
			boxesToMove = append(boxesToMove, coord{boxCoord.x + 1, boxCoord.y})
		}

		// check if there is a gap above the box
		if grid[boxCoord.y+c][boxCoord.x] == '.' && grid[boxCoord.y+c][boxCoord.x+1] == '.' {
			return true, boxesToMove
		}

		// check if there is a wall above the box
		if grid[boxCoord.y+c][boxCoord.x] == '#' || grid[boxCoord.y+c][boxCoord.x+1] == '#' {
			return false, boxesToMove
		}

		// above box is not offset
		if grid[boxCoord.y+c][boxCoord.x] == '[' {
			return canItMove(dir, grid, coord{boxCoord.x, boxCoord.y + c}, boxesToMove)
		}

		// above box is offset
		// ][
		if grid[boxCoord.y+c][boxCoord.x] == ']' && grid[boxCoord.y+c][boxCoord.x+1] == '[' {
			left, boxesToMove := canItMove(dir, grid, coord{boxCoord.x - 1, boxCoord.y + c}, boxesToMove)
			right, boxesToMove := canItMove(dir, grid, coord{boxCoord.x + 1, boxCoord.y + c}, boxesToMove)
			if !left || !right {
				return false, boxesToMove
			}
			if left && right {
				return true, boxesToMove
			}
		}

		// ].
		if grid[boxCoord.y+c][boxCoord.x] == ']' && grid[boxCoord.y+c][boxCoord.x+1] == '.' {
			return canItMove(dir, grid, coord{boxCoord.x - 1, boxCoord.y + c}, boxesToMove)
		}

		// .[
		if grid[boxCoord.y+c][boxCoord.x] == '.' && grid[boxCoord.y+c][boxCoord.x+1] == '[' {
			return canItMove(dir, grid, coord{boxCoord.x + 1, boxCoord.y + c}, boxesToMove)
		}
	}

	if grid[boxCoord.y][boxCoord.x] == ']' {

		newBox := true
		for _, b := range boxesToMove {
			if b == boxCoord {
				newBox = false
			}
		}
		if newBox && grid[boxCoord.y][boxCoord.x] == ']' {
			boxesToMove = append(boxesToMove, boxCoord)
			boxesToMove = append(boxesToMove, coord{boxCoord.x - 1, boxCoord.y})
		}

		// check if there is a gap above the box
		if grid[boxCoord.y+c][boxCoord.x] == '.' && grid[boxCoord.y+c][boxCoord.x-1] == '.' {
			return true, boxesToMove
		}

		// check if there is a wall above the box
		if grid[boxCoord.y+c][boxCoord.x] == '#' || grid[boxCoord.y+c][boxCoord.x-1] == '#' {
			return false, boxesToMove
		}

		// above box is not offset
		if grid[boxCoord.y+c][boxCoord.x] == ']' {
			return canItMove(dir, grid, coord{boxCoord.x, boxCoord.y + c}, boxesToMove)
		}

		// above box is offset
		// ][
		if grid[boxCoord.y+c][boxCoord.x] == '[' && grid[boxCoord.y+c][boxCoord.x-1] == ']' {
			left, boxesToMove := canItMove(dir, grid, coord{boxCoord.x - 1, boxCoord.y + c}, boxesToMove)
			right, boxesToMove := canItMove(dir, grid, coord{boxCoord.x + 1, boxCoord.y + c}, boxesToMove)
			if !left || !right {
				return false, boxesToMove
			}
			if left && right {
				return true, boxesToMove
			}
		}

		// .]
		if grid[boxCoord.y+c][boxCoord.x] == '.' && grid[boxCoord.y+c][boxCoord.x-1] == ']' {
			return canItMove(dir, grid, coord{boxCoord.x + 1, boxCoord.y + c}, boxesToMove)
		}

		// [.
		if grid[boxCoord.y+c][boxCoord.x] == '[' && grid[boxCoord.y+c][boxCoord.x-1] == '.' {
			return canItMove(dir, grid, coord{boxCoord.x - 1, boxCoord.y + c}, boxesToMove)
		}
	}

	return false, boxesToMove
}

func moveTriangle(dir string, grid [][]rune, boxesToMove []coord) [][]rune {

	// make a copy
	newGrid := [][]rune{}
	for y := 0; y < len(grid); y++ {
		newLine := []rune{}
		for x := 0; x < len(grid[0]); x++ {
			newLine = append(newLine, grid[y][x])
		}
		newGrid = append(newGrid, newLine)
	}

	if dir == "up" {
		sort.Slice(boxesToMove, func(i, j int) bool {
			return boxesToMove[i].y < boxesToMove[j].y
		})

		for _, boxCoord := range boxesToMove {
			newGrid[boxCoord.y][boxCoord.x] = '.'
			newGrid[boxCoord.y-1][boxCoord.x] = grid[boxCoord.y][boxCoord.x]
		}

	} else {
		sort.Slice(boxesToMove, func(i, j int) bool {
			return boxesToMove[i].y > boxesToMove[j].y
		})

		for _, boxCoord := range boxesToMove {
			newGrid[boxCoord.y][boxCoord.x] = '.'
			newGrid[boxCoord.y+1][boxCoord.x] = grid[boxCoord.y][boxCoord.x]
		}
	}

	return newGrid
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		for _, r := range line {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func boxTotal(grid [][]rune) int {
	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' || grid[y][x] == '[' {
				total += x + (y * 100)
			}
		}
	}
	return total
}

func widenGrid(grid [][]rune) ([][]rune, coord) {
	newGrid := [][]rune{}
	for y := 0; y < len(grid); y++ {
		newLine := []rune{}
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				newLine = append(newLine, '#')
				newLine = append(newLine, '#')
			}
			if grid[y][x] == '.' {
				newLine = append(newLine, '.')
				newLine = append(newLine, '.')
			}
			if grid[y][x] == 'O' {
				newLine = append(newLine, '[')
				newLine = append(newLine, ']')
			}
			if grid[y][x] == '@' {
				newLine = append(newLine, '@')
				newLine = append(newLine, '.')
			}
		}
		newGrid = append(newGrid, newLine)
	}

	robotCoord := coord{}
	for y := 0; y < len(newGrid); y++ {
		for x := 0; x < len(newGrid[0]); x++ {
			if newGrid[y][x] == '@' {
				robotCoord = coord{x, y}
			}
		}
	}

	return newGrid, robotCoord
}

func Part1(filename string, debug bool) int {
	grid, robotCoord, movements := parse(filename)

	if debug {
		printGrid(grid)
		fmt.Println(robotCoord)
		fmt.Println(movements)
		fmt.Println()
	}

	grid, _ = moveRobot(grid, robotCoord, movements, 1, debug)
	if debug {
		printGrid(grid)
	}

	ans := boxTotal(grid)
	return ans
}

func Part2(filename string, debug bool) int {
	grid, robotCoord, movements := parse(filename)
	grid, robotCoord = widenGrid(grid)

	if debug {
		printGrid(grid)
		fmt.Println(robotCoord)
		fmt.Println(movements)
		fmt.Println()
	}

	grid, _ = moveRobot(grid, robotCoord, movements, 2, debug)

	
	if debug{
		printGrid(grid)
	}
	ans := boxTotal(grid)
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day15")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
