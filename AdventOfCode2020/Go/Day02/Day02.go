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

		min := strings.Split(line, "-")[0]
		max := strings.Split(strings.Split(line, " ")[0], "-")[1]
		char := strings.Split(strings.Split(line, " ")[1], ":")[0]
		password := strings.Split(line, ": ")[1]

		entry := map[string]string{
			"char":           char,
			"min-occurences": min,
			"max-occurences": max,
			"password":       password,
		}

		data = append(data, entry)
	}

	return data
}

func Part1(filename string) int {

	data := Parse(filename)

	numOFValidPasswords := 0
	for _, entry := range data {

		min, err := strconv.Atoi(entry["min-occurences"])
		if err != nil {
			panic(err)
		}

		max, err := strconv.Atoi(entry["max-occurences"])
		if err != nil {
			panic(err)
		}

		charCount := 0
		for i := 0; i < (len(entry["password"])); i++ {
			if string(entry["password"][i]) == entry["char"] {
				charCount++
			}
		}

		if charCount >= min && charCount <= max {
			numOFValidPasswords++
		}
	}

	return numOFValidPasswords
}

func Part2(filename string) int {

	data := Parse(filename)

	numOFValidPasswords := 0
	for _, entry := range data {

		firstIndex, err := strconv.Atoi(entry["min-occurences"])
		if err != nil {
			panic(err)
		}

		secondIndex, err := strconv.Atoi(entry["max-occurences"])
		if err != nil {
			panic(err)
		}

		if (string(entry["password"][firstIndex-1]) == entry["char"] &&
			string(entry["password"][secondIndex-1]) != entry["char"]) ||
			(string(entry["password"][firstIndex-1]) != entry["char"] &&
				string(entry["password"][secondIndex-1]) == entry["char"]) {
			numOFValidPasswords++
		}

	}
	return numOFValidPasswords
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
