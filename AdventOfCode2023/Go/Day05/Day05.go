package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeConfig struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLength      int
}

type SourceToDestinationMap struct {
	source string
	destination   string
	config []RangeConfig
}

func isInt(char byte) bool {
	if _, err := strconv.Atoi(string(char)); err == nil {
		return true
	}
	return false
}

func Parse(filename string) ([]int, []SourceToDestinationMap) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	seeds := []int{}
	sourceToDestinationMaps := []SourceToDestinationMap{}
	sourceToDestinationMap := SourceToDestinationMap{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		if y == 0 {
			for _, num := range strings.Split(line[7:], " ") {
				seed, _ := strconv.Atoi(string(num))
				seeds = append(seeds, seed)
			}
		} else {
			if len(line) != 0 && !isInt(line[0]) {
				sourceToDestinationMap = SourceToDestinationMap{}
				sourceToDestinationMap.source = strings.Split(strings.Split(line, " ")[0], "-")[0]
				sourceToDestinationMap.destination = strings.Split(strings.Split(line, " ")[0], "-")[2]
			}
			if  len(line) != 0 && isInt(line[0]) {
				rangeNumbers := strings.Split(line, " ")
				rangeConfig := RangeConfig{}
				rangeConfig.destRangeStart, _ = strconv.Atoi(rangeNumbers[0])
				rangeConfig.sourceRangeStart, _ = strconv.Atoi(rangeNumbers[1])
				rangeConfig.rangeLength, _ = strconv.Atoi(rangeNumbers[2])
				sourceToDestinationMap.config = append(sourceToDestinationMap.config, rangeConfig)
			}
			if len(line) == 0 && sourceToDestinationMap.source != ""{
				sourceToDestinationMaps = append(sourceToDestinationMaps, sourceToDestinationMap)
			}
		}
		y++
	}
	sourceToDestinationMaps = append(sourceToDestinationMaps, sourceToDestinationMap)

	return seeds, sourceToDestinationMaps
}

func MapFunction(object int, rangeConfigs []RangeConfig) int {

	for _, rangeConfig := range rangeConfigs {
		if (rangeConfig.sourceRangeStart <= object) && object < (rangeConfig.sourceRangeStart + rangeConfig.rangeLength) {
			return  rangeConfig.destRangeStart + (object - rangeConfig.sourceRangeStart)
		}
	}
	return object
}

func Recursion(object int, destination string, sourceToDestinationMaps []SourceToDestinationMap  ) int {
	for _, destinationMap := range sourceToDestinationMaps {
		if destinationMap.source == destination {
			for _, sourceMap := range sourceToDestinationMaps {
				if sourceMap.destination == destinationMap.source {
					return Recursion(MapFunction(object, sourceMap.config), destinationMap.destination, sourceToDestinationMaps)
				}
			}
		}
	}
	return object
}

func Part1(filename string) int {
	seeds, sourceToDestinationMaps := Parse(filename)

	// dirty hack to recurse through the final map
	mapIWant := SourceToDestinationMap{}
	for _, sourceToDestinationMap := range sourceToDestinationMaps {
		if sourceToDestinationMap.destination == "location" {
			mapIWant = sourceToDestinationMap
		}
	}

	lowestLocation := 0
	for _, seed := range seeds {
		location := MapFunction(Recursion(seed, "soil", sourceToDestinationMaps), mapIWant.config)
		if location < lowestLocation || lowestLocation == 0 {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func Part2(filename string) int {
	seedRanges, sourceToDestinationMaps := Parse(filename)
	
	// dirty hack to recurse through the final map
	mapIWant := SourceToDestinationMap{}
	for _, sourceToDestinationMap := range sourceToDestinationMaps {
		if sourceToDestinationMap.destination == "location" {
			mapIWant = sourceToDestinationMap
		}
	}

	lowestLocation := 0
	for i := 0; i <= len(seedRanges)-1; i+=2 {
		fmt.Println(seedRanges[i], seedRanges[i+1])
		for seed := seedRanges[i]; seed < seedRanges[i] + seedRanges[i+1]; seed++ {
			location := MapFunction(Recursion(seed, "soil", sourceToDestinationMaps), mapIWant.config)
			if location < lowestLocation || lowestLocation == 0 {
				lowestLocation = location
			}
		}
	}

	return lowestLocation
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day05")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}