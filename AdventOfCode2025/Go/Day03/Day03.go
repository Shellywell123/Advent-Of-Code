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

		list1 = append(list1, line)
	}

	return list1
}

func Part1(filename string, debug bool) int {
	list1 := Parse(filename)

	joltages := []int{}
	for _, bank := range list1 {
		if debug {
			fmt.Printf("Processing bank: %v\n", bank)
		}

		digit1 := 0
		for i := 0; i < len(bank)-1; i++ {

			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit1 {
				digit1 = digitCheck
			}
		}

		digit2 := 0
		for i := strings.Index(bank, strconv.Itoa(digit1)) + 1; i < len(bank); i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit2 {
				digit2 = digitCheck
			}
		}

		if debug {
			fmt.Printf(" %d%d\n", digit1, digit2)
		}
		joltages = append(joltages, digit1*10+digit2)

	}

	ans := 0
	for _, joltage := range joltages {
		ans += joltage
	}

	return ans
}

func Part2(filename string, debug bool) int {
	list1 := Parse(filename)

	joltages := []int{}
	for _, bank := range list1 {
		if debug {
			fmt.Printf("Processing bank: %v\n", bank)
		}

		// did it the lazy way as short on time
		digit1 := 0
		digit1index := 0
		for i := 0; i < len(bank)-11; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit1 {
				digit1 = digitCheck
				digit1index = i
			}
		}

		digit2 := 0
		digit2index := 0
		for i := digit1index + 1; i < len(bank)-10; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit2 {
				digit2 = digitCheck
				digit2index = i
			}
		}

		digit3 := 0
		digit3index := 0
		for i := digit2index + 1; i < len(bank)-9; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit3 {
				digit3 = digitCheck
				digit3index = i
			}
		}

		digit4 := 0
		digit4index := 0
		for i := digit3index + 1; i < len(bank)-8; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit4 {
				digit4 = digitCheck
				digit4index = i
			}
		}

		digit5 := 0
		digit5index := 0
		for i := digit4index + 1; i < len(bank)-7; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit5 {
				digit5 = digitCheck
				digit5index = i
			}
		}

		digit6 := 0
		digit6index := 0
		for i := digit5index + 1; i < len(bank)-6; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit6 {
				digit6 = digitCheck
				digit6index = i
			}
		}

		digit7 := 0
		digit7index := 0
		for i := digit6index + 1; i < len(bank)-5; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit7 {
				digit7 = digitCheck
				digit7index = i
			}
		}

		digit8 := 0
		digit8index := 0
		for i := digit7index + 1; i < len(bank)-4; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit8 {
				digit8 = digitCheck
				digit8index = i
			}
		}

		digit9 := 0
		digit9index := 0
		for i := digit8index + 1; i < len(bank)-3; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit9 {
				digit9 = digitCheck
				digit9index = i
			}
		}

		digit10 := 0
		digit10index := 0
		for i := digit9index + 1; i < len(bank)-2; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit10 {
				digit10 = digitCheck
				digit10index = i
			}
		}

		digit11 := 0
		digit11index := 0
		for i := digit10index + 1; i < len(bank)-1; i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit11 {
				digit11 = digitCheck
				digit11index = i
			}
		}

		digit12 := 0
		for i := digit11index + 1; i < len(bank); i++ {
			digitCheck, _ := strconv.Atoi(string(bank[i]))
			if digitCheck > digit12 {
				digit12 = digitCheck
			}
		}

		if debug {
			fmt.Printf(" %d%d%d%d%d%d%d%d%d%d%d%d\n", digit1, digit2, digit3, digit4, digit5, digit6, digit7, digit8, digit9, digit10, digit11, digit12)
		}

		joltages = append(joltages, digit1*100000000000+digit2*10000000000+digit3*1000000000+digit4*100000000+digit5*10000000+digit6*1000000+digit7*100000+digit8*10000+digit9*1000+digit10*100+digit11*10+digit12)

	}

	ans := 0
	for _, joltage := range joltages {
		ans += joltage
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day03")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
