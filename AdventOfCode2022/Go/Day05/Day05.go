package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) ([][]string, [][]int) {

	// read file into string
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	data := strings.Split(string(b), "\n\n")

	// reformat stacks data
	stacksData := strings.Split(data[0], "\n")

	stacks := [][]string{}
	stack := []string{}

	for x := 1; x < len(string(stacksData[0])); x = x + 4 {
		for y := len(stacksData) - 2; y > -1; y-- {

			// create new stack
			if y == len(stacksData)-2 {
				stack = []string{}
			}

			// add crate to stack
			if string(stacksData[y][x]) != " " {
				stack = append(stack, string(stacksData[y][x]))
			}

			// add stack to stacks
			if y == 0 {
				stacks = append(stacks, stack)
			}
		}
	}

	//reformat moves data
	movesData := strings.Split(data[1], "\n")

	moves := [][]int{}
	move := []int{}

	for _, line := range movesData {

		quantity, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			panic(err)
		}

		from, err := strconv.Atoi(strings.Split(line, " ")[3])
		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(strings.Split(line, " ")[5])
		if err != nil {
			panic(err)
		}

		move = []int{quantity, from, to}
		moves = append(moves, move)
	}

	return stacks, moves
}

func Part1(filename string) string {

	stacks, moves := Parse(filename)

	for _, move := range moves {
		quantity := move[0]
		from := move[1]
		to := move[2]

		// move crates
		stacksToMove := stacks[from-1][len(stacks[from-1])-quantity:]
		stacks[from-1] = stacks[from-1][:len(stacks[from-1])-quantity]

		for i := len(stacksToMove) - 1; i > -1; i-- {
			stacks[to-1] = append(stacks[to-1], stacksToMove[i])
		}
	}

	answer := ""
	for _, stack := range stacks {
		answer += stack[len(stack)-1]
	}
	return answer
}

func Part2(filename string) string {

	stacks, moves := Parse(filename)

	for _, move := range moves {
		quantity := move[0]
		from := move[1]
		to := move[2]

		// move crates
		stacksToMove := stacks[from-1][len(stacks[from-1])-quantity:]
		stacks[from-1] = stacks[from-1][:len(stacks[from-1])-quantity]
		stacks[to-1] = append(stacks[to-1], stacksToMove...)

	}

	answer := ""
	for _, stack := range stacks {
		answer += stack[len(stack)-1]
	}
	return answer
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day05")
	fmt.Printf("Tests : Answer to Part 1 = %s\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %s\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %s\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %s\n", Part2(inputfile))
}
