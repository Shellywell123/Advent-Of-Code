package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	left string
	right string
}

func Parse(filename string) (map[string]node, string) {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	nodes := map[string]node{}
	instructions := ""

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		if y == 0 {
			instructions = line
			y++
			continue
		}
		if line != "" {
			newNode := node{}
			newNode.left = string(line[7:10])
			newNode.right = string(line[12:15])
			nodes[string(line[0:3])] = newNode
			y++
		}
	}

	return nodes, instructions
}

func Part1(filename string) int {

	nodes, instructions := Parse(filename)
	currentNode := "AAA"
	instruction := ""

	steps := 0
	for currentNode != "ZZZ" {
		for i, _ := range instructions {
			instruction = string(instructions[i])

			if instruction == "L" {
				currentNode = nodes[currentNode].left
				steps++
			}
			if instruction == "R" {
				currentNode = nodes[currentNode].right
				steps++
			}
		}
	}
	return steps
}

func gcd(first, second int) int {
	if first == 0 {
		return second
	}
	return gcd(second % first, first)
}
 
func lcmm(args []int) int {
    if(len(args) == 2){
        return args[0] *  args[1] / gcd(args[0],args[1])
    } else {
        var arg0 = args[0];
        argsShort := args[1:]
		args = argsShort
        return arg0 *  lcmm(args) / gcd(arg0,lcmm(args))
    }
}

func Part2(filename string) int {
	nodes, instructions := Parse(filename)

	currentNodes := []string{}
	for k := range nodes {
		if string(k[2]) == "A" {
			currentNodes = append(currentNodes, k)
		}
	}

	instruction := ""
	steps := 0
	list :=[]int{}
	for len(list) != len(currentNodes) {
		for i, _ := range instructions {
			
			for n, _ := range currentNodes {
				
				if string(currentNodes[n][2]) == "Z" {
					list = append(list, steps)
				}
				instruction = string(instructions[i])

				if instruction == "L" {
					currentNodes[n] = nodes[currentNodes[n]].left
				}
				if instruction == "R" {
					currentNodes[n] = nodes[currentNodes[n]].right
				}
			}
			steps++
		}
	}
	return lcmm(list)
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day08")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests2.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
