package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) [][]string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	megaagrid := [][]string{}

	for fscanner.Scan() {

		line := fscanner.Text()
		megaagrid = append(megaagrid, strings.Split(line, ""))
	}

	return megaagrid
}

func Part1(filename string, debug bool) int {
	grid := Parse(filename)

	// find start position
	current_x := 0
	for x := 0; x < len(grid[0]); x++ {
		if grid[0][x] == "S" {
			current_x = x
			break
		}
	}

	current_y := 0
	current_xs := []int{current_x}
	ans := 0

	for {

		if debug {
			fmt.Println("Current Position:", current_xs, current_y)
			for y := 0; y < len(grid); y++ {
				fmt.Println(grid[y])
			}
		}

		next_current_xs := []int{}
		current_y++
		if current_y == len(grid) {
			break
		}
		for _, current_x := range current_xs {

			if grid[current_y][current_x] == "^" {
				ans++
				next_current_xs = append(next_current_xs, current_x+1)
				next_current_xs = append(next_current_xs, current_x-1)
				grid[current_y][current_x+1] = "|"
				grid[current_y][current_x-1] = "|"
				continue
			}

			if grid[current_y][current_x] == "." {
				next_current_xs = append(next_current_xs, current_x)
				grid[current_y][current_x] = "|"
				continue
			}
		}

		current_xs = next_current_xs
	}

	return ans
}

// i hate this code, but I am hungover and it works
func Part2(filename string, debug bool) int {
	grid := Parse(filename)

	// find start position
	current_x := 0
	for x := 0; x < len(grid[0]); x++ {
		if grid[0][x] == "S" {
			current_x = x
			break
		}
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "^" {
				grid[y][x] = "0"
			}
		}
	}

	if debug {
		for y := 0; y < len(grid); y++ {
			fmt.Println(grid[y])
		}
	}
	current_y := 0

	current_xs := []int{current_x}
	for {
		current_y++
		if debug {
			fmt.Println("Current Y:", current_y)
		}

		if debug {
			fmt.Println("Current Position:", current_x, current_y)
			for y := 0; y < len(grid); y++ {
				fmt.Println(grid[y])
			}
		}

		if current_y == len(grid) {
			break
		}
		new_current_xs := []int{}
		current_keep_xs := []int{}
		for _, current_x := range current_xs {

			if grid[current_y][current_x] == "." {
				grid[current_y][current_x] = "|"
				current_keep_xs = append(current_keep_xs, current_x)
				continue
			}

			if grid[current_y][current_x] == "0" {

				passes := 0
				if current_y == 2 {
					passes = 1
				} else {

					for y := current_y - 2; y >= 0; y-- {
						if grid[y][current_x] == "." {
							break
						}
						if grid[y][current_x+1] != "|" {
							p, _ := strconv.Atoi(grid[y][current_x+1])
							passes += p
						}
					}
					for y := current_y - 2; y >= 0; y-- {
						if grid[y][current_x] == "." {
							break
						}
						if grid[y][current_x-1] != "|" && grid[y][current_x-1] != "." {
							p, _ := strconv.Atoi(grid[y][current_x-1])
							passes += p
						}
					}
				}

				grid[current_y][current_x] = fmt.Sprintf("%d", passes)

				grid[current_y][current_x+1] = "|"
				new_current_xs = append(new_current_xs, current_x+1)

				grid[current_y][current_x-1] = "|"
				new_current_xs = append(new_current_xs, current_x-1)

			}

		}
		current_xs = new_current_xs
		current_xs = append(current_xs, current_keep_xs...)
	}

	ans := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != "." && grid[y][x] != "|" {
				passes, _ := strconv.Atoi(grid[y][x])
				ans += passes
			}
		}
	}

	return ans + 1
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day07")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
