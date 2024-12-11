package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) map[int]int {
	contentBytes, _ := os.ReadFile(filename)
	content := string(contentBytes)

	// note there were no duplicates in my input
	stones := map[int]int{}
	for _, c := range strings.Split(content, " ") {
		n, _ := strconv.Atoi(string(c))
		stones[n] = 1
	}

	return stones
}

func blink(stones map[int]int, blinks int, debug bool) map[int]int {

	if debug {
		ansTemp := 0
		for _, v := range stones {
			ansTemp += v
		}
		fmt.Println(blinks, ansTemp, stones)
	}

	if blinks == 0 {
		return stones
	}

	newStones := map[int]int{}
	for stone, n := range stones {

		stoneString := fmt.Sprintf("%d", stone)

		if stone == 0 {
			_, ok := newStones[1]
			if ok {
				newStones[1] += n
			} else {
				newStones[1] = n
			}
			continue

		}

		if len(stoneString)%2 == 0 {

			left, _ := strconv.Atoi(stoneString[:(len(stoneString)+1)/2])
			right, _ := strconv.Atoi(stoneString[(len(stoneString)+1)/2:])

			_, okl := newStones[left]
			if okl {
				newStones[left] += n
			} else {
				newStones[left] = n
			}

			_, okr := newStones[right]
			if okr {
				newStones[right] += n
			} else {
				newStones[right] = n
			}

			continue
		}

		_, oke := newStones[stone*2024]
		if oke {
			newStones[stone*2024] += n
		} else {
			newStones[stone*2024] = n
		}
	}

	return blink(newStones, blinks-1, debug)
}

// trick to this one is to leverage maps
func solve(filename string, debug bool, blinks int) int {
	stones := parse(filename)

	ans := 0
	finalStones := blink(stones, blinks, debug)
	for _, n := range finalStones {
		ans += n
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day11")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", solve("tests.txt", false, 25))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", solve("inputs.txt", false, 25))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", solve("inputs.txt", false, 75))
}
