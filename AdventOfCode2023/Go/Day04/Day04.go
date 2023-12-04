package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isInt(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

type ScratchCard struct {
	number         int
	duplicates     int
	winningNumbers []int
	scratchNumbers []int
	matches        int
}

func Parse(filename string) map[int]ScratchCard {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	scratchCards := map[int]ScratchCard{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		newScratchCard := ScratchCard{}
		newScratchCard.number = y + 1
		newScratchCard.duplicates = 1
		x := strings.Index(line, ": ")
		x += 3

		newScratchCard.winningNumbers = []int{}
		newScratchCard.scratchNumbers = []int{}

		for _, num := range strings.Split(line[strings.Index(line, ": ") + 2:strings.Index(line, " | ") + 2], " ") {
			val, _ := strconv.Atoi(num)
			if val > 0 {
				newScratchCard.winningNumbers = append(newScratchCard.winningNumbers, val)
			}
		}

		for _, num := range strings.Split(line[strings.Index(line, " | ") + 2:], " ") {
			val, _ := strconv.Atoi(num)
			if val > 0 {
				newScratchCard.scratchNumbers = append(newScratchCard.scratchNumbers, val)
			}
		}

		scratchCards[newScratchCard.number] = newScratchCard
		y++
	}

	return scratchCards
}

func Part1(filename string) int {
	scratchCards := Parse(filename)
	totalValue := 0
	for _, scratchCard := range scratchCards {
		matches := 0
		for _, result := range scratchCard.scratchNumbers {
			for _, winner := range scratchCard.winningNumbers {
				if result == winner {
					matches++
					break
				}
			}
		}
		totalValue += int(math.Pow(2, float64(matches-1)))
	}

	return totalValue
}

func Part2(filename string) int {
	scratchCards := Parse(filename)

	for _, scratchCard := range scratchCards {
		matches := 0
		for _, result := range scratchCard.scratchNumbers {
			for _, winner := range scratchCard.winningNumbers {
				if result == winner {
					matches++
					break
				}
			}
		}
		scratchCard.matches = matches
		scratchCards[scratchCard.number] = scratchCard
	}

	for s := 1; s <= len(scratchCards); s++ {
		for i := 1; i <= scratchCards[s].matches; i++ {
			temp := scratchCards[s+i]
			temp.duplicates += (scratchCards[s].duplicates)
			scratchCards[s+i] = temp
		}
	}

	numOfScratchCards := 0
	for _, scratchCard := range scratchCards {
		numOfScratchCards += scratchCard.duplicates
	}

	return numOfScratchCards
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day04")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
