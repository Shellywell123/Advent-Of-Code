package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func Parse(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	histories := [][]int{}
	for fscanner.Scan() {
		line := fscanner.Text()

		history := []int{}

		for _, number := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(string(number))
			history = append(history, n)
		}
		histories = append(histories, history)
	}

	return histories
}

func AllEqualOne(sequence []int) bool {
	for _, n := range sequence {
		if math.Sqrt(float64(n*n)) >= 1.0 {
			return false
		}
	}
	return true
}

func Part1(filename string) int {

	histories := Parse(filename)
	total := 0

	for _, history := range histories {

		// interpret history
		sequences := [][]int{}
		currentSequence := history
		for !AllEqualOne(currentSequence) {

			sequences = append(sequences, currentSequence)
			nextSequence := []int{}
			for i :=0; i < len(currentSequence) - 1; i++ {
				nextSequence = append(nextSequence, currentSequence[i+1] - currentSequence[i])
			}
			
			currentSequence = nextSequence
		}

		// predict next value
		increment := 0
		for i, _ := range sequences {
			increment += sequences[i][len(sequences[i])-1]
		}
		total += increment
	}
	return total
}

func Part2(filename string) int {

	histories := Parse(filename)
	total := 0

	for _, history := range histories {

		// interpret history
		sequences := [][]int{}
		
		// quick hack is to reverse the arrays 
		historyReversed := []int{}

		for i := len(history) - 1 ; i >=0; i-- {
			historyReversed = append(historyReversed, history[i])
		}

		currentSequence := historyReversed
		for !AllEqualOne(currentSequence) {

			sequences = append(sequences, currentSequence)
			nextSequence := []int{}
			for i :=0; i < len(currentSequence) - 1; i++ {
				nextSequence = append(nextSequence, currentSequence[i+1] - currentSequence[i])
			}
			
			currentSequence = nextSequence
		}
		fmt.Println(sequences)

		// predict next value
		increment := 0
		for i, _ := range sequences {
			increment += sequences[i][len(sequences[i])-1]
		}
		total += increment
	}
	return total
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day09")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
