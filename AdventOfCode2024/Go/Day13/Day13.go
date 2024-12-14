package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x float64
	y float64
}

type machine struct {
	a coord
	b coord
	p coord
}

func parse(filename string) []machine {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	machines := []machine{}
	newMachine := machine{}
	for fscanner.Scan() {
		line := fscanner.Text()

		if line == "" {
			machines = append(machines, newMachine)
			newMachine = machine{}
		}

		if strings.Split(line, ":")[0] == "Button A" {
			aX, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[1], ",")[0], "+")[1])
			aY, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[1], ",")[1], "+")[1])
			newMachine.a = coord{float64(aX), float64(aY)}
		}

		if strings.Split(line, ":")[0] == "Button B" {
			bX, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[1], ",")[0], "+")[1])
			bY, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[1], ",")[1], "+")[1])
			newMachine.b = coord{float64(bX), float64(bY)}
		}

		if strings.Split(line, ":")[0] == "Prize" {
			pX, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[1], ",")[0], "=")[1])
			pY, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, ":")[1], ",")[1], "=")[1])
			newMachine.p = coord{float64(pX), float64(pY)}
		}
	}
	machines = append(machines, newMachine)

	return machines
}

func tokensToWin(m machine, constant float64) int {

	px := m.p.x + (constant)
	py := m.p.y + (constant)

	// SO INT() IS THE SAME AS FLOOR() IN GO
	aPresses := math.Round((px - (m.b.x*m.p.y)/m.b.y) / (m.a.x - (m.b.x*m.a.y)/m.b.y))
	bPresses := math.Round((py - (m.a.y * aPresses)) / m.b.y)

	// cant have negative presses
	if aPresses < 0 || bPresses < 0 {
		return 0
	}

	// as we are using floats we need to check if the presses are whole numbers
	// easist way to do this is to assert if the eqs work
	if ((aPresses*m.a.x)+(bPresses*m.b.x) != px) || ((aPresses*m.a.y)+(bPresses*m.b.y) != py) {
		return 0
	}
	return int(aPresses*3 + bPresses)
}

func Part1(filename string, debug bool) int {
	machines := parse(filename)

	if debug {
		fmt.Println(machines)
	}
	ans := 0
	for _, m := range machines {
		ans += tokensToWin(m, 0)
	}

	return ans
}

func Part2(filename string, debug bool) int {
	machines := parse(filename)

	if debug {
		fmt.Println(machines)
	}
	ans := 0
	for _, m := range machines {
		ans += tokensToWin(m, 10000000000000)
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day13")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Println("had to use python for part 2 :(")
}
