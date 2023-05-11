package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	isFile   bool
	children map[string]*node
	parent   *node
}

func Parse(filename string) []map[string]string {

	// read file into string
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")

	history := []map[string]string{}

	currentCommand := map[string]string{
		"command": "",
		"output":  "",
	}

	for _, line := range lines {
		if string(line[0]) == "$" {
			if currentCommand["command"] == "" {
				currentCommand["command"] = line
			} else {
				history = append(history, currentCommand)
				currentCommand = map[string]string{"command": line, "output": ""}
			}
		} else {
			currentCommand["output"] += (line + ",")
		}
	}
	history = append(history, currentCommand)

	return history
}

func calcSize(root node) (size int) {

	if root.isFile {
		return root.size

		// if dir sum file contents
	} else {
		for _, child := range root.children {
			size += calcSize(*child)
		}
	}
	return size
}

func Part1(filename string) int {

	history := Parse(filename)

	var cwd *node
	nodes := []*node{}

	for _, entry := range history {

		// cd commands
		if string(entry["command"][:4]) == "$ cd" {
			cdLocation := string(entry["command"][5:])

			// cd to root
			if cdLocation == "/" {
				cwd = &node{"/", 0, false, make(map[string]*node), nil}

				// cd to up a dir
			} else if cdLocation == ".." {
				cwd = cwd.parent

				// cd to a dir
			} else {
				cwd = cwd.children[cdLocation]
			}
		}

		// ls commands
		if string(entry["command"][:4]) == "$ ls" {

			for _, result := range strings.Split(entry["output"], ",") {

				// if dir contents not empty
				if result != "" {

					// nested dirs
					if string(result[:3]) == "dir" {
						dir := strings.Split(result, " ")[1]
						cwd.children[dir] = &node{dir, 0, false, make(map[string]*node), cwd}
						nodes = append(nodes, cwd.children[dir])

						// nested files
					} else {
						filesize, file := result[:strings.Index(result, " ")], result[strings.Index(result, " ")+1:]
						cwd.children[file] = &node{file, 0, true, make(map[string]*node), cwd}
						cwd.children[file].size, _ = strconv.Atoi(filesize)
					}
				}
			}
		}
	}

	var totalSize int

	for _, node := range nodes {

		size := calcSize(*node)

		// if size less than 100000 add to total
		if size <= 100000 {
			totalSize += size
		}
	}

	return totalSize
}

func Part2(filename string) int {

	history := Parse(filename)

	var cwd *node
	nodes := []*node{}

	for _, entry := range history {

		// cd commands
		if string(entry["command"][:4]) == "$ cd" {
			cdLocation := string(entry["command"][5:])

			// cd to root
			if cdLocation == "/" {
				cwd = &node{"/", 0, false, make(map[string]*node), nil}
				nodes = append(nodes, cwd)

				// cd to up a dir
			} else if cdLocation == ".." {
				cwd = cwd.parent

				// cd to a dir
			} else {
				cwd = cwd.children[cdLocation]
			}
		}

		// ls commands
		if string(entry["command"][:4]) == "$ ls" {

			for _, result := range strings.Split(entry["output"], ",") {

				// if dir contents not empty
				if result != "" {

					// nested dirs
					if string(result[:3]) == "dir" {
						dir := strings.Split(result, " ")[1]
						cwd.children[dir] = &node{dir, 0, false, make(map[string]*node), cwd}
						nodes = append(nodes, cwd.children[dir])

						// nested files
					} else {
						filesize, file := result[:strings.Index(result, " ")], result[strings.Index(result, " ")+1:]
						cwd.children[file] = &node{file, 0, true, make(map[string]*node), cwd}
						cwd.children[file].size, _ = strconv.Atoi(filesize)
					}
				}
			}
		}
	}

	var totalSize int

	for _, node := range nodes {

		if node.name == "/" {
			totalSize = calcSize(*node)
		}
	}

	// dumb var for comparison
	answer := 100000000000000000

	for _, node := range nodes {

		size := calcSize(*node)

		if size < answer && 70000000-(totalSize-size) >= 30000000 {
			answer = size
		}
	}

	return answer
}

func main() {

	testfile := "tests.txt"
	inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day07")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1(inputfile))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2(testfile))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2(inputfile))
}
