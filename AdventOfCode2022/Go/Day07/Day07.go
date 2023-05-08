package main

import (
	"fmt"
	"os"
	"strings"
)

func Parse(filename string) []map[string]string {

	// read file into string
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(b), "\n")

	history := []map[string]string{}

	currentCommand := map[string]string{
		"command":"",
		"output": "",
	}

	for _, line := range lines {
		if string(line[0]) == "$" {
			if currentCommand["command"] == "" {
				currentCommand["command"] = line
			} else {
				history = append(history, currentCommand)
				currentCommand = map[string]string{ "command": line, "output": ""}
			}
		} else {
			currentCommand["output"] += (line + ",")
		}
	}
	history = append(history, currentCommand)

	return history
}

func Part1(filename string, n int) int {

	history := Parse(filename)


	tree := map[string]any{}
	cwd := ""

	for _, entry := range history {

		// cd commands
		if string(entry["command"][:4]) == "$ cd" {
			cdLocation := string(entry["command"][5:])

			if cdLocation == "/" {
				cwd = "/"
				tree["/"] = map[string]any{}
			} else if cdLocation == ".."{
				newCwd := ""
				cwdDirs := strings.Split(cwd, "/")
                for _, dir := range cwdDirs[:len(cwdDirs)-2] {
					newCwd += dir + "/"
				}
                cwd = newCwd
			} else {
				cwd = cdLocation + "/"
			}
		}

		// list commands
		if string(entry["command"][:4]) == "$ ls" {

			for _, result := range strings.Split(",") {

				if result != nil {

					// dirs
					if string(result[:3]) == "dir" {
						dir := strings.Split(result, " ")[1]
						tree[] = map[string]any{dir: {}}
					}
				}
				
			}

			for result in output.split(','):
                if result:
                    # dirs
                    if result[:3] == 'dir':
                        dir = result.split(' ')[1]
                        tree_str = '["/"]'
                        for dirp in (cwd.split('/')[1:-1]):
                            tree_str += f'["{dirp}"]'
                        eval(f'tree{tree_str}.update('+'{"'+dir+'":{}})')

                    # files
                    else:
                        filesize,file = result.split(' ')[0], result.split(' ')[1]
                        tree_str = '["/"]'
                        for dirp in (cwd.split('/')[1:-1]):
                            tree_str += f'["{dirp}"]'
                        eval(f'tree{tree_str}.update('+'{"'+file+'":'+filesize+'})')
		}


	}

	return 0
}

func main() {

	testfile := "tests.txt"
	//inputfile := "inputs.txt"

	fmt.Println("Advent-Of-Code 2022 - Day06")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1(testfile, 4))
	// fmt.Printf("Inputs: Answer to Part 1 = %i\n", Solve(inputfile,4))
	// fmt.Printf("Tests : Answer to Part 2 = %i\n", Solve(testfile,14))
	// fmt.Printf("Inputs: Answer to Part 2 = %i\n", Solve(inputfile,14))
}
