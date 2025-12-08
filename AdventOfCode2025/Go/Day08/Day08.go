package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x       int
	y       int
	z       int
	curcuit int
	name    int
}

func Parse(filename string) map[int]Point {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	points := map[int]Point{}

	lineNumber := 0
	for fscanner.Scan() {

		line := fscanner.Text()
		nums := []int{}
		for _, char := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(char)
			nums = append(nums, num)
		}

		point := Point{}

		point.x = nums[0]
		point.y = nums[1]
		point.z = nums[2]
		point.curcuit = 0
		point.name = lineNumber
		points[lineNumber] = point
		lineNumber++
	}

	return points
}

func EuclideanDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(float64(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2)))
}

func Part1(filename string, numIterations int, debug bool) int {
	points := Parse(filename)

	connections := 0
	iterations := 0

	distances := map[string]int{}
	// start chain

	for _, p1 := range points {
		for _, p2 := range points {
			if p1.name == p2.name {
				continue
			}
			distance := int(EuclideanDistance(p1, p2))
			if p1.name > p2.name {
				distances[fmt.Sprintf("%d-%d", p1.name, p2.name)] = distance
			} else {
				distances[fmt.Sprintf("%d-%d", p2.name, p1.name)] = distance
			}
		}
	}

	// sort distances lowest to highest
	distancesSorted := []string{}

	for d := range distances {
		distancesSorted = append(distancesSorted, d)
	}

	sort.Slice(distancesSorted, func(i, j int) bool {
		return distances[distancesSorted[i]] < distances[distancesSorted[j]]
	})

	for _, minDistanceKey := range distancesSorted {

		current_number_of_curcuits := -1
		highest_current_curcuit := 0
		for i := 0; i < len(points); i++ {
			point := points[i]
			if point.curcuit > 0 {
				if point.curcuit > highest_current_curcuit {
					highest_current_curcuit = point.curcuit
				}
			}
			if point.curcuit > current_number_of_curcuits {
				current_number_of_curcuits = point.curcuit
			}
		}

		if debug {
			fmt.Println("Current number of connections:", connections, "Current number of curcuits:", current_number_of_curcuits)
		}

		if iterations == numIterations {
			break
		}

		pointsInChainStr := strings.Split(minDistanceKey, "-")
		p1Index, _ := strconv.Atoi(pointsInChainStr[0])
		p2Index, _ := strconv.Atoi(pointsInChainStr[1])

		point1New := points[p1Index]
		point2New := points[p2Index]

		// p1 has a curcuit, p2 does not
		if point1New.curcuit > 0 && point2New.curcuit == 0 {
			point2New.curcuit = point1New.curcuit
			points[p1Index] = point1New
			points[p2Index] = point2New
			connections++
			iterations++
			continue
		}

		// p2 has a curcuit, p1 does not
		if point2New.curcuit > 0 && point1New.curcuit == 0 {
			point1New.curcuit = point2New.curcuit
			points[p1Index] = point1New
			points[p2Index] = point2New
			connections++
			iterations++
			continue
		}

		// neither have curcuits
		if point2New.curcuit == 0 && point1New.curcuit == 0 {
			new_current_num := highest_current_curcuit + 1
			point1New.curcuit = new_current_num
			point2New.curcuit = new_current_num
			points[p1Index] = point1New
			points[p2Index] = point2New
			connections++
			iterations++
			continue
		}

		// both have curcuits
		if point2New.curcuit > 0 && point1New.curcuit > 0 {
			if point1New.curcuit == point2New.curcuit {
				// Count the iteration but not the connection
				iterations++
				continue
			}
			// merge curcuits and use lower number
			if point1New.curcuit < point2New.curcuit {
				for i := 0; i < len(points); i++ {
					p := points[i]
					if p.curcuit == point2New.curcuit {
						p.curcuit = point1New.curcuit
						points[i] = p
					}
				}
				point2New.curcuit = point1New.curcuit
			} else if point2New.curcuit < point1New.curcuit {
				for i := 0; i < len(points); i++ {
					p := points[i]
					if p.curcuit == point1New.curcuit {
						p.curcuit = point2New.curcuit
						points[i] = p
					}
				}
				point1New.curcuit = point2New.curcuit
			}
			connections++
			iterations++
			continue
		}

	}

	// number of points in each curcuit
	curcuits := map[int]int{}
	for i := 0; i < len(points); i++ {
		point := points[i]
		if point.curcuit > 0 {
			curcuits[point.curcuit]++
		}
	}

	if debug {
		fmt.Println("Curcuits found:", curcuits)
	}

	max1 := 0
	max2 := 0
	max3 := 0
	for _, v := range curcuits {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}

	ans := max1 * max2 * max3
	return ans
}

func Part2(filename string, debug bool) int {
	points := Parse(filename)
	connections := 0
	iterations := 0

	distances := map[string]int{}
	// start chain

	// find distances between all points
	for _, p1 := range points {
		for _, p2 := range points {
			if p1.name == p2.name {
				continue
			}
			distance := int(EuclideanDistance(p1, p2))
			if p1.name > p2.name {
				distances[fmt.Sprintf("%d-%d", p1.name, p2.name)] = distance
			} else {
				distances[fmt.Sprintf("%d-%d", p2.name, p1.name)] = distance
			}
		}
	}

	// sort distances lowest to highest
	distancesSorted := []string{}

	for d := range distances {
		distancesSorted = append(distancesSorted, d)
	}

	sort.Slice(distancesSorted, func(i, j int) bool {
		return distances[distancesSorted[i]] < distances[distancesSorted[j]]
	})

	finalConnection := ""

	for n, minDistanceKey := range distancesSorted {

		current_number_of_curcuits := -1
		highest_current_curcuit := 0
		for i := 0; i < len(points); i++ {
			point := points[i]
			if point.curcuit > 0 {
				if point.curcuit > highest_current_curcuit {
					highest_current_curcuit = point.curcuit
				}
			}
			if point.curcuit > current_number_of_curcuits {
				current_number_of_curcuits = point.curcuit
			}
		}

		if debug {
			fmt.Println("Current number of connections:", connections, "Current number of curcuits:", current_number_of_curcuits)
		}

		number_of_points_in_first_curcuit := 0
		for i := 0; i < len(points); i++ {
			point := points[i]
			if point.curcuit == 1 {
				number_of_points_in_first_curcuit++
			}
		}

		if number_of_points_in_first_curcuit == len(points) {
			finalConnection = distancesSorted[n-1]
			break
		}

		pointsInChainStr := strings.Split(minDistanceKey, "-")
		p1Index, _ := strconv.Atoi(pointsInChainStr[0])
		p2Index, _ := strconv.Atoi(pointsInChainStr[1])

		point1New := points[p1Index]
		point2New := points[p2Index]

		// p1 has a curcuit, p2 does not
		if point1New.curcuit > 0 && point2New.curcuit == 0 {
			point2New.curcuit = point1New.curcuit
			points[p1Index] = point1New
			points[p2Index] = point2New
			connections++
			iterations++
			continue
		}

		// p2 has a curcuit, p1 does not
		if point2New.curcuit > 0 && point1New.curcuit == 0 {
			point1New.curcuit = point2New.curcuit
			points[p1Index] = point1New
			points[p2Index] = point2New
			connections++
			iterations++
			continue
		}

		// neither have curcuits
		if point2New.curcuit == 0 && point1New.curcuit == 0 {

			new_current_num := highest_current_curcuit + 1
			point1New.curcuit = new_current_num
			point2New.curcuit = new_current_num
			points[p1Index] = point1New
			points[p2Index] = point2New
			connections++
			iterations++
			continue
		}

		// both have curcuits
		if point2New.curcuit > 0 && point1New.curcuit > 0 {
			if point1New.curcuit == point2New.curcuit {
				// Count the iteration but not the connection
				iterations++
				continue
			}
			// merge curcuits and use lower number
			if point1New.curcuit < point2New.curcuit {
				for i := 0; i < len(points); i++ {
					p := points[i]
					if p.curcuit == point2New.curcuit {
						p.curcuit = point1New.curcuit
						points[i] = p
					}
				}
				point2New.curcuit = point1New.curcuit
			} else if point2New.curcuit < point1New.curcuit {
				for i := 0; i < len(points); i++ {
					p := points[i]
					if p.curcuit == point1New.curcuit {
						p.curcuit = point2New.curcuit
						points[i] = p
					}
				}
				point1New.curcuit = point2New.curcuit
			}
			connections++
			iterations++
			continue
		}

	}

	ansPointsX := strings.Split(finalConnection, "-")
	p1Index, _ := strconv.Atoi(ansPointsX[0])
	p2Index, _ := strconv.Atoi(ansPointsX[1])

	point1Final := points[p1Index]
	point2Final := points[p2Index]

	ans := point1Final.x * point2Final.x
	return ans
}

func main() {
	// the code is a mess but it works :) - to tired to clean it up
	fmt.Println("Advent-Of-Code 2025 - Day08")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", 10, false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", 1000, false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
