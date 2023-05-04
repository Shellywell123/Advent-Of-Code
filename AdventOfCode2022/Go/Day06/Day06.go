package main

import (
	"fmt"
	"os"
)

func Parse(filename string) string {

	// read file into string
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func Solve(filename string, n int) int {

	data := Parse(filename)

	for i := 0; i < len(data); i++ {

		distinct := map[string]int {}
		chars := string(data[i:i+n])
		for i := 0; i < len(chars); i++ {
			distinct[string(chars[i])] = 0
		}

		// checck if all characters are distinct
		if len(distinct) == n  {
			return i + n
		}
	}
	return 0
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day06")
	fmt.Printf("Tests : Answer to Part 1 = %i\n", Solve(testfile, 4))
	fmt.Printf("Inputs: Answer to Part 1 = %i\n", Solve(inputfile,4))
	fmt.Printf("Tests : Answer to Part 2 = %i\n", Solve(testfile,14))
	fmt.Printf("Inputs: Answer to Part 2 = %i\n", Solve(inputfile,14))
}
