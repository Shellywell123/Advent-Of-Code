package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func Catchup(headPosition []int, tailPosition []int) []int {
	newTailPosition := tailPosition

	// head = centre of clock
	// tail = number on clock

	// 1
	if headPosition[0]+1 == tailPosition[0] && headPosition[1]-2 == tailPosition[1] {
		newTailPosition[1]--
		newTailPosition[1]++
	}

	// 2
	if headPosition[0]+2 == tailPosition[0] && headPosition[1]-1 == tailPosition[1] {
		newTailPosition[0]--
		newTailPosition[1]++
	}

	// 3
	if headPosition[0]+2 == tailPosition[0] && headPosition[1] == tailPosition[1] {
		newTailPosition[0]--
	}

	// 4
	if headPosition[0]+2 == tailPosition[0] && headPosition[1]+1 == tailPosition[1] {
		newTailPosition[0]--
		newTailPosition[1]--
	}

	// 5
	if headPosition[0]+1 == tailPosition[0] && headPosition[1]+2 == tailPosition[1] {
		newTailPosition[0]--
		newTailPosition[1]--
	}

	// 6
	if headPosition[0] == tailPosition[0] && headPosition[1]+2 == tailPosition[1] {
		newTailPosition[0]--
	}

	// 7
	if headPosition[0]-1 == tailPosition[0] && headPosition[1]+2 == tailPosition[1] {
		newTailPosition[0]++
		newTailPosition[1]--
	}

	// 8
	if headPosition[0]-2 == tailPosition[0] && headPosition[1]+1 == tailPosition[1] {
		newTailPosition[0]++
		newTailPosition[1]--
	}

	// 9
	if headPosition[0]-2 == tailPosition[0] && headPosition[1] == tailPosition[1] {
		newTailPosition[0]++
	}

	// 10
	if headPosition[0]-2 == tailPosition[0] && headPosition[1]-1 == tailPosition[1] {
		newTailPosition[0]++
		newTailPosition[1]++
	}

	// 11
	if headPosition[0]-1 == tailPosition[0] && headPosition[1]-2 == tailPosition[1] {
		newTailPosition[0]++
		newTailPosition[1]++
	}

	// 12
	if headPosition[0] == tailPosition[0] && headPosition[1] -2== tailPosition[1] {
		newTailPosition[1]++
	}

	return newTailPosition
}

func Part1(filename string) int {

	data := Parse(filename)

	tailPositions := map[string]string{fmt.Sprintf("%v-%v", 0, 99): ""}

	currentHeadPosition := []int{0, 99}
	currentTailPosition := []int{0, 99}

	// draw path
	for _, move := range data {

		direction, magnitudeString := move["direction"], move["magnitude"]
		magnitude, _ := strconv.Atoi(magnitudeString)

		// fmt.Println(direction, magnitude)

		switch direction {
		case "U":
			for i := 0; i < magnitude; i++ {
				// fmt.Print(currentHeadPosition,currentTailPosition,"\n")
				currentHeadPosition[1]--
				currentTailPosition = Catchup(currentHeadPosition, currentTailPosition)
				tailPositions[fmt.Sprintf("%v-%v", currentTailPosition[0], currentTailPosition[1])] = ""
			}
		case "D":
			for i := 0; i < magnitude; i++ {
				// fmt.Print(currentHeadPosition,currentTailPosition,"\n")
				currentHeadPosition[1]++
				currentTailPosition = Catchup(currentHeadPosition, currentTailPosition)
				tailPositions[fmt.Sprintf("%v-%v", currentTailPosition[0], currentTailPosition[1])] = ""
			}
		case "L":
			for i := 0; i < magnitude; i++ {
				// fmt.Print(currentHeadPosition,currentTailPosition,"\n")
				currentHeadPosition[0]--
				currentTailPosition = Catchup(currentHeadPosition, currentTailPosition)
				tailPositions[fmt.Sprintf("%v-%v", currentTailPosition[0], currentTailPosition[1])] = ""
			}
		case "R":
			for i := 0; i < magnitude; i++ {
				// fmt.Print(currentHeadPosition,currentTailPosition,"\n")
				currentHeadPosition[0]++
				currentTailPosition = Catchup(currentHeadPosition, currentTailPosition)
				tailPositions[fmt.Sprintf("%v-%v", currentTailPosition[0], currentTailPosition[1])] = ""
			}
		}
	}

	// for k, v := range tailPositions {
	// 	fmt.Println(k, v,)
	// }

	return len(tailPositions)
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day09")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	// fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile))
	// fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))
}
