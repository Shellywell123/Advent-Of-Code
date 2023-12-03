package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Outcome struct {
	red int
	green int
	blue int
}

type Game struct {
	id int
	outcomes []Outcome
}

func Parse(filename string) []Game {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	games := []Game{}

	for fscanner.Scan() {
		line := fscanner.Text()

		game := Game{}
		outcomes := []Outcome{}
		game.id, _ = strconv.Atoi(line[5:strings.Index(line, ":")])

		raw_outcomes := strings.Split(line[strings.Index(line, ":")+1:], ";")

		for _, raw_outcome  := range raw_outcomes {
			outcome := Outcome{red: 0, green: 0, blue :0}

			cubes := strings.Split(raw_outcome, ",")
			for _, cube_color := range cubes {
				if strings.Contains(cube_color, "red") {
					outcome.red, _ = strconv.Atoi(strings.Split(cube_color, " ")[1])
				}
				if strings.Contains(cube_color, "blue") {
					outcome.blue, _ = strconv.Atoi(strings.Split(cube_color, " ")[1])
				}
				if strings.Contains(cube_color, "green") {
					outcome.green, _ = strconv.Atoi(strings.Split(cube_color, " ")[1])
				}
			}

			outcomes = append(outcomes, outcome)

		}
		game.outcomes = outcomes
		games = append(games, game)
	}

	return games
}

func Part1(filename string) int {
	games := Parse(filename)

	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	total := 0
	for _, game := range games {
		possible := true
		for _, outcome := range game.outcomes {
			if outcome.red > redLimit || outcome.green > greenLimit || outcome.blue > blueLimit || outcome.red + outcome.green + outcome.blue > redLimit + greenLimit + blueLimit {
				possible = false
			}
		}
		if possible {
			total += game.id
		}
	}
	return total
}

func Part2(filename string) int {
	games := Parse(filename)

	totalPower := 0

	for _, game := range games {

		redMinimum := 0
		greenMinimum := 0
		blueMinimum := 0

		fmt.Println(game)
		for _, outcome := range game.outcomes {
			if outcome.red > redMinimum {
				redMinimum = outcome.red
			}
			if outcome.green > greenMinimum {
				greenMinimum = outcome.green
			}
			if outcome.blue > blueMinimum {
				blueMinimum = outcome.blue
			}
		}
		
		totalPower += (redMinimum * greenMinimum * blueMinimum)
		
	}
	return totalPower
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day02")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
