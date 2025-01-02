package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) (map[string]int, []int) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	registers := map[string]int{}
	program := []int{}

	for fscanner.Scan() {
		line := fscanner.Text()

		if line == "" {
			continue
		}
		if line[:8] == "Register" {
			n, _ := strconv.Atoi(string(line[12:]))
			registers[string(line[9])] = n
		}
		if line[:7] == "Program" {
			for _, r := range strings.Split(line[9:], ",") {
				n, _ := strconv.Atoi(r)
				program = append(program, n)
			}
		}
	}

	return registers, program
}

func combooperand(operand int, registers map[string]int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers["A"]
	case 5:
		return registers["B"]
	case 6:
		return registers["C"]
	case 7:
		fmt.Println("Error: combooperand(operand, registers) = 7")
		os.Exit(0)
	}
	fmt.Println("Error: combooperand(operand, registers) = 0")
	return 0
}
func dv(registers map[string]int, operand int) int {
	return int((registers["A"]) / (1 << combooperand(operand, registers)))
}

// 0
func adv(registers map[string]int, operand int) map[string]int {
	registers["A"] = dv(registers, operand)
	return registers
}

// 1
func bxl(registers map[string]int, operand int) map[string]int {
	registers["B"] = registers["B"] ^ operand
	return registers
}

// 2
func bst(registers map[string]int, operand int) map[string]int {
	registers["B"] = combooperand(operand, registers) % 8
	return registers
}

// 3
func jnz(registers map[string]int, operand, instructionPointer int) int {
	if registers["A"] != 0 {
		instructionPointer := operand
		return instructionPointer - 2
	}
	return instructionPointer
}

// 4
func bxc(registers map[string]int, operand int) map[string]int {
	registers["B"] = registers["B"] ^ registers["C"]
	return registers
}

// 5
func out(registers map[string]int, operand int) string {
	return fmt.Sprintf("%d,", (combooperand(operand, registers) % 8))
}

// 6
func bdv(registers map[string]int, operand int) map[string]int {
	registers["B"] = dv(registers, operand)
	return registers
}

// 7
func cdv(registers map[string]int, operand int) map[string]int {
	registers["C"] = dv(registers, operand)
	return registers
}

func run(registers map[string]int, program []int, debug bool) string {
	outputs := ""

	for i := 0; i < len(program)-1; i += 2 {

		opcode := program[i]
		operand := program[i+1]

		if debug {
			fmt.Println("b", opcode, operand, registers)
		}
		switch opcode {
		case 0:
			registers = adv(registers, operand)
			break
		case 1:
			registers = bxl(registers, operand)
			break
		case 2:
			registers = bst(registers, operand)
			break
		case 3:
			i = jnz(registers, operand, i)
			break
		case 4:
			registers = bxc(registers, operand)
			break
		case 5:
			outputs += out(registers, operand)
			break
		case 6:
			registers = bdv(registers, operand)
			break
		case 7:
			registers = cdv(registers, operand)
			break
		}
		if debug {
			fmt.Println("a", opcode, operand, registers)
		}
	}

	return outputs[:len(outputs)-1]
}

func Part1(filename string, debug bool) string {
	registers, program := parse(filename)

	if debug {
		fmt.Println(registers, program)
	}

	ans := run(registers, program, debug)

	return ans
}

func Part2(filename string, debug bool) int {

	// TODO fix part 2

	registers, program := parse(filename)

	programString := ""
	for _, s := range program {
		programString += fmt.Sprintf("%d,", s)
	}
	programString = programString[:len(programString)-1]

	if debug {
		fmt.Println(registers, program)
	}

	// 117440 = ( (58* 2024) +48)

	// thanks dan for the optimizations
	outputs := ""
	ans := registers["A"]
	for len(outputs) < len(programString) {
		outputs = run(map[string]int{"A": ans, "B": 0, "C": 0}, program, debug)
		ans += registers["A"]
	}

	// ans := registers["A"]

	// outputs := run(map[string]int{"A": ans, "B": 0, "C": 0}, program, debug)

	for {
		outputs := run(map[string]int{"A": ans, "B": 0, "C": 0}, program, debug)
		if ans == 0 {
			break
		}
		fmt.Println(ans, "=", int(ans/registers["A"]), "*", registers["A"], "+", ans%registers["A"], "want :", programString, "got :", outputs)
		correctAns := true
		if len(outputs) > len(programString) {
			fmt.Println("Error: Part2", len(outputs) ,len(programString))
			fmt.Println(ans)
			ans /= 8
			continue
		}
		x := 1
		for i := len(programString) - 1; i >= 0; i-=2 {
			if outputs[i] != programString[i] {
				correctAns = false
				// to increment each digit of the outpputed number
				// number = + 8**didgiti index
				// ans += (i << 3)
				fmt.Println("add", int(float64((-outputs[i]+programString[i]))),int(float64((-outputs[i]+programString[i]))),int(math.Pow(8, float64(i/2))),int(float64((-outputs[i]+programString[i])))*int(math.Pow(8, float64(i/2))), ans, ans+int(float64((-outputs[i]+programString[i])))*int(math.Pow(8, float64(i/2))))
				// os.Exit(0)
				// a,_:=strconv.Atoi(string(outputs[i]))
				// b,_:=strconv.Atoi(string(programString[i]))
				ans += int(float64((-outputs[i]+programString[i])))*int(math.Pow(8, float64(i/2))) //*(b-a)
				fmt.Println(ans)
				outputs = run(map[string]int{"A": ans, "B": 0, "C": 0}, program, debug)
				break
			}
			x++
		}
		if correctAns && len(outputs) == len(programString) {
			break
		}
	}

	// // 190593311063055 is valid but too high
	// // 190593313094671

	// r := run(map[string]int{"A": ans, "B": 0, "C": 0}, program, debug)
	// fmt.Println("", r, "\n", programString)
	// if r != programString {
	// 	fmt.Println("Error: Part2")
	// 	os.Exit(0)
	// }

	//(8 * 8 * 8 * 8) - 2048 = 2048
	// 8 * (0*( 8 ^0)+ 3*(8^1) + 5*(8^2) + 4*(8^3)+ 3 *(8^4) + 0*(8^5)) = 117440
	// 8 * (2*( 8 ^0)+ 4*( 8 ^1)+ 1*( 8 ^2)+ 2*( 8 ^3)+ 7*( 8 ^4)+ 5*( 8 ^5)+ 1*( 8 ^6)+ 7*( 8 ^7)+ 4*( 8 ^8)+ 4*( 8 ^9)+ 0*( 8 ^10)+ 3*( 8 ^11)+ 5*( 8 ^12)+ 5*( 8 ^13)+ 3*( 8 ^14)+ 0*( 8 ^15)) = 130503239246608
	// 2,4,1,2,7,5,1,7,4,4,0,3,5,5,3,0

	// fmt.Println(run(map[string]int{"A": ((2024 * 58) + 48), "B": 0, "C": 0}, program, debug))
	// fmt.Println(run(map[string]int{"A": (2024 * 56), "B": 0, "C": 0}, program, debug))
	// fmt.Println(run(map[string]int{"A": ((2024 * 56) + (8)), "B": 0, "C": 0}, program, debug))
	// fmt.Println(run(map[string]int{"A": ((2024 * 56) + (8 * 8)), "B": 0, "C": 0}, program, debug))
	// fmt.Println(run(map[string]int{"A": 8 * (0*int(math.Pow(8,0))+ 3*int(math.Pow(8,1)) + 5*int(math.Pow(8,2)) + 4*int(math.Pow(8,3))+ 3 *int(math.Pow(8,4)) + 0*int(math.Pow(8,5))), "B": 0, "C": 0}, program, debug))

	// fmt.Println(run(map[string]int{"A": 8 * (0*int(math.Pow(8,0))+ 3*int(math.Pow(8,1)) + 5*int(math.Pow(8,2)) + 4*int(math.Pow(8,3))+ 3 *int(math.Pow(8,4)) + 0*int(math.Pow(8,5))), "B": 0, "C": 0}, program, debug))
 	x := 8*(2*int(math.Pow( 8 ,0))+ 4*int(math.Pow( 8 ,1))+ 1*int(math.Pow( 8 ,2))+ 2*int(math.Pow( 8 ,3))+ 7*int(math.Pow( 8 ,4))+ 5*int(math.Pow( 8 ,5))+ 1*int(math.Pow( 8 ,6))+ 7*int(math.Pow( 8 ,7))+ 4*int(math.Pow( 8 ,8))+ 4*int(math.Pow( 8 ,9))+ 0*int(math.Pow( 8 ,10))+ 3*int(math.Pow( 8 ,11))+ 3*int(math.Pow( 8 ,12))+ 3*int(math.Pow( 8 ,13))+ 0*int(math.Pow( 8 ,14))+ 0*int(math.Pow(8 ,15)))
	fmt.Println(run(map[string]int{"A": x, "B": 0, "C": 0}, program, debug))

	// for i := 0; i < 35000; i++ {
	// 	o := run(map[string]int{"A": 130503239247*i/8, "B": 0, "C": 0}, program, debug)
	// 	fmt.Println(i, o, programString)
	// 	if o == programString {
	// 		fmt.Println("Found", i)
	// 		break
	// 	}
	// // }
	// fmt.Println(run(map[string]int{"A": 2018*130503239247, "B": 0, "C": 0}, program, debug))
	// // fmt.Println(run(map[string]int{"A": (117440), "B": 0, "C": 0}, program, debug))
	// // fmt.Println(run(map[string]int{"A": (130503239246608), "B": 0, "C": 0}, program, debug))
	// fmt.Println(run(map[string]int{"A": (190593311063055), "B": 0, "C": 0}, program, debug))
	// fmt.Println(run(map[string]int{"A": (190593313094671), "B": 0, "C": 0}, program, debug))
	// fmt.Println(programString)
	// fmt.Println((190593313094671 - 190593311063055))
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day17")
	fmt.Println()
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
