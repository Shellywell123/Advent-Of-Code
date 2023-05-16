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

func Catchup(headPosition coordinate, tailPosition coordinate, lastDirection string) coordinate {
	newTailPosition := tailPosition

	if math.Abs(float64(headPosition.x)-float64(tailPosition.x)) > 1 || math.Abs(float64(headPosition.y)-float64(tailPosition.y)) > 1 {

		switch lastDirection {
		case "U":
			newTailPosition.x = headPosition.x
			newTailPosition.y = headPosition.y + 1
		case "D":
			newTailPosition.x = headPosition.x
			newTailPosition.y = headPosition.y - 1
		case "L":
			newTailPosition.x = headPosition.x + 1
			newTailPosition.y = headPosition.y
		case "R":
			newTailPosition.x = headPosition.x - 1
			newTailPosition.y = headPosition.y
		}
	}

	return newTailPosition
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
			switch direction {
			case "U":
				currentHeadPosition.y--
			case "D":
				currentHeadPosition.y++
			case "L":
				currentHeadPosition.x--
			case "R":
				currentHeadPosition.x++
			}

			// move tail along
			currentTailPosition = Catchup(currentHeadPosition, currentTailPosition, direction)

			// add tail position to map
			tailPositions[fmt.Sprintf("%v,%v", currentTailPosition.x, currentTailPosition.y)] = ""
		}
	}

	return len(tailPositions)
}

func Part2(filename string) int {

	data := Parse(filename)

	tailPositions := map[string]interface{}{}

	// initialize knot positions as 0,0
	knots := []knot{}
	for i := 0; i < 10; i++ {
		currentHeadPosition := coordinate{0, 0}
		currentTailPosition := currentHeadPosition
		knots = append(knots, knot{head: currentHeadPosition, tail: currentTailPosition})
	}

	// draw path
	for _, move := range data {

		direction, magnitudeString := move["direction"], move["magnitude"]
		magnitude, _ := strconv.Atoi(magnitudeString)

		fmt.Print(direction, " ", magnitude, "\n")

		for i := 0; i < magnitude; i++ {

			// move head along
			switch direction {
			case "U":
				knots[0].head.y--
			case "D":
				knots[0].head.y++
			case "L":
				knots[0].head.x--
			case "R":
				knots[0].head.x++
			}

			// move tails along
			for k := 0; k < len(knots); k++ {
				if k != 0 {
					knots[k].tail = Catchup(knots[k-1].tail, knots[k].tail, direction)
				} else {
					knots[k].tail = Catchup(knots[k].head, knots[k].tail, direction)
				}
			}

			// add end tail position to map
			tailPositions[fmt.Sprintf("%v,%v", knots[len(knots)-1].tail.x, knots[len(knots)-1].tail.y)] = ""
		}
	}

	fmt.Print(tailPositions)
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
	// fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))
}
