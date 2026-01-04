package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Device struct {
	name    string
	outputs []string
}

func Parse(filename string) map[string]Device {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	devices := map[string]Device{}

	for fscanner.Scan() {

		line := fscanner.Text()

		name := strings.Split(line, ": ")[0]
		outputs := strings.Split(strings.Split(line, ": ")[1], " ")

		device := Device{
			name:    name,
			outputs: outputs,
		}

		devices[name] = device
	}

	return devices
}

func findAllRoutes(startDevice Device, endDeviceName string, Devices map[string]Device, currentPath string, visited map[string]bool, depth int, filterOut string) []string {
	// Prevent infinite recursion with depth limit
	if depth > 100 {
		return []string{}
	}

	currentPath += "->" + startDevice.name

	// Check if we've already visited this device in the current path (cycle detection)
	if visited[startDevice.name] {
		return []string{}
	}

	// Mark this device as visited
	visited[startDevice.name] = true

	// Collect all paths from ALL outputs (not just first one)
	allPaths := []string{}
	for _, output := range startDevice.outputs {
		// Check if the device exists in the map
		if output == endDeviceName {
			// Found a complete path to "out"
			if filterOut != "" && strings.Contains(currentPath, filterOut) {
				continue
			}
			allPaths = append(allPaths, currentPath)
			continue
		}

		// Create a copy of visited map for this branch
		visitedCopy := make(map[string]bool)
		for k, v := range visited {
			visitedCopy[k] = v
		}

		paths := findAllRoutes(Devices[output], endDeviceName, Devices, currentPath, visitedCopy, depth+1, "")
		allPaths = append(allPaths, paths...)
	}

	return allPaths
}

func findAllRoutesVia(device Device, endDeviceName string, Devices map[string]Device, via1 string, via2 string) int {

	// took a hunch that this is the intended approach for part 2

	StoV1 := findAllRoutes(device, via1, Devices, "", make(map[string]bool), 0, via2)
	fmt.Println("Paths from Start to", via1, ":", len(StoV1))
	StoV2 := findAllRoutes(device, via2, Devices, "", make(map[string]bool), 0, via1)
	fmt.Println("Paths from Start to", via2, ":", len(StoV2))
	V1toE := findAllRoutes(Devices[via1], endDeviceName, Devices, "", make(map[string]bool), 0, via2)
	fmt.Println("Paths from", via1, "to End:", len(V1toE))
	V2toE := findAllRoutes(Devices[via2], endDeviceName, Devices, "", make(map[string]bool), 0, via1)
	fmt.Println("Paths from", via2, "to End:", len(V2toE))
	V1toV2 := findAllRoutes(Devices[via1], via2, Devices, "", make(map[string]bool), 0, "")
	fmt.Println("Paths from", via1, "to", via2, ":", len(V1toV2))
	V2toV1 := findAllRoutes(Devices[via2], via1, Devices, "", make(map[string]bool), 0, "")
	fmt.Println("Paths from", via2, "to", via1, ":", len(V2toV1))

	fmt.Println("Paths from Start to", via1, ":", len(StoV1))
	fmt.Println("Paths from Start to", via2, ":", len(StoV2))
	fmt.Println("Paths from", via1, "to End:", len(V1toE))
	fmt.Println("Paths from", via2, "to End:", len(V2toE))
	fmt.Println("Paths from", via1, "to", via2, ":", len(V1toV2))
	fmt.Println("Paths from", via2, "to", via1, ":", len(V2toV1))

	return (len(StoV1) * len(V1toV2) * len(V2toE)) + (len(StoV2) * len(V2toV1) * len(V1toE))
}

func Part1(filename string, debug bool) int {
	devices := Parse(filename)

	paths := findAllRoutes(devices["you"], "out", devices, "", make(map[string]bool), 0, "")

	if debug {
		fmt.Println("\nAll paths found:")
		for i, path := range paths {
			fmt.Printf("Path %d: %s\n", i+1, path)
		}
	}

	ans := len(paths)
	return ans
}

func Part2(filename string, debug bool) int {
	devices := Parse(filename)

	ans := findAllRoutesVia(devices["svr"], "out", devices, "dac", "fft")

	// fmt.Println("Finding all paths (this may take a while)...")
	// paths := findAllRoutes(devices["svr"], "out", devices, "", make(map[string]bool), 0)
	// fmt.Printf("Found %d total paths\n", len(paths))

	// // might need to identify loops

	// ans := 0
	// if debug {
	// 	fmt.Println("\nAll paths found:")
	// 	for i, path := range paths {
	// 		fmt.Printf("Path %d: %s\n", i+1, path)
	// 		if strings.Contains(path, "dac") && strings.Contains(path, "fft") {
	// 			ans++
	// 		}
	// 	}
	// } else {
	// 	// Just count without printing
	// 	for _, path := range paths {
	// 		if strings.Contains(path, "dac") && strings.Contains(path, "fft") {
	// 			ans++
	// 		}
	// 	}
	// }

	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day11")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests2.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
