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

type region struct {
	name      string
	coords    []coord
	perimeter int
	area      int
	sides     int
}

func parse(filename string) [][]rune {
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

	return grid
}

func calculateSides(region region, grid [][]rune) int {

	// quick wins
	if len(region.coords) == 1 {
		return 4
	}

	// my approach
	// scanning through the coords from each direction
	// if a edge coord has a preceding coord out of the region it is part of a side
	// group each slice scan by shared x or y will give the number of sides in that slice
	// could optimize by cleaning up the duplicate logic

	topSides := 0
	bottomSides := 0
	leftSides := 0
	rightSides := 0

	// top down
	for y := 0; y < len(grid); y++ {

		// for each vertical slice check if we have any region coords
		regionCoords := []coord{}
		for _, coord := range region.coords {
			if coord.y == y {
				regionCoords = append(regionCoords, coord)
			}
		}

		// skip empty slices
		if len(regionCoords) == 0 {
			continue
		}

		// check if coord is a edge coord
		edgeCoordsX := []int{}
		for _, rc := range regionCoords {
			if rc.y == 0 || grid[rc.y-1][rc.x] != rune(region.name[0]) {
				// top edge
				edgeCoordsX = append(edgeCoordsX, rc.x)
			}
		}

		// skip empty slices
		if len(edgeCoordsX) == 0 {
			continue
		}

		// sort edge coords by common x
		sort.Ints(edgeCoordsX)
		for i := 0; i < len(edgeCoordsX); i++ {
			if i == 0 {
				topSides++
				continue
			}
			if edgeCoordsX[i]-1 != edgeCoordsX[i-1] {
				topSides++
			}
		}
	}

	// left to right
	for x := 0; x < len(grid[0]); x++ {

		// for each horizontal slice check if we have any region coords
		regionCoords := []coord{}
		for _, coord := range region.coords {
			if coord.x == x {
				regionCoords = append(regionCoords, coord)
			}
		}

		// skip empty slices
		if len(regionCoords) == 0 {
			continue
		}

		// check if coord is a edge coord
		edgeCoordsY := []int{}
		for _, rc := range regionCoords {
			if rc.x == 0 || grid[rc.y][rc.x-1] != rune(region.name[0]) {
				// left edge
				edgeCoordsY = append(edgeCoordsY, rc.y)
			}
		}

		// skip empty slices
		if len(edgeCoordsY) == 0 {
			continue
		}

		// sort edge coords by common y
		sort.Ints(edgeCoordsY)
		for i := 0; i < len(edgeCoordsY); i++ {
			if i == 0 {
				leftSides++
				continue
			}
			if edgeCoordsY[i]-1 != edgeCoordsY[i-1] {
				leftSides++
			}
		}
	}

	// right to left
	for x := len(grid[0]) - 1; x >= 0; x-- {

		// for each horizontal slice check if we have any region coords
		regionCoords := []coord{}
		for _, coord := range region.coords {
			if coord.x == x {
				regionCoords = append(regionCoords, coord)
			}
		}

		// skip empty slices
		if len(regionCoords) == 0 {
			continue
		}

		// check if coord is a edge coord
		edgeCoordsY := []int{}
		for _, rc := range regionCoords {
			if rc.x == len(grid[0])-1 || grid[rc.y][rc.x+1] != rune(region.name[0]) {
				// right edge
				edgeCoordsY = append(edgeCoordsY, rc.y)
			}
		}

		// skip empty slices
		if len(edgeCoordsY) == 0 {
			continue
		}

		// sort edge coords by common y
		sort.Ints(edgeCoordsY)
		for i := 0; i < len(edgeCoordsY); i++ {
			if i == 0 {
				rightSides++
				continue
			}
			if edgeCoordsY[i]-1 != edgeCoordsY[i-1] {
				rightSides++
			}
		}
	}

	// bottom up
	for y := len(grid) - 1; y >= 0; y-- {

		// for each vertical slice check if we have any region coords
		regionCoords := []coord{}
		for _, coord := range region.coords {
			if coord.y == y {
				regionCoords = append(regionCoords, coord)
			}
		}

		// skip empty slices
		if len(regionCoords) == 0 {
			continue
		}

		// check if coord is a edge coord
		edgeCoordsX := []int{}
		for _, rc := range regionCoords {
			if rc.y == len(grid)-1 || grid[rc.y+1][rc.x] != rune(region.name[0]) {
				// bottom edge
				edgeCoordsX = append(edgeCoordsX, rc.x)
			}
		}

		// skip empty slices
		if len(edgeCoordsX) == 0 {
			continue
		}

		// sort edge coords by common x
		sort.Ints(edgeCoordsX)
		for i := 0; i < len(edgeCoordsX); i++ {
			if i == 0 {
				bottomSides++
				continue
			}
			if edgeCoordsX[i]-1 != edgeCoordsX[i-1] {
				bottomSides++
			}
		}
	}

	return topSides + bottomSides + leftSides + rightSides
}

func findRegionProperties(grid [][]rune, start coord, region region) region {

	// add coord to region
	region.coords = append(region.coords, start)

	upCoord := coord{start.x, start.y - 1}
	downCoord := coord{start.x, start.y + 1}
	leftCoord := coord{start.x - 1, start.y}
	rightCoord := coord{start.x + 1, start.y}

	for _, nextCoord := range []coord{upCoord, downCoord, leftCoord, rightCoord} {

		// check if the nextCoord is in the grid
		if nextCoord.y < 0 || nextCoord.y >= len(grid) || nextCoord.x < 0 || nextCoord.x >= len(grid[0]) {
			region.perimeter++
			continue
		}

		// check if the nextCoord is valid
		if grid[nextCoord.y][nextCoord.x] == '.' || grid[nextCoord.y][nextCoord.x] != rune(region.name[0]) {
			region.perimeter++
			continue
		}

		// check if the nextCoord is already in the region
		alreadyInRegion := false
		for _, coord := range region.coords {
			if coord == nextCoord {
				alreadyInRegion = true
			}
		}

		if alreadyInRegion {
			continue
		}

		// add the nextCoord to the region and recurse
		region = findRegionProperties(grid, nextCoord, region)
	}

	// can move this out to optimize
	region.area = len(region.coords)
	region.sides = calculateSides(region, grid)

	return region
}

func findRegions(grid [][]rune) []region {

	// iterate over the grid removing regions as they are located
	regions := []region{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {

			// already regioned
			if grid[y][x] == '.' {
				continue
			}

			// find the region
			region := findRegionProperties(grid, coord{x, y}, region{name: string(grid[y][x])})
			regions = append(regions, region)

			// remove the region from the grid
			for _, coord := range region.coords {
				grid[coord.y][coord.x] = '.'
			}
		}
	}
	return regions
}

func Part1(filename string, debug bool) int {
	grid := parse(filename)

	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				fmt.Printf("%s", string(grid[y][x]))
			}
			fmt.Println()
		}
	}

	regions := findRegions(grid)

	ans := 0
	for _, region := range regions {
		ans += region.area * region.perimeter
	}

	return ans
}

func Part2(filename string, debug bool) int {
	grid := parse(filename)

	if debug {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				fmt.Printf("%s", string(grid[y][x]))
			}
			fmt.Println()
		}
	}

	regions := findRegions(grid)

	ans := 0
	for _, region := range regions {
		if debug {
			fmt.Printf("Region %v: %v\n", region.name, region)
		}
		ans += region.area * region.sides
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day12")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
