package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type robot struct {
	p coord
	v coord
}

func parse(filename string) []robot {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	robots := []robot{}
	for fscanner.Scan() {
		line := fscanner.Text()

		p := []int{}
		for _, n := range strings.Split(strings.Split(strings.Split(line, " ")[0], "=")[1], ",") {
			pp, _ := strconv.Atoi(n)
			p = append(p, pp)
		}

		v := []int{}
		for _, n := range strings.Split(strings.Split(strings.Split(line, " ")[1], "=")[1], ",") {
			vv, _ := strconv.Atoi(n)
			v = append(v, vv)
		}
		robots = append(robots, robot{coord{p[0], p[1]}, coord{v[0], v[1]}})
	}

	return robots
}

func seconds(robots []robot, bathroom coord, seconds int) []robot {

	bathroomWidth := bathroom.x
	bathroomHeight := bathroom.y

	newRobots := []robot{}
	for _, r := range robots {


		newPx := (r.p.x + r.v.x * seconds ) % bathroomWidth
		newPy := (r.p.y + r.v.y * seconds ) % bathroomHeight

		if newPx < 0 {
			newPx += bathroomWidth
		}

		if newPy >= bathroomHeight {
			newPy -= bathroomHeight
		}

		if newPy < 0 {
			newPy += bathroomHeight
		}

		if newPx >= bathroomWidth {
			newPx -= bathroomWidth
		}

		newRobots = append(newRobots, robot{coord{newPx, newPy}, r.v})
	}

	return newRobots
}

func quadrants(robots []robot, bathroom coord) int {

	bathroomWidth := bathroom.x
	bathroomHeight := bathroom.y

	xMid := int(float64(bathroomWidth) / 2)
	yMid := int(float64(bathroomHeight) / 2)

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, r := range robots {

		if r.p.x < xMid && r.p.y < yMid {
			q1++
		}

		if r.p.x > xMid && r.p.y < yMid {
			q2++
		}

		if r.p.x < xMid && r.p.y > yMid {
			q3++
		}

		if r.p.x > xMid && r.p.y > yMid {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func printRobots(robots []robot, bathroom coord) {
	bathroomWidth := bathroom.x
	bathroomHeight := bathroom.y

	f := 0
	for y := 0; y < bathroomHeight; y++ {
		for x := 0; x < bathroomWidth; x++ {
			found := false
			for _, r := range robots {
				if r.p.x == x && r.p.y == y {
					fmt.Print("#")
					found = true
					f++
					break
				}
			}
			if !found {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	// had a hunch that this would be the answer
	if f == len(robots) {
		fmt.Println("All found")
		os.Exit(0)
	}
}

func Part1(filename string, bathroom coord, debug bool) int {
	robots := parse(filename)

	robots = seconds(robots, bathroom, 100)
	ans := quadrants(robots, bathroom)
	return ans
}

func Part2(filename string, bathroom coord, debug bool) int {
	robots := parse(filename)
	printRobots(robots, bathroom)

	s := 0
	for {
		printRobots(robots, bathroom)
		robots = seconds(robots, bathroom, 1)	
		s++	
		fmt.Println("Seconds: ", s)
	}
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day14")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", coord{11, 7}, false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", coord{101, 103}, false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", coord{101, 103}, false))
}
