package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1(filename string) int {

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	commmons := []string {}

	// loop through file reading one line at a time
	for fscanner.Scan() {
		rucksack := fscanner.Text()

		// split rucksack into compartments
		compartment1 := strings.Split(rucksack[:len(rucksack)/2],"")
		compartment2 := strings.Split(rucksack[len(rucksack)/2:],"")

		compartment1Map := map[string]string{}
		compartment2Map := map[string]string{}


		// convert compartments into Maps to remove duplicate entries
		for _, item := range compartment1 {
			compartment1Map[item] = item
		}

		for _, item := range compartment2 {
			compartment2Map[item] = item
		}
		
		// collect common elements into "commons"
		for item1, _ := range compartment1Map {
			for item2, _ := range compartment2Map {
				if item1 == item2 {
					commmons = append(commmons, item1)
				}			
			}
		}
	}

	// calculate the total priotity value of all common items
	priority := map[string]int {}
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 52; i++ {
		priority[strings.Split(alphabet,"")[i]] = i+1
	}

	total := 0
	for _, item := range commmons {
		total += priority[item]
	}

	return total
}

func Part2(filename string) int {

	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	commmons := []string {}

	groupmember := 1
	group := []string {}

	// loop through file reading one line at a time
	for fscanner.Scan() {
		rucksack := fscanner.Text()
		group = append(group, rucksack)

		if groupmember == 3 {

			member1Rucksack := strings.Split(group[0],"")
			member2Rucksack := strings.Split(group[1],"")
			member3Rucksack := strings.Split(group[2],"")

			// convert member rucksacks into Maps to remove duplicate entries
			member1MapRucksack := map[string]string{}
			member2MapRucksack := map[string]string{}
			member3MapRucksack := map[string]string{}

			for _, item := range member1Rucksack {
				member1MapRucksack[item] = item
			}

			for _, item := range member2Rucksack {
				member2MapRucksack[item] = item
			}

			for _, item := range member3Rucksack {
				member3MapRucksack[item] = item
			}
			
			// collect common elements into "commons"
			for item1, _ := range member1MapRucksack {
				for item2, _ := range member2MapRucksack {
					for item3, _ := range member3MapRucksack {
						if item1 == item2 && item2 == item3 {
							commmons = append(commmons, item1)
						}
					}
				}
			}
			group = nil
			groupmember = 0
		}
		groupmember += 1	
		
	}

	// calculate the total priotity value of all common items
	priority := map[string]int {}
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 52; i++ {
		priority[strings.Split(alphabet,"")[i]] = i+1
	}

	total := 0
	for _, item := range commmons {
		total += priority[item]
	}

	return total
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day03")
	fmt.Printf("Tests : Answer to Part 1 = %d\n", Part1(testfile))	
	fmt.Printf("Inputs: Answer to Part 1 = %d\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %d\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %d\n", Part2(inputfile))
}
