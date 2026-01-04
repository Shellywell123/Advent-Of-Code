package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Machine struct {
	intialIndicators  string
	desiredIndicators string
	lowestButtonInput []int
	buttons           [][]int
	initalJoltages    []int
	desiredJoltages   []int
}

func Parse(filename string) map[int]Machine {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	machines := map[int]Machine{}

	lineNumber := 0
	for fscanner.Scan() {

		line := fscanner.Text()

		desiredIndicators := strings.Split(line, "]")[0][1:]

		initialIndicators := ""
		for i := 0; i < len(desiredIndicators); i++ {
			initialIndicators += "."
		}

		buttons := [][]int{}
		buttonsStringParts := strings.Split(strings.Split(strings.Split(line, "] (")[1], ") {")[0], ") (")
		for _, buttonString := range buttonsStringParts {
			buttonNums := []int{}
			buttonParts := strings.Split(buttonString, ",")
			for _, buttonPart := range buttonParts {
				num, _ := strconv.Atoi(buttonPart)
				buttonNums = append(buttonNums, num)
			}
			buttons = append(buttons, buttonNums)
		}

		initalJoltages := []int{}
		desiredJoltages := []int{}
		for _, char := range strings.Split(strings.Split(strings.Split(line, "{")[1], "}")[0], ",") {
			num, _ := strconv.Atoi(char)
			desiredJoltages = append(desiredJoltages, num)
			initalJoltages = append(initalJoltages, 0)
		}

		machines[lineNumber] = Machine{
			intialIndicators:  initialIndicators,
			desiredIndicators: desiredIndicators,
			lowestButtonInput: []int{},
			buttons:           buttons,
			initalJoltages:    initalJoltages,
			desiredJoltages:   desiredJoltages,
		}
		lineNumber++
	}

	return machines
}

func switchIndicator(currentIndicator string, indicatorToSwitch int) string {
	currentPosition := currentIndicator[indicatorToSwitch]
	if currentPosition == '.' {
		return currentIndicator[:indicatorToSwitch] + "#" + currentIndicator[indicatorToSwitch+1:]
	} else {
		return currentIndicator[:indicatorToSwitch] + "." + currentIndicator[indicatorToSwitch+1:]
	}
}

func increaseJoltage(currentJoltages []int, joltageToIncrease int) []int {
	currentJoltages[joltageToIncrease] += 1
	return currentJoltages
}

func pressButtonWiring(machine Machine, buttonIndex int, currentIndicators string) string {

	for _, indicatorToSwitch := range machine.buttons[buttonIndex] {
		currentIndicators = switchIndicator(currentIndicators, indicatorToSwitch)
	}
	return currentIndicators
}

func pressButtonJoltage(machine Machine, buttonIndex int, currentJoltages []int) []int {

	for _, joltagesToIncrease := range machine.buttons[buttonIndex] {
		currentJoltages = increaseJoltage(currentJoltages, joltagesToIncrease)
	}
	return currentJoltages
}

func findLowestButtonInputWiring(machine Machine, debug bool, purpose string) int {

	lowestPresses := 10000

	// brute force it for 10,000 random attempts
	for i := 0; i < 10000; i++ {

		buttonPresses := 0

		currentIndicators := machine.intialIndicators
		for machine.desiredIndicators != currentIndicators {

			if buttonPresses > lowestPresses {
				break
			}

			// press a random button
			currentIndicators = pressButtonWiring(machine, rand.Intn(len(machine.buttons)), currentIndicators)
			buttonPresses += 1

			if debug {
				fmt.Println("Current Indicators:", currentIndicators)
			}
		}

		if buttonPresses < lowestPresses {
			lowestPresses = buttonPresses
		}
	}

	return lowestPresses
}

func joltageTooHigh(desiredJoltages []int, currentJoltages []int) bool {
	for i := 0; i < len(desiredJoltages); i++ {
		if currentJoltages[i] > desiredJoltages[i] {
			return true
		}
	}
	return false
}

func joltageEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func joltageIsdivisible(desiredJoltages []int, currentJoltages []int) (bool, int) {

	for c := range currentJoltages {
		if currentJoltages[c] == 0 {
			return false, 0
		}
	}

	divisibleBy := float64(desiredJoltages[0]) / float64(currentJoltages[0])
	if divisibleBy != float64(int(divisibleBy)) {
		return false, 0
	}

	// fmt.Println("Checking divisibility by", divisibleBy)
	for i := 0; i < len(desiredJoltages); i++ {
		if float64(desiredJoltages[i])/float64(currentJoltages[i]) != divisibleBy {
			return false, 0
		}
	}

	return true, int(divisibleBy)
}

func findLowestButtonInputJoltage(machine Machine, debug bool, purpose string, initialLowest int) int {

	lowestPresses := initialLowest
	if lowestPresses == 0 {
		lowestPresses = 1000000
	}

	// brute force with random attempts (more iterations if no solution cached)
	iterations := 1000000
	if initialLowest >= 1000000 || initialLowest == 0 {
		iterations = 5000000
	}
	for i := 0; i < iterations; i++ {

		buttonPresses := 0

		// Make a proper copy of the slice
		currentJoltages := make([]int, len(machine.initalJoltages))
		copy(currentJoltages, machine.initalJoltages)

		currentButtons := make([][]int, len(machine.buttons))
		copy(currentButtons, machine.buttons)

		for !joltageEqual(machine.desiredJoltages, currentJoltages) {

			if buttonPresses >= lowestPresses || joltageTooHigh(machine.desiredJoltages, currentJoltages) || len(currentButtons) == 0 {
				break
			}

			// Check if current joltages are a divisor of desired (optimization to exit early)
			isDivisible, divisibleBy := joltageIsdivisible(machine.desiredJoltages, currentJoltages)
			if isDivisible && divisibleBy > 1 {

				if buttonPresses*divisibleBy < lowestPresses {

					// We can scale up - the total cost is buttonPresses * divisibleBy
					buttonPresses *= divisibleBy
					// Update joltages to match desired
					for i := 0; i < len(currentJoltages); i++ {
						currentJoltages[i] *= divisibleBy
					}
					break
				}
			}
			buttonToPress := rand.Intn(len(currentButtons))

			// allButtonsAreEven := true

			// for _, btn := range currentButtons[buttonToPress] {
			// 	if btn%2 != 0 {
			// 		allButtonsAreEven = false
			// 		break
			// 	}
			// }

			// Press the button directly using the button from currentButtons
			for _, joltagesToIncrease := range currentButtons[buttonToPress] {
				currentJoltages = increaseJoltage(currentJoltages, joltagesToIncrease)
			}

			// if allButtonsAreEven {
			// 	// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/?share_id=fUwRdxhFQjdSP0Z_yhx-O&utm_content=2&utm_medium=android_app&utm_name=androidcss&utm_source=share&utm_term=2
			// 	// only trial press each button once and check if joltages can be multiplied to reach target
			// 	newCurrentButtons := [][]int{}
			// 	newCurrentButtons = append(newCurrentButtons, currentButtons[:buttonToPress]...)
			// 	newCurrentButtons = append(newCurrentButtons, currentButtons[buttonToPress+1:]...)
			// 	currentButtons = newCurrentButtons
			// }

			buttonPresses += 1

			if debug && buttonPresses%100 == 0 {
				fmt.Println("Button presses:", buttonPresses, "Current:", currentJoltages, "Desired:", machine.desiredJoltages)
			}
		}

		if buttonPresses < lowestPresses {
			if joltageEqual(machine.desiredJoltages, currentJoltages) {
				lowestPresses = buttonPresses
				if debug {
					fmt.Println("New lowest:", lowestPresses)
				}
			}
		}
	}

	if lowestPresses == 1000000 {
		return findLowestButtonInputJoltage(machine, debug, purpose, initialLowest)
	}

	return lowestPresses
}

func loadCache(filename string) map[int]int {
	cache := make(map[int]int)
	cacheFile := strings.Replace(filename, ".txt", "_cache.json", 1)
	data, err := os.ReadFile(cacheFile)
	if err == nil {
		json.Unmarshal(data, &cache)
	}
	return cache
}

func saveCache(filename string, cache map[int]int) {
	cacheFile := strings.Replace(filename, ".txt", "_cache.json", 1)
	data, err := json.MarshalIndent(cache, "", "  ")
	if err == nil {
		os.WriteFile(cacheFile, data, 0644)
	}
}

func linearAlgerbra(mahcine Machine) int {
	target := mahcine.desiredJoltages
	numButtons := len(mahcine.buttons)
	numJoltages := len(target)

	// Build the coefficient matrix A where A[i][j] = 1 if button j affects joltage i
	// A is [numJoltages x numButtons]
	AData := make([]float64, numJoltages*numButtons)
	bData := make([]float64, numJoltages)

	for joltageIdx := 0; joltageIdx < numJoltages; joltageIdx++ {
		bData[joltageIdx] = float64(target[joltageIdx])

		for buttonIdx := 0; buttonIdx < numButtons; buttonIdx++ {
			// Check if this button affects this joltage
			affectsJoltage := false
			for _, j := range mahcine.buttons[buttonIdx] {
				if j == joltageIdx {
					affectsJoltage = true
					break
				}
			}
			if affectsJoltage {
				AData[joltageIdx*numButtons+buttonIdx] = 1
			}
		}
	}

	A := mat.NewDense(numJoltages, numButtons, AData)
	b := mat.NewVecDense(numJoltages, bData)

	var x mat.VecDense
	
	// For underdetermined systems (more buttons than joltages), find minimum norm solution
	// For overdetermined systems (more joltages than buttons), find least squares solution
	// For square systems, find exact solution
	// SolveVec handles all cases using QR decomposition
	err := x.SolveVec(A, b)
	if err != nil {
		// If direct solve fails, try pseudo-inverse approach for underdetermined systems
		if numButtons > numJoltages {
			// Minimum norm solution: x = A^T(AA^T)^-1 b
			var AT mat.Dense
			AT.CloneFrom(A.T())
			
			var AAT mat.Dense
			AAT.Mul(A, &AT)
			
			var AATinv mat.Dense
			err := AATinv.Inverse(&AAT)
			if err != nil {
				return -1
			}
			
			var temp mat.VecDense
			temp.MulVec(&AATinv, b)
			
			x.MulVec(&AT, &temp)
		} else {
			return -1
		}
	}

	// Extract solution and verify it's valid (non-negative integers)
	solution := make([]int, numButtons)
	totalPresses := 0

	for i := 0; i < numButtons; i++ {
		val := x.AtVec(i)

		// Must be non-negative
		if val < -0.001 {
			return -1
		}

		// Round to nearest integer
		rounded := int(math.Round(val))
		if rounded < 0 {
			rounded = 0
		}

		// Check if rounding error is too large
		if math.Abs(val-float64(rounded)) > 0.1 {
			return -1
		}

		solution[i] = rounded
		totalPresses += rounded
	}

	// Verify the solution actually reaches the target
	result := make([]int, numJoltages)
	for buttonIdx := 0; buttonIdx < numButtons; buttonIdx++ {
		for _, joltageIdx := range mahcine.buttons[buttonIdx] {
			result[joltageIdx] += solution[buttonIdx]
		}
	}

	// Check if solution matches target
	for i := 0; i < numJoltages; i++ {
		if result[i] != target[i] {
			return -1
		}
	}

	return totalPresses
}

// Helper function for bounded search
func boundedSearch(mahcine Machine, startingSolution []float64) int {
	numButtons := len(mahcine.buttons)
	numJoltages := len(mahcine.desiredJoltages)
	target := mahcine.desiredJoltages
	
	// Create search bounds based on number of buttons
	var searchBound int
	if numButtons <= 5 {
		searchBound = 15
	} else if numButtons <= 8 {
		searchBound = 10
	} else if numButtons <= 12 {
		searchBound = 6
	} else if numButtons <= 15 {
		searchBound = 4
	} else {
		searchBound = 3
	}
	
	baseSolution := make([]int, numButtons)
	for i := 0; i < numButtons; i++ {
		baseSolution[i] = int(math.Round(startingSolution[i]))
		if baseSolution[i] < 0 {
			baseSolution[i] = 0
		}
	}
	
	// Calculate total search space
	totalCombos := 1
	for i := 0; i < numButtons && totalCombos <= 1000000; i++ {
		totalCombos *= (2*searchBound + 1)
	}
	
	// If search space is too large, reduce bound
	maxCombos := 500000
	for totalCombos > maxCombos && searchBound > 1 {
		searchBound--
		totalCombos = 1
		for i := 0; i < numButtons && totalCombos <= maxCombos; i++ {
			totalCombos *= (2*searchBound + 1)
		}
	}
	
	if totalCombos > maxCombos {
		return -1 // Too large to search
	}
	
	// Recursive search through all combinations
	bestSolution := -1
	candidate := make([]int, numButtons)
	copy(candidate, baseSolution)
	
	var search func(int)
	search = func(buttonIdx int) {
		if buttonIdx >= numButtons {
			// Check if this candidate is valid
			resultJoltages := make([]int, numJoltages)
			for btnIdx, presses := range candidate {
				for _, joltageIdx := range mahcine.buttons[btnIdx] {
					resultJoltages[joltageIdx] += presses
				}
			}
			
			// Verify matches target
			valid := true
			for i := 0; i < numJoltages; i++ {
				if resultJoltages[i] != target[i] {
					valid = false
					break
				}
			}
			
			if valid {
				sum := 0
				for _, p := range candidate {
					sum += p
				}
				if bestSolution == -1 || sum < bestSolution {
					bestSolution = sum
				}
			}
			return
		}
		
		// Try values in range [base-bound, base+bound]
		baseVal := baseSolution[buttonIdx]
		for offset := -searchBound; offset <= searchBound; offset++ {
			val := baseVal + offset
			if val < 0 {
				val = 0
			}
			candidate[buttonIdx] = val
			
			// Early pruning: if current sum already exceeds best, skip
			if bestSolution > 0 {
				currentSum := 0
				for i := 0; i <= buttonIdx; i++ {
					currentSum += candidate[i]
				}
				if currentSum > bestSolution {
					continue
				}
			}
			
			search(buttonIdx + 1)
		}
	}
	
	search(0)
	return bestSolution
}

// Exhaustive search starting from zero for square/degenerate systems
func exhaustiveFromZero(mahcine Machine) int {
	numButtons := len(mahcine.buttons)
	numJoltages := len(mahcine.desiredJoltages)
	target := mahcine.desiredJoltages
	
	// Calculate max possible value we might need
	maxTarget := 0
	for _, t := range target {
		if t > maxTarget {
			maxTarget = t
		}
	}
	
	// Set reasonable bounds
	maxPerButton := maxTarget + 10
	if maxPerButton > 50 {
		maxPerButton = 50
	}
	
	bestSolution := -1
	candidate := make([]int, numButtons)
	
	var search func(int, []int)
	search = func(buttonIdx int, currentJoltages []int) {
		if buttonIdx >= numButtons {
			// Check if solution is valid
			valid := true
			for i := 0; i < numJoltages; i++ {
				if currentJoltages[i] != target[i] {
					valid = false
					break
				}
			}
			
			if valid {
				sum := 0
				for _, p := range candidate {
					sum += p
				}
				if bestSolution == -1 || sum < bestSolution {
					bestSolution = sum
				}
			}
			return
		}
		
		// Prune if current sum exceeds best
		if bestSolution > 0 {
			currentSum := 0
			for i := 0; i < buttonIdx; i++ {
				currentSum += candidate[i]
			}
			if currentSum >= bestSolution {
				return
			}
		}
		
		buttonAffects := mahcine.buttons[buttonIdx]
		
		// Calculate smart upper bound for this button
		maxNeeded := 0
		for _, joltageIdx := range buttonAffects {
			remaining := target[joltageIdx] - currentJoltages[joltageIdx]
			if remaining > maxNeeded {
				maxNeeded = remaining
			}
		}
		
		upperLimit := maxNeeded + 5
		if upperLimit > maxPerButton {
			upperLimit = maxPerButton
		}
		
		// Try values from 0 to upper limit
		for val := 0; val <= upperLimit; val++ {
			candidate[buttonIdx] = val
			
			// Update joltages
			newJoltages := make([]int, numJoltages)
			copy(newJoltages, currentJoltages)
			for _, joltageIdx := range buttonAffects {
				newJoltages[joltageIdx] += val
			}
			
			// Prune if exceeded any target
			exceeded := false
			for i := 0; i < numJoltages; i++ {
				if newJoltages[i] > target[i] {
					exceeded = true
					break
				}
			}
			
			if !exceeded {
				search(buttonIdx+1, newJoltages)
			}
			
			// Early exit if found solution
			if bestSolution > 0 && val > 0 {
				break
			}
		}
	}
	
	initialJoltages := make([]int, numJoltages)
	search(0, initialJoltages)
	
	return bestSolution
}

// Branch and bound search with constraint propagation
func branchAndBound(mahcine Machine, startingSolution []float64) int {
	numButtons := len(mahcine.buttons)
	numJoltages := len(mahcine.desiredJoltages)
	target := mahcine.desiredJoltages
	
	// Create initial bounds based on machine size
	minBound := make([]int, numButtons)
	maxBound := make([]int, numButtons)
	
	// Check if we have a meaningful LP solution
	hasLPSolution := false
	for i := 0; i < numButtons; i++ {
		if startingSolution[i] > 0.1 {
			hasLPSolution = true
			break
		}
	}
	
	var searchRadius int
	if hasLPSolution {
		// Use tighter bounds around LP solution
		if numButtons <= 5 {
			searchRadius = 15
		} else if numButtons <= 7 {
			searchRadius = 12
		} else if numButtons <= 10 {
			searchRadius = 10
		} else {
			searchRadius = 8
		}
	} else {
		// No LP solution, need wider search from zero
		if numButtons <= 5 {
			searchRadius = 30
		} else if numButtons <= 7 {
			searchRadius = 20
		} else if numButtons <= 10 {
			searchRadius = 15
		} else {
			searchRadius = 12
		}
	}
	
	for i := 0; i < numButtons; i++ {
		base := int(math.Round(startingSolution[i]))
		if base < 0 {
			base = 0
		}
		minBound[i] = base - searchRadius
		if minBound[i] < 0 {
			minBound[i] = 0
		}
		maxBound[i] = base + searchRadius
	}
	
	bestSolution := -1
	candidate := make([]int, numButtons)
	
	var search func(int, []int)
	search = func(buttonIdx int, currentJoltages []int) {
		if buttonIdx >= numButtons {
			// Check if solution is valid
			valid := true
			for i := 0; i < numJoltages; i++ {
				if currentJoltages[i] != target[i] {
					valid = false
					break
				}
			}
			
			if valid {
				sum := 0
				for _, p := range candidate {
					sum += p
				}
				if bestSolution == -1 || sum < bestSolution {
					bestSolution = sum
				}
			}
			return
		}
		
		// Calculate current sum for pruning
		currentSum := 0
		for i := 0; i < buttonIdx; i++ {
			currentSum += candidate[i]
		}
		
		// Prune if current sum already exceeds best
		if bestSolution > 0 && currentSum >= bestSolution {
			return
		}
		
		// Calculate lower bound: how many more button presses we need at minimum
		remainingDeficit := 0
		for i := 0; i < numJoltages; i++ {
			if currentJoltages[i] < target[i] {
				remainingDeficit += (target[i] - currentJoltages[i])
			}
		}
		
		// Prune if lower bound exceeds best known solution
		if bestSolution > 0 && currentSum+remainingDeficit >= bestSolution {
			return
		}
		
		// Try values for this button
		buttonAffects := mahcine.buttons[buttonIdx]
		
		// Calculate useful range based on what we still need
		minNeeded := 0
		
		// If this button affects joltages we need to increase, prioritize higher values
		needsIncrease := false
		for _, joltageIdx := range buttonAffects {
			if currentJoltages[joltageIdx] < target[joltageIdx] {
				needsIncrease = true
				deficit := target[joltageIdx] - currentJoltages[joltageIdx]
				if deficit > minNeeded {
					minNeeded = deficit
				}
			}
		}
		
		// Constrain search range
		start := minBound[buttonIdx]
		if needsIncrease && start < minNeeded {
			start = minNeeded
		}
		end := maxBound[buttonIdx]
		
		// Try values in order (prefer smaller values to minimize total)
		for val := start; val <= end; val++ {
			candidate[buttonIdx] = val
			
			// Update joltages for next level
			newJoltages := make([]int, numJoltages)
			copy(newJoltages, currentJoltages)
			for _, joltageIdx := range buttonAffects {
				newJoltages[joltageIdx] += val
			}
			
			// Check if any joltage exceeded target (prune this branch)
			exceeded := false
			for i := 0; i < numJoltages; i++ {
				if newJoltages[i] > target[i] {
					exceeded = true
					break
				}
			}
			
			if !exceeded {
				search(buttonIdx+1, newJoltages)
			}
		}
	}
	
	initialJoltages := make([]int, numJoltages)
	search(0, initialJoltages)
	
	return bestSolution
}

func integerLinearAlgebra(mahcine Machine) int {
	target := mahcine.desiredJoltages
	numButtons := len(mahcine.buttons)
	numJoltages := len(target)

	// First, try standard linear algebra
	result := linearAlgerbra(mahcine)
	if result > 0 {
		return result // Already found integer solution
	}

	// Build the coefficient matrix to check for degeneracy
	AData := make([]float64, numJoltages*numButtons)
	bData := make([]float64, numJoltages)

	for joltageIdx := 0; joltageIdx < numJoltages; joltageIdx++ {
		bData[joltageIdx] = float64(target[joltageIdx])

		for buttonIdx := 0; buttonIdx < numButtons; buttonIdx++ {
			affectsJoltage := false
			for _, j := range mahcine.buttons[buttonIdx] {
				if j == joltageIdx {
					affectsJoltage = true
					break
				}
			}
			if affectsJoltage {
				AData[joltageIdx*numButtons+buttonIdx] = 1
			}
		}
	}

	A := mat.NewDense(numJoltages, numButtons, AData)
	b := mat.NewVecDense(numJoltages, bData)


	// Get the continuous solution as a starting point
	AData = make([]float64, numJoltages*numButtons)
	bData = make([]float64, numJoltages)

	for joltageIdx := 0; joltageIdx < numJoltages; joltageIdx++ {
		bData[joltageIdx] = float64(target[joltageIdx])

		for buttonIdx := 0; buttonIdx < numButtons; buttonIdx++ {
			affectsJoltage := false
			for _, j := range mahcine.buttons[buttonIdx] {
				if j == joltageIdx {
					affectsJoltage = true
					break
				}
			}
			if affectsJoltage {
				AData[joltageIdx*numButtons+buttonIdx] = 1
			}
		}
	}

	A = mat.NewDense(numJoltages, numButtons, AData)
	b = mat.NewVecDense(numJoltages, bData)

	// Get the continuous solution as a starting point
	var x mat.VecDense
	err := x.SolveVec(A, b)

	var startingSolution []float64
	if err == nil {
		startingSolution = make([]float64, numButtons)
		for i := 0; i < numButtons; i++ {
			startingSolution[i] = x.AtVec(i)
		}
	} else {
		// Try minimum norm solution for underdetermined
		if numButtons > numJoltages {
			var AT mat.Dense
			AT.CloneFrom(A.T())

			var AAT mat.Dense
			AAT.Mul(A, &AT)

			var AATinv mat.Dense
			err := AATinv.Inverse(&AAT)
			if err == nil {
				var temp mat.VecDense
				temp.MulVec(&AATinv, b)

				x.MulVec(&AT, &temp)

				startingSolution = make([]float64, numButtons)
				for i := 0; i < numButtons; i++ {
					startingSolution[i] = x.AtVec(i)
				}
			}
		}
	}
	
	// If no LP solution found, create a zero starting solution for bounded search
	if startingSolution == nil {
		startingSolution = make([]float64, numButtons)
		// Initialize with zero or small values
		for i := 0; i < numButtons; i++ {
			startingSolution[i] = 0
		}
	}

	// Try different rounding strategies
	if startingSolution != nil {
		bestSolution := -1
		
		strategies := []func(float64) int{
			func(v float64) int { return int(math.Floor(v)) },
			func(v float64) int { return int(math.Ceil(v)) },
			func(v float64) int { return int(math.Round(v)) },
		}

		// Try each strategy
		for _, strategy := range strategies {
			candidate := make([]int, numButtons)
			for i := 0; i < numButtons; i++ {
				val := strategy(startingSolution[i])
				if val < 0 {
					val = 0
				}
				candidate[i] = val
			}

			// Check if valid
			resultJoltages := make([]int, numJoltages)
			for btnIdx, presses := range candidate {
				for _, joltageIdx := range mahcine.buttons[btnIdx] {
					resultJoltages[joltageIdx] += presses
				}
			}

			valid := true
			for i := 0; i < numJoltages; i++ {
				if resultJoltages[i] != target[i] {
					valid = false
					break
				}
			}

			if valid {
				sum := 0
				for _, p := range candidate {
					sum += p
				}
				if bestSolution == -1 || sum < bestSolution {
					bestSolution = sum
				}
			}
		}

		// Try bounded exhaustive search around the LP solution
		boundedResult := boundedSearch(mahcine, startingSolution)
		if boundedResult > 0 {
			if bestSolution == -1 || boundedResult < bestSolution {
				bestSolution = boundedResult
			}
		}
		
		// Try branch-and-bound with constraint propagation
		branchBoundResult := branchAndBound(mahcine, startingSolution)
		if branchBoundResult > 0 {
			if bestSolution == -1 || branchBoundResult < bestSolution {
				bestSolution = branchBoundResult
			}
		}

		// Try adjusting the rounded solution to make it valid
		candidate := make([]int, numButtons)
		for i := 0; i < numButtons; i++ {
			candidate[i] = int(math.Round(startingSolution[i]))
			if candidate[i] < 0 {
				candidate[i] = 0
			}
		}

		// Iteratively adjust to meet constraints (simple greedy approach)
		maxIterations := 1000
		for iter := 0; iter < maxIterations; iter++ {
			resultJoltages := make([]int, numJoltages)
			for btnIdx, presses := range candidate {
				for _, joltageIdx := range mahcine.buttons[btnIdx] {
					resultJoltages[joltageIdx] += presses
				}
			}

			// Check differences
			allMatch := true
			for i := 0; i < numJoltages; i++ {
				diff := target[i] - resultJoltages[i]
				if diff != 0 {
					allMatch = false
					// Find a button that affects this joltage
					for btnIdx := 0; btnIdx < numButtons; btnIdx++ {
						affects := false
						for _, j := range mahcine.buttons[btnIdx] {
							if j == i {
								affects = true
								break
							}
						}
						if affects {
							if diff > 0 {
								candidate[btnIdx]++
							} else if candidate[btnIdx] > 0 {
								candidate[btnIdx]--
							}
							break
						}
					}
					break
				}
			}

			if allMatch {
				sum := 0
				for _, p := range candidate {
					sum += p
				}
				// Only update if this is better than what we already found
				if bestSolution == -1 || sum < bestSolution {
					bestSolution = sum
				}
				break
			}
		}
		
		// Return the best solution found across all strategies
		if bestSolution > 0 {
			return bestSolution
		}
	}

	return -1 // Failed to find integer solution
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Part1(filename string, debug bool) int {
	machines := Parse(filename)

	ans := 0
	for _, machine := range machines {

		ans += findLowestButtonInputWiring(machine, debug, "wiring")
	}

	return ans
}

// func Part2(filename string, debug bool) int {
// 	machines := Parse(filename)
// 	cache := loadCache(filename)

// 	ans := 0
// 	i := 0
// 	for machineID, machine := range machines {
// 		fmt.Println("Processing Machines", i*100/len(machines), "%")
// 		initialLowest := cache[machineID]
// 		if initialLowest > 0 && initialLowest <= 1000000 {
// 			fmt.Printf("  Machine %d: Using cached lowest = %d\n", machineID, initialLowest)
// 		}

// 		result := findLowestButtonInputJoltage(machine, debug, "joltage", initialLowest)
// 		// Only save real solutions to cache (not max values)
// 		if result < 1000000 && (result < initialLowest || initialLowest == 0) {
// 			cache[machineID] = result
// 			fmt.Printf("  Machine %d: New lowest = %d (updated cache)\n", machineID, result)
// 		} else if result >= 1000000 {
// 			fmt.Printf("  Machine %d: No solution found yet\n", machineID)
// 		}
// 		ans += result
// 		i++
// 	}

// 	saveCache(filename, cache)
// 	fmt.Println("Cache saved!")
// 	return ans
// }

func Part2(filename string, debug bool) int {
	machines := Parse(filename)
	cache := loadCache(filename)

	// First pass: Use ILP to solve as many machines as possible
	fmt.Println("=== First Pass: Integer Linear Algebra ===")
	for machineID, machine := range machines {
		if machineID%10 == 0 {
			fmt.Printf("ILP Progress: %d/%d machines (%d%%)\n", machineID, len(machines), machineID*100/len(machines))
		}
		
		// Reload cache to get latest values from other processes
		cache = loadCache(filename)
		
		// Skip if already in cache
		if cache[machineID] > 0 && cache[machineID] < 1000000 {
			continue
		}
		
		result := integerLinearAlgebra(machine)
		if result > 0 {
			// Reload cache before saving to avoid overwriting better solutions
			cache = loadCache(filename)
			if cache[machineID] == 0 || cache[machineID] >= 1000000 || result < cache[machineID] {
				cache[machineID] = result
				saveCache(filename, cache)
				fmt.Printf("  Machine %d: ILP found %d (saved)\n", machineID, result)
			}
		}
	}
	fmt.Println("First pass complete!")
	
	// Reload cache to get all solutions found by this and other processes
	cache = loadCache(filename)
	
	// Count how many unsolved
	unsolved := 0
	unsolvedIDs := []int{}
	for machineID := range machines {
		if cache[machineID] == 0 || cache[machineID] >= 1000000 {
			unsolved++
			unsolvedIDs = append(unsolvedIDs, machineID)
		}
	}
	fmt.Printf("ILP solved %d/%d machines. %d remaining.\n", len(machines)-unsolved, len(machines), unsolved)
	
	// Second pass: Use random search for unsolved machines
	if unsolved > 0 {
		fmt.Println("\n=== Second Pass: Random Search for Unsolved Machines ===")
		for i, machineID := range unsolvedIDs {
			// Reload cache to check if another process already solved it
			cache = loadCache(filename)
			if cache[machineID] > 0 && cache[machineID] < 1000000 {
				fmt.Printf("Random search %d/%d: Machine %d already solved by another process (%d)\n", i+1, unsolved, machineID, cache[machineID])
				continue
			}
			
			fmt.Printf("Random search %d/%d: Machine %d...\n", i+1, unsolved, machineID)
			machine := machines[machineID]
			result := findLowestButtonInputJoltage(machine, false, "joltage", 0)
			if result < 1000000 {
				// Reload cache before saving
				cache = loadCache(filename)
				if cache[machineID] == 0 || cache[machineID] >= 1000000 || result < cache[machineID] {
					cache[machineID] = result
					saveCache(filename, cache)
					fmt.Printf("  Machine %d: Random search found = %d (saved)\n", machineID, result)
				} else {
					fmt.Printf("  Machine %d: Found %d but cache has better %d\n", machineID, result, cache[machineID])
				}
			} else {
				fmt.Printf("  Machine %d: No solution found\n", machineID)
			}
		}
	}
	
	// Calculate final answer from latest cache
	cache = loadCache(filename)
	ans := 0
	for machineID := range machines {
		result := cache[machineID]
		if result > 0 && result < 1000000 {
			ans += result
		}
	}
	
	fmt.Println("All passes complete!")
	return ans
}

func main() {
	fmt.Println("Advent-Of-Code 2025 - Day10")
	// after creating a cache and trying a few optimizations, I ended up cheating and used AI for part 2
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt", false))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt", false))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt", true)) // should be 33
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt", false))
}
