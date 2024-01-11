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

// Hell Yeah, Dynamic Programming Baby!
var hashMap = make(map[string]int)

func count(cfg string, cfgIndex int, nums []int, numsIndex int) int {

	cfgTemp := ""
	numsTemp := []int{}


	if cfgIndex < len(cfg) {
		cfgTemp = string(cfg[cfgIndex:])
	}

	if numsIndex < len(nums) {
		numsTemp = nums[numsIndex:]
	}
	
	// if the no config is possible return 0
	// when
	// cfgTemp is empty string and numsTemp is a populated list
	// or
	// cfgTemp is populated string and numsTemp is an empty list
	sum := 0
	for _, n := range numsTemp {
		sum += n
	}
	// fmt.Println(sum, len(cfgTemp))
	if sum > len(cfgTemp) ||  (cfgTemp == "" && len(numsTemp) != 0) || (len(numsTemp) == 0 && strings.Contains(cfgTemp, "#")) {
		return 0
	}

	// if the 1 config is possible return 1
	// when
	// numsTemp is an empty list and cfgTemp is empty string
	// or 
	// numsTemp is an empty list and cfgTemp doesnt contain #
	if len(numsTemp) == 0 && (cfgTemp == "" || !strings.Contains(cfgTemp, "#")) {
		return 1
	}

	key := fmt.Sprintf("%s,%d,%d,%d", cfg, nums, cfgIndex, numsIndex)
	if v, ok := hashMap[key]; ok {
		return v
	}

	// handle each ? in 2 ways 
	result := 0

	// 1 - handle .,? as .
	if strings.Contains(".?", string(cfgTemp[0])){
		// recursive count removing the first char
		result += count(cfg, cfgIndex+1, nums, numsIndex)
	}

	// 2 - handle #,? as #
	if strings.Contains("#?", string(cfgTemp[0])) &&
	    numsTemp[0] <= len(cfgTemp) && // lengthwise it is possible for cfgTemp to contains the first num
		!strings.Contains(cfgTemp[:numsTemp[0]], ".") && // #,? found in a row for size of first num
		(numsTemp[0] == len(cfgTemp) || string(cfgTemp[numsTemp[0]]) != "#") {
			// recursive count removing the first match of numsTemp
			result += count(cfg, cfgIndex+numsTemp[0]+1, nums, numsIndex+1)
	}

	hashMap[key] = result
	return result
}

func ExpandSpringss(springss []springs) []springs {
	springssExpanded := []springs{}

	for i, _ := range springss {
		d := springss[i].data
		f := springss[i].failed

		c := 0
		for c < 4 {
			d += "?"+springss[i].data
			for j := 0; j < len(springss[i].failed); j++ {
				f = append(f,springss[i].failed[j])
			}
			c++
		}
		springssExpanded = append(springssExpanded, springs{d,f})
	}
	return springssExpanded
}

func Part1(filename string) int {
	springss := (Parse(filename))
	
	arrangements := 0
	for _, s := range springss {
		arrangements += count(s.data, 0, s.failed, 0)	
	}

	return arrangements
}

func Part2(filename string) int {
	springss := (Parse(filename))

	arrangements := 0
	for _, s := range ExpandSpringss(springss) {
		arrangements += count(s.data, 0, s.failed, 0)
	}

	return arrangements
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day12")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
