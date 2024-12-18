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

var size int

func parse(filename string) ([][]rune, coord) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	y := 0

	startCoord := coord{}
	grid := [][]rune{}
	for fscanner.Scan() {
		line := fscanner.Text()

		gridLine := []rune{}
		for x := 0; x < len(line); x++ {
			gridLine = append(gridLine, rune(line[x]))
			for x := 0; x < len(line); x++ {
				if line[x] == 'S' {
					startCoord = coord{x, y}
				}
			}
		}
		y++
		grid = append(grid, gridLine)
	}

	return grid, startCoord
}

func printGrid(grid [][]rune) {
	for y, line := range grid {
		for _, r := range line {
			fmt.Print(string(r))
		}
		fmt.Print(" ", y)
		fmt.Println()
	}
	for x := 0; x < 10; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println()
}

func printPath(path []coord, grid [][]rune) int {

	newGrid := [][]rune{}
	for y := 0; y < len(grid); y++ {
		newGridLine := []rune{}
		for x := 0; x < len(grid[y]); x++ {
			newGridLine = append(newGridLine, grid[y][x])
		}
		newGrid = append(newGrid, newGridLine)
	}

	for i := 0; i < len(path)-1; i++ {
		if path[i].y == path[i+1].y && path[i].x < path[i+1].x {
			newGrid[path[i].y][path[i].x] = '>'
		}
		if path[i].y == path[i+1].y && path[i].x > path[i+1].x {
			newGrid[path[i].y][path[i].x] = '<'
		}
		if path[i].x == path[i+1].x && path[i].y < path[i+1].y {
			newGrid[path[i].y][path[i].x] = 'v'
		}
		if path[i].x == path[i+1].x && path[i].y > path[i+1].y {
			newGrid[path[i].y][path[i].x] = '^'
		}
	}
	printGrid(newGrid)
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func alreadyBeenHere(path []coord, c coord) bool {
	for _, p := range path {
		if p == c {
			return true
		}
	}
	return false
}

func whoLetTheReindeersOut(grid [][]rune, startOrientation rune, path []coord, currentAns int, values map[string]int, debug bool) int {

	startCoord := path[len(path)-1]
	v := valuePath(path)
	startRune := grid[startCoord.y][startCoord.x]

	// if we have reached the end, return the answer
	if startRune == 'E' {

		if debug {
			fmt.Println("END!", v, startCoord, path[len(path)-1], path[len(path)-1])
		}

		return v
	}

	// don't bother continuing if we have a been here before for cheaper
	key := fmt.Sprintf("%d-%d", startCoord.x, startCoord.y)
	_, ok := values[key]
	// adding the && here somehow sped up everything?
	// guessing map searches are expensive
	if ok && values[key] < v {
		return currentAns
	} else {
		values[key] = v
		if debug {
			fmt.Println(len(values), "/", size)
		}
	}

	// recursive check for each direction
	upCoord := coord{startCoord.x, startCoord.y - 1}
	downCoord := coord{startCoord.x, startCoord.y + 1}
	leftCoord := coord{startCoord.x - 1, startCoord.y}
	rightCoord := coord{startCoord.x + 1, startCoord.y}

	currentAnsUp := currentAns
	currentAnsDown := currentAns
	currentAnsLeft := currentAns
	currentAnsRight := currentAns

	for _, c := range []coord{upCoord, downCoord, leftCoord, rightCoord} {

		// dont hit walls
		if grid[c.y][c.x] == '#' { 
			continue
		}

		switch c {
		case downCoord:
			switch startOrientation {
			case 'v', '>', '<':
				currentAnsDown = whoLetTheReindeersOut(grid, 'v', append(path, c), currentAns, values, debug)
				break
			}
			break

		case upCoord:
			switch startOrientation {
			case '^', '>', '<':
				currentAnsUp = whoLetTheReindeersOut(grid, '^', append(path, c), currentAns, values, debug)
				break
			}
			break

		case leftCoord:
			switch startOrientation {
			case '<', '^', 'v':
				currentAnsLeft = whoLetTheReindeersOut(grid, '<', append(path, c), currentAns, values, debug)
				break
			}
			break

		case rightCoord:
			switch startOrientation {
			case '>', '^', 'v':
				currentAnsRight = whoLetTheReindeersOut(grid, '>', append(path, c), currentAns, values, debug)
				break
			}
			break
		}
	}

	ans := 10000000000000000
	for _, c := range []int{currentAnsDown, currentAnsUp, currentAnsLeft, currentAnsRight} {
		if c == 0 {
			continue
		}
		if c < ans {
			ans = c
		}
	}

	return ans
}

// duplicated as I dont want to slow down part 1
func whoLetTheReindeersOut2(grid [][]rune, startOrientation rune, paths [][]coord, path1 []coord, currentAns int, values map[string]int, debug bool) (int, [][]coord) {

	path := make([]coord, len(path1))
	copy(path, path1)

	startCoord := path[len(path)-1]
	v := valuePath(path)
	startRune := grid[startCoord.y][startCoord.x]

	// if we have reached the end, return the answer
	if startRune == 'E' {

		paths = append(paths, path)

		if debug {
			fmt.Println("END!", v, startCoord, path[len(path)-1], path[len(path)-1])
		}

		return v, paths
	}

	// don't bother continuing if we have a been here before for cheaper
	key := fmt.Sprintf("%d-%d", startCoord.x, startCoord.y)
	_, ok := values[key]
	// adding the && here somehow sped up everything?
	// guessing map searches are expensive

	// hack to make sure we dont skip answers
	// the bigger the buffer the slower, but more accurate the answer
	if ok && values[key] < v - 1000 {
		return currentAns, paths
	} else {
		values[key] = v
		if debug {
			fmt.Println(len(values), "/", size)
		}
	}

	// recursive check for each direction
	upCoord := coord{startCoord.x, startCoord.y - 1}
	downCoord := coord{startCoord.x, startCoord.y + 1}
	leftCoord := coord{startCoord.x - 1, startCoord.y}
	rightCoord := coord{startCoord.x + 1, startCoord.y}

	currentAnsUp := currentAns
	currentAnsDown := currentAns
	currentAnsLeft := currentAns
	currentAnsRight := currentAns

	for _, c := range []coord{upCoord, downCoord, leftCoord, rightCoord} {

		// dont hit walls and dont back track
		if grid[c.y][c.x] == '#' || alreadyBeenHere(path, c) { // diff here is we prevent individual paths from looping apposed to cheapest path to each grid coord
			continue
		}

		switch c {
		case downCoord:
			switch startOrientation {
			case 'v', '>', '<':
				currentAnsDown, paths = whoLetTheReindeersOut2(grid, 'v', paths, append(path, c), currentAns, values, debug)
				break
			}
			break

		case upCoord:
			switch startOrientation {
			case '^', '>', '<':
				currentAnsUp, paths = whoLetTheReindeersOut2(grid, '^', paths, append(path, c), currentAns, values, debug)
				break
			}
			break

		case leftCoord:
			switch startOrientation {
			case '<', '^', 'v':
				currentAnsLeft, paths = whoLetTheReindeersOut2(grid, '<', paths, append(path, c), currentAns, values, debug)
				break
			}
			break

		case rightCoord:
			switch startOrientation {
			case '>', '^', 'v':
				currentAnsRight, paths = whoLetTheReindeersOut2(grid, '>', paths, append(path, c), currentAns, values, debug)
				break
			}
			break
		}
	}

	ans := 10000000000000000
	for _, c := range []int{currentAnsDown, currentAnsUp, currentAnsLeft, currentAnsRight} {
		if c == 0 {
			continue
		}
		if c < ans {
			ans = c
		}
	}

	return ans, paths
}

func valuePath(path []coord) int {
	corners := -1
	currentAxis := "x"

	// check for initial corner
	if len(path) > 2 && path[1].x == 1 {
		corners++
		currentAxis = "y"
	}

	for i := 0; i < len(path)-1; i++ {
		if currentAxis == "x" {
			if path[i].x != path[i+1].x {
				currentAxis = "y"
				corners++
				continue
			}
		}
		if currentAxis == "y" {
			if path[i].y != path[i+1].y {
				currentAxis = "x"
				corners++
				continue
			}
		}
	}

	pathValue := corners*1000 + (len(path) - 1)
	return pathValue
}

func countUniqueCoords(paths [][]coord) int {

	unique := []coord{}
	for _, path := range paths {
		for _, c1 := range path {
			if !alreadyBeenHere(unique, c1) {
				unique = append(unique, c1)
			}
		}
	}

	return len(unique)
}

func orderPathsByValue(paths [][]coord) [][]coord {
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths)-1; j++ {
			if valuePath(paths[j]) > valuePath(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
	return paths
}

func Part1(filename string, debug bool) int {
	grid, startCoord := parse(filename)

	if debug {
		printGrid(grid)
	}

	// count number of empty spaces to get max size of path
	// used to print % progress
	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == '.' {
					size++
				}
			}
		}
	}

	ans := whoLetTheReindeersOut(grid, '>', []coord{startCoord}, 0, map[string]int{}, debug)

	return ans
}

func Part2(filename string, debug bool) int {
	grid, startCoord := parse(filename)

	if debug {
		printGrid(grid)
	}

	// count number of empty spaces to get max size of path
	// used to print % progress
	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == '.' {
					size++
				}
			}
		}
	}

	_, paths := whoLetTheReindeersOut2(grid, '>', [][]coord{}, []coord{startCoord}, 0, map[string]int{}, debug)

	// determine best paths
	paths = orderPathsByValue(paths)
	bestPaths := [][]coord{}
	for p := 0; p < len(paths); p++ {
		if valuePath(paths[p]) == valuePath(paths[0]) {
			if debug {
				fmt.Println(p, valuePath(paths[p]))
				printPath(paths[p], grid)
			}
			bestPaths = append(bestPaths, paths[p])
		}
	}

	ans := countUniqueCoords(bestPaths)

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day16")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false)) // slow
}
