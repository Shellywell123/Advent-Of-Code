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

	data := []string{}
	for fscanner.Scan() {
		line := fscanner.Text()
		data = strings.Split(line, ",")
	}
	return data
}

func Calculate(hash string) int {
	currentValue := 0
	for _, r := range hash {
		currentValue += int(r)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func Part1(filename string) int {
	data := Parse(filename)

	sum := 0
	for _, hash := range data {
		sum += Calculate(hash)
	}

	return sum
}

type Lens struct {
	slot        int
	focalLength int
}

func (l *Lens) SetfocalLength(focalLength int) {
	l.focalLength = focalLength
}

func (l *Lens) SetSlot(slot int) {
	l.slot = slot
}

func Part2(filename string) int {
	data := Parse(filename)

	boxHashmaps := map[int]map[string]Lens{}

	for _, hash := range data {

		operation := "="
		if strings.Contains(hash, "-") {
			operation = "-"

		}

		hashSplit := strings.Split(hash, operation)
		label := hashSplit[0]
		value, _ := strconv.Atoi(hashSplit[1])
		box := Calculate(label)

		switch operation {

		// remove lens
		case "-":

			// shuffle all other slots
			s := boxHashmaps[box][label].slot

			// remove lens
			delete(boxHashmaps[box], label)

			for lensLabel, _ := range boxHashmaps[box] {
				if s > 0 && boxHashmaps[box][lensLabel].slot > s {
					lenso := boxHashmaps[box][lensLabel]
					lenso.SetSlot(lenso.slot - 1)
					boxHashmaps[box][lensLabel] = lenso
				}
			}

		// insert/replace lens
		case "=":

			// create a new box if needed
			if boxHashmaps[box] == nil {
				boxHashmaps[box] = map[string]Lens{}
			}

			// create a new lens if needed
			if _, exists := boxHashmaps[box][label]; !exists {
				boxHashmaps[box][label] = Lens{len(boxHashmaps[box]) + 1, 0}
			}

			// replace lense
			lenso := boxHashmaps[box][label]
			lenso.SetfocalLength(value)
			boxHashmaps[box][label] = lenso
		}
	}
	sum := 0
	for boxNumber := 1; boxNumber <= 256; boxNumber++ {
		for lensLabel, _ := range boxHashmaps[boxNumber] {
			sum += (boxNumber + 1) * boxHashmaps[boxNumber][lensLabel].slot * boxHashmaps[boxNumber][lensLabel].focalLength
		}
	}

	return sum
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day15")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
