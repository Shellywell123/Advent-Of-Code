package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type springs struct {
	data string
	failed []int
}

func Parse(filename string) []springs {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	
	springss := []springs{}

	for fscanner.Scan() {
		line := fscanner.Text()

		springs := springs{}

		springs.data = strings.Split(line, " ")[0]

		fails := []int{}
		for _, num := range strings.Split(strings.Split(line, " ")[1], ",") {
			numm, _ := strconv.Atoi(num)
			fails = append(fails, numm)
		}
		springs.failed = fails

		springss = append(springss, springs)
	}

	return springss
}

func count(cfgI string, nums []int) int {
	
	fmt.Println(cfgI, nums)

	// if the no config is possible return 0
	// when
	// cfg is empty string and nums is a populated list
	// or
	// cfg is populated string and nums is an empty list
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum, len(cfgI))
	if sum > len(cfgI) || (cfgI == "" && len(nums) != 0) || (len(nums) == 0 && strings.Contains(cfgI, "#")) {
		return 0
	}

	// if the 1 config is possible return 1
	// when
	// nums is an empty list and cfg is empty string
	// or 
	// nums is an empty list and cfg doesnt contain #
	if len(nums) == 0 && (cfgI == "" || !strings.Contains(cfgI, "#")) {
		fmt.Print("found")
		return 1
	}

	// handle each ? in 2 ways 
	result := 0

	// 1 - handle .,? as .
	if strings.Contains(".?", string(cfgI[0])){
		fmt.Println(1)
		// recursive count removing the first char
		result += count(string(cfgI[1:]), nums)

	}

	// 2 - handle #,? as #
	if strings.Contains("#?", string(cfgI[0])){
		fmt.Println(2)
		// conditional logic before recursion
		if nums[0] <= len(cfgI) && // lengthwise it is possible for cfg to contains the first num
			!strings.Contains(cfgI[:nums[0]], ".") && // #,? found in a row for size of first num
			(nums[0] == len(cfgI) || string(cfgI[nums[0]]) != "#") {
				// recursive count removing the first match of nums
				result += count(string(cfgI[nums[0]+1]), nums[1:])
			}
	}

	return result
}

func Part1(filename string) int {
	springss := (Parse(filename))
	
	arrangements := 0
	for _, s := range springss {
		arrangements += count(s.data, s.failed)	
	}

	return arrangements
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day12")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	// fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	// fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
