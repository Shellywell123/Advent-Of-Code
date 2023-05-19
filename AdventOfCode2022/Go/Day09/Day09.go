package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type knot struct {
	head coordinate
	tail coordinate
}

func Parse(filename string) []map[string]string {

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	data := []map[string]string{}

	for fscanner.Scan() {
		line := fscanner.Text()
		direction, magnitude := line[:strings.Index(line, " ")], line[strings.Index(line, " ")+1:]

		lineData := map[string]string{"direction": direction, "magnitude": magnitude}
		data = append(data, lineData)
	}

	return data
}

func MoveHead(headPosition coordinate, direction string) coordinate {

	switch direction {
	case "U":
		headPosition.y--
	case "D":
		headPosition.y++
	case "L":
		headPosition.x--
	case "R":
		headPosition.x++
	}

	return headPosition
}

func Catchup(headPosition coordinate, tailPosition coordinate) coordinate {

	newTailPosition := tailPosition

	dx := float64(tailPosition.x - headPosition.x)
	dy := float64(tailPosition.y - headPosition.y)

	theta := math.Atan2(dx, dy)

	if math.Pow(dx, 2)+math.Pow(dy, 2) > 2 {
		newTailPosition.x -= int(math.Round(1.4 * math.Sin(theta)))
		newTailPosition.y -= int(math.Round(1.4 * math.Cos(theta)))
	}

	return newTailPosition
}

func PrintKnots(knots []knot) {

	minX := 0
	maxX := 0

	minY := 0
	maxY := 0

	for _, knot := range knots {
		if knot.head.x < minX {
			minX = knot.head.x
		}
		if knot.head.x > maxX {
			maxX = knot.head.x
		}
		if knot.head.y < minY {
			minY = knot.head.y
		}
		if knot.head.y > maxY {
			maxY = knot.head.y
		}
	}

	canvas := [][]string{}
	for y := minY; y <= maxY+1; y++ {
		row := []string{}
		for x := minX; x <= maxX+1; x++ {
			row = append(row, ".")
		}
		canvas = append(canvas, row)
	}

	for k := 0; k < len(knots); k++ {
		canvas[knots[k].head.y-minY][knots[k].head.x-minX] = strconv.Itoa(k)
	}
	canvas[-minY][-minX] = "s"

	for row := range canvas {
		fmt.Println(canvas[row])
	}
}

func PrintTailPositions(tailPositions map[string]interface{}) {

	minX := 0
	maxX := 0

	minY := 0
	maxY := 0

	for key := range tailPositions {
		x, _ := strconv.Atoi(strings.Split(key, ",")[0])
		y, _ := strconv.Atoi(strings.Split(key, ",")[1])

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	canvas := [][]string{}
	for y := minY; y <= maxY+1; y++ {
		row := []string{}
		for x := minX; x <= maxX+1; x++ {
			row = append(row, ".")
		}
		canvas = append(canvas, row)
	}

	for key := range tailPositions {
		x, _ := strconv.Atoi(strings.Split(key, ",")[0])
		y, _ := strconv.Atoi(strings.Split(key, ",")[1])

		canvas[y-minY][x-minX] = "#"
	}

	for row := range canvas {
		fmt.Println(canvas[row])
	}
}

func Part1(filename string) int {

	data := Parse(filename)

	tailPositions := map[string]interface{}{}

	// initialize head position as 0,0
	currentHeadPosition := coordinate{0, 0}
	currentTailPosition := currentHeadPosition

	// draw path
	for _, move := range data {

		direction, magnitudeString := move["direction"], move["magnitude"]
		magnitude, _ := strconv.Atoi(magnitudeString)

		for i := 0; i < magnitude; i++ {

			// move head along
			currentHeadPosition = MoveHead(currentHeadPosition, direction)

			// move tail along
			currentTailPosition = Catchup(currentHeadPosition, currentTailPosition)

			// add tail position to map
			tailPositions[fmt.Sprintf("%v,%v", currentTailPosition.x, currentTailPosition.y)] = ""
		}
	}

	return len(tailPositions)
}

func Part2(filename string) int {

	data := Parse(filename)

	numOfKnots := 9

	tailPositions := map[string]interface{}{}

	// initialize knot positions as 0,0
	knots := []knot{}
	for i := 0; i < numOfKnots; i++ {
		currentHeadPosition := coordinate{0, 0}
		currentTailPosition := currentHeadPosition
		knots = append(knots, knot{head: currentHeadPosition, tail: currentTailPosition})
	}

	// draw path
	for _, move := range data {

		direction, magnitudeString := move["direction"], move["magnitude"]
		// catchupDirection := direction
		magnitude, _ := strconv.Atoi(magnitudeString)

		for i := 0; i < magnitude; i++ {

			// move head along
			knots[0].head = MoveHead(knots[0].head, direction)

			// move tails along
			for k := 0; k < numOfKnots; k++ {
				if k != 0 {
					knots[k].head = knots[k-1].tail
				}
				knots[k].tail = Catchup(knots[k].head, knots[k].tail)
			}

			// add end tail position to map
			tailPositions[fmt.Sprintf("%v,%v", knots[numOfKnots-1].tail.x, knots[numOfKnots-1].tail.y)] = ""

			// PrintKnots(knots)
		}
	}

	// PrintTailPositions(tailPositions)
	return len(tailPositions)
}

func main() {

	testfile := "tests.txt"
	testfile2 := "tests2.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day09")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile2))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))

}
