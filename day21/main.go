package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

func main() {

	fmt.Println("Solution1", solution1("input.txt", 64, false))
	fmt.Println("Solution2", solution2("input.txt", 26501365))

}

func solution2(input string, maxSteps int) int {

	lines := utils.FileToArray(input)

	field := make([][]string, len(lines))

	start := [2]int{0, 0}
	for i, line := range lines {

		field[i] = utils.StringToStringArray(line, "")

		if strings.ContainsAny(line, "S") {
			start = [2]int{i, strings.IndexAny(line, "S")}
		}
	}
	field[start[0]][start[1]] = "."

	stateMap := map[string]int{}

	// breakpoint for the loop below
	b := maxSteps

	nextSteps := make([][2]int, 0)
	nextSteps = append(nextSteps, start)

	cycleStart, cycleLength := 0, 0
	for i := 0; i < maxSteps; i++ {
		// cycle the map
		ns := make([][2]int, 0)
		for _, nextStep := range nextSteps {
			ns = append(ns, makeStep(field, nextStep)...)
		}
		nextSteps = uniqueSteps(ns)

		h := joinMatrix(field) + "|" + fmt.Sprint(nextSteps)
		// if we haven't changed the breakpoint yet

		utils.PrintMatrix(field)
		if b == maxSteps {
			// check if we have seen this state before
			if prev, ok := stateMap[h]; ok {
				// if we have, we can calculate the breakpoint
				// current step + remaining steps % cycle_length
				b = i + (maxSteps-prev)%(i-prev)
				fmt.Println("b", b, "i", i, "prev", prev, "maxSteps", maxSteps)
			} else {
				stateMap[h] = i
			}
		}

		if i == b {
			cycleStart = stateMap[h]
			cycleLength = i - stateMap[h]
			println("Cycle start", cycleStart, "cycle length", cycleLength)
			break
		}
	}

	t := calculateField(field)

	fmt.Println(t)

	k := maxSteps / cycleStart
	fmt.Println(k)

	return 0
}

func calculateField(field [][]string) int {
	count := 0
	for _, line := range field {
		for _, char := range line {
			if char == "O" {
				count++

			}
		}
	}
	return count
}

func joinMatrix(m [][]string) string {
	result := ""
	for _, row := range m {
		for _, item := range row {
			result += item
		}
	}
	return result
}

// part 1
func solution1(input string, maxSteps int, wrap bool) int {
	lines := utils.FileToArray(input)

	field := make([][]string, len(lines))

	start := [2]int{0, 0}
	for i, line := range lines {

		field[i] = utils.StringToStringArray(line, "")

		if strings.ContainsAny(line, "S") {
			start = [2]int{i, strings.IndexAny(line, "S")}
		}
	}
	field[start[0]][start[1]] = "."

	calcField := make([][]string, len(field))
	for i, line := range field {
		calcField[i] = make([]string, len(line))
		copy(calcField[i], line)
	}

	nextSteps := make([][2]int, 0)
	nextSteps = append(nextSteps, start)
	for step := 0; step < maxSteps; step++ {

		// on each step, we check all the next steps and then overwrite them
		ns := make([][2]int, 0)
		for _, nextStep := range nextSteps {
			ns = append(ns, makeStep(calcField, nextStep)...)
		}
		nextSteps = uniqueSteps(ns)

		// utils.PrintMatrix(calcField)
	}

	count := 0
	for _, line := range calcField {
		for _, char := range line {
			if char == "O" {
				count++

			}
		}
	}

	return count

}

func uniqueSteps(steps [][2]int) [][2]int {
	unique := make(map[[2]int]bool)
	for _, step := range steps {
		unique[step] = true
	}

	result := make([][2]int, 0)
	for step := range unique {
		result = append(result, step)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i][0] == result[j][0] {
			return result[i][1] < result[j][1]
		}
		return result[i][0] < result[j][0]
	})
	return result
}

var directions = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func makeStep(field [][]string, start [2]int) [][2]int {

	nextSteps := make([][2]int, 0)
	for _, direction := range directions {
		x := start[0] + direction[0]
		y := start[1] + direction[1]

		// borders
		if x < 0 || y < 0 || x >= len(field) || y >= len(field[0]) {
			continue
		}

		next := field[x][y]

		if next == "." || next == "O" {
			field[x][y] = "O"
			field[start[0]][start[1]] = "."

			nextSteps = append(nextSteps, [2]int{x, y})
		}
	}
	return nextSteps
}

// // part 2

// func solution2(input string, maxSteps int, wrap bool) int {
// 	lines := utils.FileToArray(input)

// 	field := make([][]string, len(lines))

// 	start := [2]int{0, 0}
// 	for i, line := range lines {
// 		field[i] = utils.StringToStringArray(line, "")
// 		if strings.ContainsAny(line, "S") {
// 			start = [2]int{i, strings.IndexAny(line, "S")}
// 		}
// 	}
// 	field[start[0]][start[1]] = "."

// 	fields := make(map[string]FieldResult)

// 	// we go top to bottom and calculate for 0 and len(field[0])-1
// 	for i := 0; i < len(field); i++ {

// 		start := [2]int{i, 0}

// 		key := HashField(field, start)
// 		if _, ok := fields[key]; !ok {
// 			fields[key] = processField(field, start)
// 		}

// 		start = [2]int{i, len(field[0]) - 1}
// 		key = HashField(field, start)
// 		if _, ok := fields[key]; !ok {
// 			fields[key] = processField(field, start)
// 		}

// 		fmt.Println("i", i, "len", len(field))
// 	}

// 	for j := 0; j < len(field[0]); j++ {

// 		start := [2]int{0, j}

// 		key := HashField(field, start)
// 		if _, ok := fields[key]; !ok {
// 			fields[key] = processField(field, start)
// 		}

// 		start = [2]int{len(field) - 1, j}
// 		key = HashField(field, start)
// 		if _, ok := fields[key]; !ok {
// 			fields[key] = processField(field, start)
// 		}

// 		fmt.Println("j", j, "len", len(field))
// 	}

// 	return len(fields)
// }

// // for field we need to know how many steps we have to do before we reach next field
// // and how many final points we have for each of those steps

// func HashField(field [][]string, start [2]int) string {
// 	result := ""
// 	for _, line := range field {
// 		result += strings.Join(line, "")
// 	}
// 	return result + fmt.Sprintf("%v", start)
// }

// type FieldResult struct {
// 	StepsToFullExit int
// 	CellsTotal      int
// 	// this will be number of cells we reached given this many steps through the field
// 	StepsToCells map[int]int
// 	// for each exit point we record how many steps it took to reach it
// 	StepsToExit []struct {
// 		exit  [2]int
// 		steps int
// 	}
// }

// // this function calculates the field for the given start point and returns
// func processField(field [][]string, start [2]int) FieldResult {

// 	result := FieldResult{CellsTotal: 0, StepsToCells: make(map[int]int), StepsToExit: make([]struct {
// 		exit  [2]int
// 		steps int
// 	}, 0)}

// 	nextSteps := make([]StepResult, 0)
// 	nextSteps = append(nextSteps, StepResult{start, false})

// 	for steps := 0; len(nextSteps) > 0; steps++ {
// 		// on each step, we check all the next steps and then overwrite them
// 		ns := make([]StepResult, 0)

// 		for _, nextStep := range nextSteps {
// 			ns = append(ns, makeStepWithExit(field, nextStep.position)...)
// 		}

// 		// we went through all the steps, we need to grab exists
// 		for _, nextStep := range ns {
// 			if nextStep.outOfBounds {

// 				result.StepsToExit = append(result.StepsToExit, struct {
// 					exit  [2]int
// 					steps int
// 				}{nextStep.position, steps})

// 			}
// 		}

// 		// then we unique them and exclude all that are out of bounds

// 		nextSteps = uniqueStepsInBounds(ns)

// 		if steps < 20 {
// 			fmt.Println("start", start, "step", steps, "nextSteps", len(nextSteps))
// 		}
// 		if len(nextSteps) == 0 {
// 			result.StepsToFullExit = steps
// 			break
// 		}

// 		count := 0
// 		for _, line := range field {
// 			for _, char := range line {
// 				if char == "O" {
// 					count++

// 				}
// 			}
// 		}

// 		result.CellsTotal = count
// 		result.StepsToCells[steps] = count
// 	}

// 	return result
// }

// type StepResult struct {
// 	position    [2]int
// 	outOfBounds bool
// }

// func makeStepWithExit(field [][]string, start [2]int) []StepResult {

// 	nextSteps := make([]StepResult, 0)

// 	for _, direction := range directions {
// 		x := start[0] + direction[0]
// 		y := start[1] + direction[1]

// 		// borders
// 		if x < 0 || y < 0 || x >= len(field) || y >= len(field[0]) {
// 			nextSteps = append(nextSteps, StepResult{position: [2]int{x, y}, outOfBounds: true})
// 			continue
// 		}

// 		next := field[x][y]

// 		if next == "." || next == "O" {
// 			field[x][y] = "O"
// 			field[start[0]][start[1]] = "."

// 			nextSteps = append(nextSteps, StepResult{position: [2]int{x, y}, outOfBounds: false})
// 		}
// 	}
// 	return nextSteps
// }

// func uniqueStepsInBounds(steps []StepResult) []StepResult {

// 	unique := make(map[[2]int]bool)

// 	for _, step := range steps {
// 		if step.outOfBounds {
// 			continue
// 		}
// 		unique[step.position] = true
// 	}

// 	result := make([]StepResult, 0)
// 	for step := range unique {
// 		result = append(result, struct {
// 			position    [2]int
// 			outOfBounds bool
// 		}{step, false})
// 	}
// 	return result
// }
