package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(filename string) ([][]int, []string) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	numbers := [][]int{}
	numbers_sorted := [][]int{}
	operators := []string{}

	for fscanner.Scan() {
		line := fscanner.Text()
		if line[0] == '*' || line[0] == '+' {
			for _, op := range strings.Split(line, " ") {
				if op == "" {
					continue
				}
				operators = append(operators, op)
			}
			break
		}

		numbers_temp := []int{}
		line = strings.Replace(line, "   ", " ", -1)
		line = strings.Replace(line, "  ", " ", -1)

		for _, char := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(char)
			if num == 0 {
				continue
			}
			numbers_temp = append(numbers_temp, num)
		}
		numbers = append(numbers, numbers_temp)

	}
	fmt.Println("lol")
	fmt.Println(len(numbers[0]))
	fmt.Println(len(operators)) // this should = 1000 not 591

	for i, _ := range operators {
		numbers_sorted = append(numbers_sorted, []int{})
		for j, _ := range numbers {
			numbers_sorted[i] = append(numbers_sorted[i], numbers[j][i])
		}
	}

	return numbers_sorted, operators
}

func Part1(filename string, debug bool) int {
	numberss, operators := Parse(filename)

	fmt.Println(len(numberss))
	fmt.Println(len(operators))

	fmt.Println(numberss)
	fmt.Println(operators)

	ans := 0
	for i, numbers := range numberss {

		mult := 1
		if operators[i] == "*" {
			if debug {
				fmt.Println("")
			}
			for _, num := range numbers {
				if debug {
					fmt.Print(num, " * ")
				}
				mult *= num
			}
			if debug {
				fmt.Print("=", mult)
			}
			ans += mult
			continue
		}

		add := 0
		if operators[i] == "+" {
			if debug {
				fmt.Println("")
			}
			for _, num := range numbers {
				if debug {
					fmt.Print(num, " + ")
				}
				add += num
			}
			if debug {
				fmt.Print("=", add)
			}
			ans += add
			continue
		}
	}
	if debug {
		fmt.Println("")
	}

	for _, n := range numberss {
		fmt.Print(n[0], " ")
	}
	return ans
}

func Parse2(filename string) ([][]int, []string) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	operators := []string{}
	megaagrid := [][]string{}

	for fscanner.Scan() {

		line := fscanner.Text()

		if line[0] == '*' || line[0] == '+' {
			for _, op := range strings.Split(line, " ") {
				if op == "" {
					continue
				}
				operators = append(operators, op)
			}
			break
		}
		megaagrid = append(megaagrid, strings.Split(line, ""))
	}

	prev_sep := -1

	numbers_sorted := [][]int{}
	for x := 0; x <= len(megaagrid[0]); x++ {
		sep_found := false
		if x == len(megaagrid[0]) { // make sure to grab last one fully 
			sep_found = true
		}
		if sep_found == false && megaagrid[0][x] == " " {
			counter := 1
			for y := 1; y < len(megaagrid); y++ {
				if megaagrid[y][x] != " " {
					break
				}
				counter++

			}
			if counter == len(megaagrid) {
				sep_found = true
			}
		}
		// get number

		if sep_found == false {
			continue
		}

		mini_grid := [][]string{}
		for y := 0; y < len(megaagrid); y++ {
			row := []string{}
			for k := prev_sep + 1; k < x; k++ {
				row = append(row, megaagrid[y][k])
			}
			mini_grid = append(mini_grid, row)
		}
		prev_sep = x

		// read vertical nums for grid
		fmt.Println("mini grid:")
		fmt.Println(mini_grid)
		numbers_sorted_mini := []int{}
		for i := len(mini_grid[0]) - 1; i >= 0; i-- {
			new_string_num := ""
			for j := 0; j < len(mini_grid); j++ {
				new_string_num += mini_grid[j][i]
			}
			new_num, _ := strconv.Atoi(strings.TrimSpace(new_string_num))
			numbers_sorted_mini = append(numbers_sorted_mini, new_num)
		}
		numbers_sorted = append(numbers_sorted, numbers_sorted_mini)

	}

	return numbers_sorted, operators
}

func Part2(filename string, debug bool) int {
	numberss, operators := Parse2(filename)

	fmt.Println(len(numberss))
	fmt.Println(len(operators))

	fmt.Println(numberss)
	fmt.Println(operators)

	ans := 0
	for i, numbers := range numberss {

		mult := 1
		if operators[i] == "*" {
			if debug {
				fmt.Println("")
			}
			for _, num := range numbers {
				if debug {
					fmt.Print(num, " * ")
				}
				mult *= num
			}
			if debug {
				fmt.Print("=", mult)
			}
			ans += mult
			continue
		}

		add := 0
		if operators[i] == "+" {
			if debug {
				fmt.Println("")
			}
			for _, num := range numbers {
				if debug {
					fmt.Print(num, " + ")
				}
				add += num
			}
			if debug {
				fmt.Print("=", add)
			}
			ans += add
			continue
		}
	}
	if debug {
		fmt.Println("")
	}

	for _, n := range numberss {
		fmt.Print(n, " ")
	}
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day06")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false)) // 2862649984428023 = toohigh
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", true))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
