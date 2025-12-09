package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	coords := [][]int{}

	for fscanner.Scan() {

		line := fscanner.Text()
		nums := []int{}
		for _, char := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(char)
			nums = append(nums, num)
		}
		coords = append(coords, nums)
	}

	return coords
}

func Part1(filename string, debug bool) int {
	coords := Parse(filename)

	ans := 0
	for _, coord1 := range coords {
		for _, coord2 := range coords {

			if coord1[0] == coord2[0] && coord1[1] == coord2[1] {
				continue
			}

			x := math.Abs(float64(coord1[0]-coord2[0])) + 1
			y := math.Abs(float64(coord1[1]-coord2[1])) + 1

			area := int(x * y)

			if area > ans {
				ans = area
			}

		}
	}
	return ans
}

func Part2(filename string, debug bool) int {
	redCoords := Parse(filename)

	greenCoords := [][]int{}

	// Connect each consecutive pair of red coordinates
	for n := 0; n < len(redCoords)-1; n++ {

		// same x so fill y
		if redCoords[n][0] == redCoords[n+1][0] {
			startY := redCoords[n][1]
			endY := redCoords[n+1][1]
			if startY > endY {
				startY, endY = endY, startY
			}
			for y := startY; y <= endY; y++ {
				greenCoords = append(greenCoords, []int{redCoords[n][0], y})
			}
		}

		// same y so fill x
		if redCoords[n][1] == redCoords[n+1][1] {
			startX := redCoords[n][0]
			endX := redCoords[n+1][0]
			if startX > endX {
				startX, endX = endX, startX
			}
			for x := startX; x <= endX; x++ {
				greenCoords = append(greenCoords, []int{x, redCoords[n][1]})
			}
		}
	}

	// complete loop - connect last red coord back to first
	lastIdx := len(redCoords) - 1
	if redCoords[0][0] == redCoords[lastIdx][0] {
		// same x, fill y
		startY := redCoords[0][1]
		endY := redCoords[lastIdx][1]
		if startY > endY {
			startY, endY = endY, startY
		}
		for y := startY; y <= endY; y++ {
			greenCoords = append(greenCoords, []int{redCoords[0][0], y})
		}
	} else if redCoords[0][1] == redCoords[lastIdx][1] {
		// same y, fill x
		startX := redCoords[0][0]
		endX := redCoords[lastIdx][0]
		if startX > endX {
			startX, endX = endX, startX
		}
		for x := startX; x <= endX; x++ {
			greenCoords = append(greenCoords, []int{x, redCoords[0][1]})
		}
	} else {
		fmt.Printf("WARNING: Last coord %v and first coord %v don't share X or Y!\n", redCoords[lastIdx], redCoords[0])
	}

	allCoords := [][]int{}
	allCoords = append(allCoords, redCoords...)
	allCoords = append(allCoords, greenCoords...)

	if debug {
		fmt.Printf("Total edge coords: %d (red: %d, green: %d)\n", len(allCoords), len(redCoords), len(greenCoords))
	}
	// Create edge map for O(1) lookups using int64 keys - NO flood fill
	edgeMap := make(map[int64]bool, len(allCoords))
	for _, coord := range allCoords {
		key := int64(coord[0])<<32 | int64(coord[1])
		edgeMap[key] = true
	}

	// Cache for ray casting results
	insideCache := make(map[int64]bool)

	// Ray casting to check if point is inside polygon
	isInside := func(x, y int) bool {
		key := int64(x)<<32 | int64(y)
		if edgeMap[key] {
			return true // On edge counts as inside
		}

		if cached, ok := insideCache[key]; ok {
			return cached
		}

		intersections := 0
		for i := 0; i < len(redCoords); i++ {
			j := (i + 1) % len(redCoords)
			x1, y1 := redCoords[i][0], redCoords[i][1]
			x2, y2 := redCoords[j][0], redCoords[j][1]
			if ((y1 > y) != (y2 > y)) && (x < (x2-x1)*(y-y1)/(y2-y1)+x1) {
				intersections++
			}
		}
		result := intersections%2 == 1
		insideCache[key] = result
		return result
	}

	ans := 0
	checked := 0
	skipped := 0
	total := len(redCoords) * len(redCoords)

	for i, coord1 := range redCoords {
		if i%50 == 0 && debug {
			fmt.Printf("Progress: %d/%d rectangles checked, current best: %d\n", i*len(redCoords), total, ans)
		}

		for _, coord2 := range redCoords {

			if coord1[0] == coord2[0] && coord1[1] == coord2[1] {
				continue
			}

			// Calculate area first - skip if it can't beat current best
			w := int(math.Abs(float64(coord1[0]-coord2[0])) + 1)
			h := int(math.Abs(float64(coord1[1]-coord2[1])) + 1)
			area := w * h

			if area <= ans {
				skipped++
				continue
			}

			checked++

			// if all sides are in the coords list then we have a valid rectangle
			x1 := 0
			y1 := 0
			x2 := 0
			y2 := 0

			if coord1[0] < coord2[0] {
				x1 = coord1[0]
				x2 = coord2[0]
			} else {
				x1 = coord2[0]
				x2 = coord1[0]
			}

			if coord1[1] < coord2[1] {
				y1 = coord1[1]
				y2 = coord2[1]
			} else {
				y1 = coord2[1]
				y2 = coord1[1]
			}

			// Check if rectangle is inside - use sampling for large rectangles
			allFound := true

			if w*h <= 5000 {
				// Small rectangle - check every pixel
				for x := x1; x <= x2 && allFound; x++ {
					for y := y1; y <= y2; y++ {
						if !isInside(x, y) {
							allFound = false
							break
						}
					}
				}
			} else {
				// Large rectangle - check corners + sparse sample
				// Check the 4 corners
				if !isInside(x1, y1) || !isInside(x1, y2) || !isInside(x2, y1) || !isInside(x2, y2) {
					allFound = false
				}

				// Sample edges sparingly (every 10th pixel)
				if allFound {
					step := 10
					for x := x1; x <= x2 && allFound; x += step {
						if !isInside(x, y1) || !isInside(x, y2) {
							allFound = false
							break
						}
					}
					for y := y1; y <= y2 && allFound; y += step {
						if !isInside(x1, y) || !isInside(x2, y) {
							allFound = false
							break
						}
					}
				}

				// Sample a very sparse grid in the interior
				if allFound {
					stepX := w / 10
					stepY := h / 10
					if stepX < 10 {
						stepX = 10
					}
					if stepY < 10 {
						stepY = 10
					}
					for x := x1 + stepX; x < x2 && allFound; x += stepX {
						for y := y1 + stepY; y < y2; y += stepY {
							if !isInside(x, y) {
								allFound = false
								break
							}
						}
					}
				}
			}

			if !allFound {
				continue
			}

			if debug {
				fmt.Println("Found rectangle between:", coord1, coord2)
			}

			if area > ans {
				if debug {
					fmt.Println(coord1, coord2, area)
				}
				ans = area
			}
		}
	}

	if debug {
		fmt.Printf("Checked %d rectangles, skipped %d\n", checked, skipped)
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day09")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}

// used some ai to help optimize:

// Ray-casting for polygon containment instead of flood fill - much faster for point-in-polygon tests
// Sparse sampling for large rectangles (checking every 10th pixel instead of all pixels) - dramatically reduces computation time
// Caching ray-cast results - avoids redundant calculations for the same coordinates
// Early termination - skipping rectangles smaller than the current best area
// This approach handles the massive 96,000 × 96,000 coordinate space efficiently by avoiding the need to fill all 9.3 billion cells, instead only checking the specific rectangles formed by pairs of red coordinates.
