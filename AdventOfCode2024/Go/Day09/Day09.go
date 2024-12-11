package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func parse(filename string) []int {

	contentBytes, _ := os.ReadFile(filename)
	content := string(contentBytes)

	fileSystem := []int{}
	for i, c := range content {
		n, _ := strconv.Atoi(string(c))

		// if even
		if i%2 == 0 {
			id := int(i+1) / 2
			for x := 0; x < n; x++ {
				fileSystem = append(fileSystem, id)
			}
			// if odd
		} else {
			for x := 0; x < n; x++ {
				fileSystem = append(fileSystem, -1)
			}
		}
	}

	return fileSystem
}

func reorderBlocks(blocks []int) []int {

	spaceToFill := true
	for spaceToFill {
		for i := 0; i < len(blocks)-1; i++ {
			// find next free space
			if blocks[i] == -1 {
				// find last block
				for j := len(blocks) - 1; j >= 0; j-- {

					if j <= i {
						spaceToFill = false
						break
					}

					// swap if block found after space
					if blocks[j] != -1 {
						blocks = slices.Concat(blocks[:i], []int{blocks[j]}, blocks[i+1:j], []int{blocks[i]}, blocks[j+1:])
						break
					}
				}
			}
		}
	}

	return blocks
}

func chunkFiles(blocks []int) [][]int {

	files := [][]int{}
	currentFile := []int{}
	prevBlock := blocks[0]
	for i := 0; i < len(blocks); i++ {

		// last chunk
		if i == len(blocks)-1 {
			currentFile = append(currentFile, blocks[i])
			files = append(files, currentFile)
			break
		}
		// new chunk
		if prevBlock != blocks[i] {
			if len(currentFile) > 0 {
				files = append(files, currentFile)
				currentFile = []int{blocks[i]}
			}
			// same chunk
		} else {
			currentFile = append(currentFile, blocks[i])
		}

		prevBlock = blocks[i]
	}

	return files
}

func reorderFiles(files [][]int, jStart int) []int {

	if jStart == -1 {
		jStart = len(files) - 1
	}

	for j := jStart; j >= 0; j-- {

		for i := 0; i < len(files); i++ {

			// find next free space
			if files[i][0] == -1 {

				// // guard clause
				if j < i {
					break
				}

				// find next block
				if files[j][0] != -1 {

					if i == j {
						return slices.Concat(files...)
					}

					diff := len(files[i]) - len(files[j])

					// doesn't fit
					if diff < 0 {
						continue
					}

					// partial/full fit
					if diff >= 0 {

						filledSpace := files[j]
						unfilledSpace := []int{}
						for g := 0; g < diff; g++ {
							unfilledSpace = append(unfilledSpace, -1)
						}

						space := []int{}
						for g := 0; g < len(files[i])-diff; g++ {
							space = append(space, -1)
						}

						if len(unfilledSpace) == 0 {
							files = slices.Concat(files[:i], [][]int{filledSpace}, files[i+1:j], [][]int{space}, files[j+1:])
						} else {
							files = slices.Concat(files[:i], [][]int{filledSpace}, [][]int{unfilledSpace}, files[i+1:j], [][]int{space}, files[j+1:])
						}
						return reorderFiles(files, j)
					}
				}
			}
		}
	}
	return slices.Concat(files...)
}

func checksum(blocks []int) int {
	checksum := 0
	for i, c := range blocks {
		if c == -1 {
			continue
		}
		checksum += (i * c)
	}

	return checksum
}

func Part1(filename string) int {
	fileSystem := parse(filename)
	blocks := reorderBlocks(fileSystem)
	ans := checksum(blocks)
	return ans
}

func Part2(filename string) int {
	fileSystem := parse(filename)
	files := chunkFiles(fileSystem)
	blocks := reorderFiles(files, -1)
	ans := checksum(blocks)
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2024 - Day09")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
