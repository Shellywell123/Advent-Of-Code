package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) []map[string]string {

	// read file into string
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")

	data := []map[string]string{}
	for _, line := range lines {

		dir := strings.Split(line, " ")[0]
		mag := strings.Split(line, " ")[1]

		entry := map[string]string{
			"direction": dir,
			"magnitude": mag,
		}

		data = append(data, entry)
	}

	return data
}

func Part1(filename string) int {

	data := Parse(filename)

	x := 0
	y := 0

	for _, entry := range data {

		mag, err := strconv.Atoi(entry["magnitude"])
		if err != nil {
			panic(err)
		}

		switch entry["direction"] {

		case "up":
			y -= mag

		case "down":
			y += mag

		case "forward":
			x += mag
		}
	}

	return (x * y)
}

func Part2(filename string) int {

	data := Parse(filename)

	x := 0
	y := 0
	aim := 0

	for _, entry := range data {

		mag, err := strconv.Atoi(entry["magnitude"])
		if err != nil {
			panic(err)
		}

		switch entry["direction"] {

		case "up":
			aim -= mag

		case "down":
			aim += mag

		case "forward":
			x += mag
			y += (aim * mag)

		}
	}

	return (x * y)
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day02")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))
}
