package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	number int
	shape  [][]string
}

type Space struct {
	identifier int
	width      int
	height     int
	quatnities []int
}

func Parse(filename string) (map[int]Present, map[int]Space) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	presents := [][][]string{}
	spaces := [][]string{}
	presentsMap := map[int]Present{}
	spaceMap := map[int]Space{}

	currentPresent := [][]string{}

	for fscanner.Scan() {

		line := fscanner.Text()

		if strings.Contains(line, "x") {
			space := []string{}
			space = append(space, strings.Split(line, ": ")[0])
			space = append(space, strings.Split(strings.Split(line, ": ")[1], " ")...)
			spaces = append(spaces, space)
			continue
		}

		if strings.Contains(line, "#") {
			currentPresent = append(currentPresent, strings.Split(line, ""))
			continue
		}

		if line == "" {
			presents = append(presents, currentPresent)
			currentPresent = [][]string{}
			continue
		}
	}

	// fmt.Println(presents)
	for i, present := range presents {
		presentsMap[i] = Present{
			number: i,
			shape:  present,
		}
	}

	for i, space := range spaces {
		widthHeight := strings.Split(string(space[0]), "x")
		// fmt.Println(widthHeight, space[0])
		width, _ := strconv.Atoi(widthHeight[0])
		height, _ := strconv.Atoi(widthHeight[1])

		quatnities := []int{}
		for j := 1; j < len(space); j++ {
			qty, _ := strconv.Atoi(string(space[j]))
			quatnities = append(quatnities, qty)
		}

		spaceMap[i] = Space{
			identifier: i,
			width:      width,
			height:     height,
			quatnities: quatnities,
		}

	}

	return presentsMap, spaceMap
}

func checkPresentsFitsInSpace(spaceGrid [][]string, present Present, presentQty int, debug bool) bool {

	if presentQty == 0 {
		return true
	}

	// Try to fit present in space in all possible rotations and positions
	presentHeight := len(present.shape)
	presentWidth := len(present.shape[0])

	presentFits := false

	for yOffset := 0; yOffset <= len(spaceGrid)-presentHeight; yOffset++ {
		for xOffset := 0; xOffset <= len(spaceGrid[0])-presentWidth; xOffset++ {

			fitsHere := true
			for py := 0; py < presentHeight; py++ {
				for px := 0; px < presentWidth; px++ {
					if present.shape[py][px] == "#" {
						if spaceGrid[yOffset+py][xOffset+px] == "#" {
							fitsHere = false
							break
						}
					}
				}
				if !fitsHere {
					break
				}
			}

			if fitsHere {
				// Mark the cells as occupied
				for py := 0; py < presentHeight; py++ {
					for px := 0; px < presentWidth; px++ {
						if present.shape[py][px] == "#" {
							spaceGrid[yOffset+py][xOffset+px] = "#"
						}
					}
				}

				// Try to fit remaining presents
				if checkPresentsFitsInSpace(spaceGrid, present, presentQty-1, debug) {
					return true
				}

				// Backtrack - unmark the cells
				for py := 0; py < presentHeight; py++ {
					for px := 0; px < presentWidth; px++ {
						if present.shape[py][px] == "#" {
							spaceGrid[yOffset+py][xOffset+px] = "."
						}
					}
				}
			}
		}
		if presentFits {
			break
		}
	}

	return presentFits
}

func presentsFit(presentsMap map[int]Present, space Space, debug bool) bool {
	for presentID, presentQty := range space.quatnities {
		if presentQty > 1 {

			spaceGrid := [][]string{}
			for h := 0; h < space.height; h++ {
				row := []string{}
				for w := 0; w < space.width; w++ {
					row = append(row, ".")
				}
				spaceGrid = append(spaceGrid, row)
			}

			if !checkPresentsFitsInSpace(spaceGrid, presentsMap[presentID], presentQty, debug) {
				return false
			}
			if debug {
				fmt.Println("Present", presentID, "fits in space", space.identifier)
			}
		}
	}
	return true
}

func Part1(filename string, debug bool) int {
	presentsMap, spaceMap := Parse(filename)

	ans := 0
	for _, space := range spaceMap {
		if presentsFit(presentsMap, space, debug) {
			ans++
			if debug {
				fmt.Println(":) All presents fit in space", space.identifier, space.quatnities)
			}
		} else {
			if debug {
				fmt.Println(":( Not all presents fit in space", space.identifier, space.quatnities)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day12")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", true)) // ans should be 2
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", true)) // ans is lower than 1000
	// fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests2.txt", false))
	// fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
