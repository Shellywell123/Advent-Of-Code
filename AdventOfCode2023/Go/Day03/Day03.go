package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isInt(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

type PartID struct {
	number int
	y      int
	x      int
}

type Gear struct {
	y int
	x int
}

func Parse(filename string) ([]string, []PartID, []Gear) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	engineSchematic := []string{}
	partIds := []PartID{}
	gears := []Gear{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		for index, char := range line {

			//find gears
			if string(char) == "*" {
				newGear := Gear{}
				newGear.x = index
				newGear.y = y
				gears = append(gears, newGear)
			}

			//find partIds
			if index > 0 && isInt(string(line[index-1])) {
				continue
			}
			if isInt(string(char)) {
				newPart := PartID{}
				number := ""
				for i := 0; i < 3; i++ {
					if isInt(string(line[index+i])) {
						number += string(line[index+i])
					}
				}
				newPart.number, _ = strconv.Atoi(number)
				newPart.y = y
				newPart.x = index
				partIds = append(partIds, newPart)
			}	
		}

		engineSchematic = append(engineSchematic, line)
		y++
	}

	return engineSchematic, partIds, gears
}

func Part1(filename string) int {
	engineSchematic, partIds, _ := Parse(filename)
	total := 0

	for y, line := range engineSchematic {

		// for each partid on that line
		for _, partId := range partIds {
			if partId.y == y {
				xL := partId.x
				xR := partId.x + len(fmt.Sprint(partId.number))

				// check surrounding coords for symbols
				for yCheck := y - 1; yCheck <= y+1; yCheck++ {
					for xCheck := xL - 1; xCheck <= xR; xCheck++ {

						// check out of bounds
						if xCheck < 0 || xCheck >= len(line) || yCheck < 0 || yCheck >= len(engineSchematic) {
							continue
						}

						str := string(engineSchematic[yCheck][xCheck])

						// increase total if adjacent symbol exists
						if (!isInt(str)) && (str != ".") {
							total += partId.number
						}
					}
				}
			}
		}
	}

	return total
}

func Part2(filename string) int {
	engineSchematic, partIds, gears := Parse(filename)
	ratios := map[int]int{}

	for y, line := range engineSchematic {

		// for each gear on that line
		for _, gear := range gears {
			if gear.y == y {

				ids := map[int]int{}
				

				// check surrounding coords for partids
				for yCheck := y - 1; yCheck <= y+1; yCheck++ {
					for xCheck := gear.x - 1; xCheck <= gear.x+1; xCheck++ {

						// check out of bounds
						if xCheck < 0 || xCheck >= len(line) || yCheck < 0 || yCheck >= len(engineSchematic) {
							continue
						}

						str := string(engineSchematic[yCheck][xCheck])

						// look for adjacent part ids 
						if isInt(str) {

							// find leading digit x ind
							leadInd := xCheck
							for i := 0; i < 3; i++ {
								if xCheck -i < 0 || !isInt(string(engineSchematic[yCheck][xCheck-i])) {
									break
								}
								leadInd = xCheck - i
								
							}

							// find part id from index
							for _, partId := range partIds {
								if partId.y == yCheck && partId.x == leadInd {
									ids[partId.number+partId.x-partId.y] = partId.number // use map w/ uid to prevent duplicate entries (not a fan of my messy uid method)
								}
							}

							// if two part ids found get gear ratio
							if len(ids) == 2 {					
								ratio := 1			
								for _,val := range ids {
									ratio *= val
								}
								ratios[ratio+gear.x+gear.y] = ratio // use map w/ uid to prevent duplicate entries
								break
							}

						}
					}
				}
			}
		}
	}

	// sum ratios
	total := 0
	for _,val := range ratios {
		total += val
	}
	return total
}
func main() {
	fmt.Println("Advent-Of-Code 2023 - Day03")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
