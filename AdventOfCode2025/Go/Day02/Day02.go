package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) []string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	list1 := []string{}

	for fscanner.Scan() {
		line := fscanner.Text()

		list1 = append(list1, strings.Split(line, ",")...)
	}

	return list1
}

func Part1(filename string, debug bool) int {
	list1 := Parse(filename)

	invalidIds := []int{}

	for _, v := range list1 {
		parts := strings.SplitN(v, "-", 2)
		id1, _ := strconv.Atoi(parts[0])
		id2, _ := strconv.Atoi(parts[1])

		if debug {
			fmt.Printf("Checking IDs from %v to %v\n", id1, id2)
		}

		// ids to check
		for i := id1; i <= id2; i++ {
			idstring := strconv.Itoa(i)

			if idstring[len(idstring)/2:] == idstring[:len(idstring)/2] {
				invalidIds = append(invalidIds, i)
			}
		}

	}

	ans := 0

	for _, v := range invalidIds {
		ans += v
	}

	return ans
}

func Part2(filename string, debug bool) int {
	list1 := Parse(filename)

	invalidIds := []int{}

	for _, v := range list1 {
		if debug {
			fmt.Printf("Processing range: %v\n", v)
		}
		parts := strings.SplitN(v, "-", 2)
		id1, _ := strconv.Atoi(parts[0])
		id2, _ := strconv.Atoi(parts[1])

		if debug {
			fmt.Printf("Checking IDs from %v to %v\n", id1, id2)
		}

		// ids to check
		for i := id1; i <= id2; i++ {
			idstring := strconv.Itoa(i)

			if debug {
				fmt.Printf("Checking ID: %v\n", i)
			}

			for j := len(idstring); j > 0; j-- {

				if debug {
					fmt.Printf("  Checking substring length %v\n", j)
				}

				if debug {
					fmt.Println("   Checking for ", string(idstring[:len(idstring)-j+1]), " appears in ", idstring, " this times", j)
				}

				substring := string(idstring[:len(idstring)-j])
				count := strings.Count(idstring, substring)

				if debug {
					fmt.Println("   Count is ", strings.Count(idstring, string(idstring[:len(idstring)-j+1])))
				}

				if count*len(substring) == len(idstring) {
					invalidIds = append(invalidIds, i)

					if debug {
						fmt.Printf("   ID %v is invalid as made up of %v repeated %v times\n", i, substring, count)
					}

					break
				}
			}
		}
	}

	ans := 0

	for _, v := range invalidIds {
		ans += v
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day02")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
