package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	number                 int
	items                  []int
	operation              string
	divisible_test         int
	if_test_true_throw_to  int
	if_test_false_throw_to int
	inspections            int
}

func Parse(filename string) []Monkey {

	file, _ := os.ReadFile(filename)

	data := strings.Split(string(file), "\n")

	monkeys := []Monkey{}
	currentMonkey := Monkey{}

	for i, _ := range data {

		line := strings.TrimSpace(data[i])
		first4 := ""

		if line != "" {
			first4 = line[:4]
		} else {
			first4 = ""
		}

		switch first4 {

		//new monkey
		case "Monk":
			currentMonkeyNumber, _ := strconv.Atoi(strings.Split(line, " ")[1][:1])
			currentMonkey = Monkey{number: currentMonkeyNumber}

		// starting items
		case "Star":
			items := strings.Split(strings.Split(line, ": ")[1], ", ")
			itemsInts := []int{}
			for _, item := range items {
				itemInt, _ := strconv.Atoi(item)
				itemsInts = append(itemsInts, itemInt)
			}
			currentMonkey.items = itemsInts

		// operation
		case "Oper":
			op := strings.Split(strings.Split(line, ": ")[1], "new = old ")[1]
			currentMonkey.operation = op

		// test case
		case "Test":
			test := strings.Split(strings.Split(line, ": ")[1], "y ")[1]
			testInt, _ := strconv.Atoi(test)
			currentMonkey.divisible_test = testInt

		// true test case
		case "If t":
			testTrue := strings.Split(strings.Split(line, ": ")[1], "y ")[1]
			testTrueInt, _ := strconv.Atoi(testTrue)
			currentMonkey.if_test_true_throw_to = testTrueInt

		// false test case
		case "If f":
			testFalse := strings.Split(strings.Split(line, ": ")[1], "y ")[1]
			testFalseInt, _ := strconv.Atoi(testFalse)
			currentMonkey.if_test_false_throw_to = testFalseInt

		//end of monkey
		case "":
			monkeys = append(monkeys, currentMonkey)
			continue
		}

	}
	// get final monkey
	monkeys = append(monkeys, currentMonkey)

	return monkeys
}

func Solve(monkeys []Monkey, rounds int, decreaseWorry func(int, []Monkey) int, debug bool) int {
	
	for r := 0; r < rounds; r++ {
		if debug {
			fmt.Print("\n\nROUND ", r)
		}
		for _, monkey := range monkeys {
			for _, item := range monkey.items {

				// opersation - messy as no eval function in go
				operation := monkey.operation
				if strings.Contains(monkey.operation, "old") {
					operation = strings.ReplaceAll(operation, "old", strconv.Itoa(item))
				}

				coeficent, _ := strconv.Atoi(operation[2:])

				switch operation[:1] {
				case "+":
					item = item + coeficent
				case "-":
					item = item - coeficent
				case "*":
					item = item * coeficent
				case "/":
				}

				// decrease worry
				item = decreaseWorry(item, monkeys)

				// test
				if item%monkey.divisible_test == 0 {
					monkeys[monkey.if_test_true_throw_to].items = append(monkeys[monkey.if_test_true_throw_to].items, item)

				} else {
					monkeys[monkey.if_test_false_throw_to].items = append(monkeys[monkey.if_test_false_throw_to].items, item)
				}
				monkeys[monkey.number].inspections++
			}
			monkeys[monkey.number].items = []int{}
		}
		if debug {
			for _, monkey := range monkeys {
				fmt.Print("\nMonkey ", monkey.number, ": ")
				for _, item := range monkey.items {
					fmt.Print(item, " ")
				}
			}
		}
	}

	// calculate monkey business
	monkeysInspections := []int{}
	for _, monkey := range monkeys {
		monkeysInspections = append(monkeysInspections, -monkey.inspections)
	}
	sort.Ints(monkeysInspections[:])
	monkeyBusiness := monkeysInspections[0] * monkeysInspections[1]

	return monkeyBusiness
}

func DecreaseWorryDivide(item int, monkeys []Monkey) int {
	return int(item / 3)
}

func DecreaseWorryLCM(item int, monkeys []Monkey) int {
	LCM := 1
	for _, monkey := range monkeys {
		LCM *= monkey.divisible_test
	}
	return int(item % LCM)
}

func main() {
	
	fmt.Println("Advent-Of-Code 2022 - Day10")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Solve(Parse("tests.txt"),  20,    DecreaseWorryDivide, false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Solve(Parse("inputs.txt"), 20,    DecreaseWorryDivide, false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Solve(Parse("tests.txt"),  10000, DecreaseWorryLCM,    false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Solve(Parse("inputs.txt"), 10000, DecreaseWorryLCM,    false))

}
