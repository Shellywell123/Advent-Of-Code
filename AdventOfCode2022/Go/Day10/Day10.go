package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) []string {

	file, _ := os.ReadFile(filename)

	data := strings.Split(string(file), "\n")

	return data
}

func Part1(filename string) int {

	instructions := Parse(filename)

	sums := 0
	cycle := 0
	x := 1

	for _, instruction := range instructions {

		cycle++

		if (cycle-20)%40 == 0 {
			sums += x * cycle
		}

		if instruction[:4] != "noop" {

			cycle++

			if (cycle-20)%40 == 0 {
				sums += x * cycle
			}

			xp, _ := strconv.Atoi((strings.TrimSpace(instruction[5:])))
			x += xp
		}
	}
	return sums
}

func Part2(filename string, printing bool) int {

	// create blank canvas
	pixels := []string{}
	for i := 0; i <= 40*6; i++ {
		pixels = append(pixels, ".")
	}

	instructions := Parse(filename)

	cycle := 1
	x := 1

	for _, instruction := range instructions {

		if printing {
			fmt.Print("Start cycle   ", cycle, ": begin executing ", instruction, "\n")
		}

		if x+(40*int(math.Floor((float64(cycle)/40)))) <= cycle && cycle <= x+2+(40*int(math.Floor((float64(cycle)/40)))) {
			pixels[cycle-1] = "#"
		}

		if printing {
			fmt.Print("During cycle  ", cycle, ": CRT draws pixel in position ", cycle-1, "\n")
			fmt.Print("Current CRT row: ", pixels[:cycle], "\n\n")
		}

		// AddX
		if instruction[:4] != "noop" {

			cycle++

			if x+(40*int(math.Floor((float64(cycle)/40)))) <= cycle && cycle <= x+2+(40*int(math.Floor((float64(cycle)/40)))) {
				pixels[cycle-1] = "#"
			}

			if printing {
				fmt.Print("During cycle  ", cycle, ": CRT draws pixel in position ", cycle-1, "\n")
				fmt.Print("Current CRT row: ", pixels[:cycle], "\n")
			}

			xp, _ := strconv.Atoi((strings.TrimSpace(instruction[5:])))
			x += xp
		}
		if printing {
			fmt.Print("End of cycle  ", cycle, ": finish executing ", instruction[:len(instruction)-1], " (Register X is now ", x, ")\n")
		}
		cycle++
	}

	// Print the CRT
	for row := 0; row < 6; row++ {
		fmt.Print(pixels[0+(40*row):40+(40*row)], "\n")
	}

	return 0
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day10")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile, false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile, false))

}
