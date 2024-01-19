package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(filename string) []string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	data := []string{}
	for fscanner.Scan() {
		line := fscanner.Text()
		data = append(data, line)
	}
	return data
}

func Rotate90(data []string) []string {

	newData := []string{}
	for x := len(data[0]) - 1; x >= 0; x-- {

		// make column slice
		colSlice := ""
		for y := len(data) - 1; y >= 0; y-- {
			colSlice += string(data[y][x])
		}
		newData = append(newData, colSlice)
	}

	return newData
}

func TiltRowLeft(slice []string) ([]string, int) {
	changes := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] == "O" && slice[i-1] == "." {
			slice[i], slice[i-1] = slice[i-1], slice[i]
			changes++
		}
	}
	return slice, changes
}

func TiltDataNorth(data []string) []string {

	newData := []string{}

	for y := 0; y < len(data); y++ {
		newData = append(newData, "")
	}

	for x := len(data[0]) - 1; x >= 0; x-- {

		// make column slice
		colSlice := []string{}
		for y := 0; y < len(data); y++ {
			colSlice = append(colSlice, string(data[y][x]))
		}

		// tilt column slice
		changes := 1
		for changes != 0 {
			colSlice, changes = TiltRowLeft(colSlice)
		}

		for z := 0; z < len(colSlice); z++ {
			newData[z] += colSlice[z]
		}
	}

	return newData
}

func CalculateLoad(data []string) int {
	load := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if data[y][x] == 'O' {
				load += len(data) - y
			}
		}
	}
	return load
}

func Part1(filename string) int {
	data := Parse(filename)
	return CalculateLoad(TiltDataNorth(data))
}

func Part2(filename string) int {
	data := Parse(filename)

	cache := []string{} // DP
	cycleStart := 0
	cycleEnd := 0
	gridString := ""
	cycles := 1000000000
	c := 0
	for cycleStart == 0 {
		for i := 0; i < 4; i++ {
			data = Rotate90(TiltDataNorth(data))
		}

		gridString = "" // has to be redfined as an empty string each time
		for _, line := range data {
			gridString += "-" + line
		}

		for j, _ := range cache {
			if gridString == cache[j] {
				cycleStart = j + 1
				cycleEnd = c + 1
			}
		}

		cache = append(cache, gridString)
		c++
	}

	d := cache[cycleStart+((cycles-cycleStart)%(cycleEnd-cycleStart))-1] // -1 ?
	finalData := strings.Split(d, "-")[1:]

	return CalculateLoad(finalData)
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day14")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
