package main

import (
	"fmt"
	"sort"
	"utils"
)

func main() {

	result := solution("input.txt")
	println(result)

	result2 := solution2("input.txt")
	println(result2)
}

type Direction [2]int

var (
	up    Direction = [2]int{-1, 0}
	down            = [2]int{1, 0}
	left            = [2]int{0, -1}
	right           = [2]int{0, 1}
)

var directions = []Direction{up, down, left, right}

// we look at next char and check if we connect to it from the direction we came from
func canConnect(matrix [][]string, currentCoordinates [2]int, incomingDirection Direction) bool {

	currentSymbol := matrix[currentCoordinates[0]][currentCoordinates[1]]

	maxI := len(matrix)
	maxJ := len(matrix[0])

	nextCoordinates := [2]int{currentCoordinates[0] + incomingDirection[0], currentCoordinates[1] + incomingDirection[1]}
	if nextCoordinates[0] < 0 || nextCoordinates[0] >= maxI || nextCoordinates[1] < 0 || nextCoordinates[1] >= maxJ {
		return false
	}
	nextSymbol := matrix[nextCoordinates[0]][nextCoordinates[1]]
	//  = matrix[start[0]+incomingDirection[0]][start[1]+incomingDirection[1]]
	// // fmt.Println("canConnect", a, b, direction)
	switch nextSymbol {
	case "|":
		if incomingDirection == up {
			// | can connect from |, L and F
			return currentSymbol == "|" || currentSymbol == "L" || currentSymbol == "J" || currentSymbol == "S"
		}
		if incomingDirection == down {
			// | can connect from 7 and F
			return currentSymbol == "|" || currentSymbol == "7" || currentSymbol == "F" || currentSymbol == "S"
		}
		return false
	case "-":
		if incomingDirection == right {
			// - can connect from -, L and F
			return currentSymbol == "-" || currentSymbol == "L" || currentSymbol == "F" || currentSymbol == "S"
		}
		if incomingDirection == left {
			// - can connect from -, 7 and J
			return currentSymbol == "-" || currentSymbol == "7" || currentSymbol == "J" || currentSymbol == "S"
		}
		return false
	case "L":
		if incomingDirection == down {
			// L can connect from |,  F , 7
			return currentSymbol == "|" || currentSymbol == "F" || currentSymbol == "7" || currentSymbol == "S"
		}
		if incomingDirection == left {
			// L can connect from - , 7 , J
			return currentSymbol == "-" || currentSymbol == "7" || currentSymbol == "J" || currentSymbol == "S"
		}
		return false
	case "7":
		if incomingDirection == up {
			// 7 can connect from |, L, J
			return currentSymbol == "|" || currentSymbol == "L" || currentSymbol == "J" || currentSymbol == "S"
		}
		if incomingDirection == right {
			// 7 can connect from - , L , F
			return currentSymbol == "-" || currentSymbol == "L" || currentSymbol == "F" || currentSymbol == "S"
		}
		return false
	case "F":
		if incomingDirection == up {
			// F can connect from |, L, J
			return currentSymbol == "|" || currentSymbol == "L" || currentSymbol == "J" || currentSymbol == "S"
		}
		if incomingDirection == left {
			// F can connect from - , 7, J
			return currentSymbol == "-" || currentSymbol == "7" || currentSymbol == "J" || currentSymbol == "S"
		}
		return false
	case "J":
		if incomingDirection == down {
			// J can connect from |, F, 7
			return currentSymbol == "|" || currentSymbol == "F" || currentSymbol == "7" || currentSymbol == "S"
		}
		if incomingDirection == right {
			// J can connect from - , F, L
			return currentSymbol == "-" || currentSymbol == "F" || currentSymbol == "L" || currentSymbol == "S"
		}
		return false
	default:
		return false
	}
}

func determineStartingSymbol(matrix [][]string, startingPoint [2]int) string {

	if canConnect(matrix, startingPoint, down) && canConnect(matrix, startingPoint, up) {
		return "|"
	}
	// if we can connect left and right we are -
	if canConnect(matrix, startingPoint, left) && canConnect(matrix, startingPoint, right) {
		return "-"
	}
	// if we can connect up and left we are J
	if canConnect(matrix, startingPoint, left) && canConnect(matrix, startingPoint, up) {
		return "J"
	}
	// if we can connect down and left we are 7
	if canConnect(matrix, startingPoint, left) && canConnect(matrix, startingPoint, down) {
		return "7"
	}
	// if we can connect down and right we are F
	if canConnect(matrix, startingPoint, right) && canConnect(matrix, startingPoint, down) {
		return "F"
	}
	// if we can connect up and right we are L
	if canConnect(matrix, startingPoint, right) && canConnect(matrix, startingPoint, up) {
		return "L"
	}
	// not ideal but we need to return something
	return ""
}

func solution(filename string) int {
	lines := utils.FileToArray(filename)

	matrix := make([][]string, len(lines))

	startingPoint := [2]int{}
	for i, line := range lines {
		matrix[i] = make([]string, len(line))
		for j, char := range line {
			matrix[i][j] = string(char)
			if string(char) == "S" {
				startingPoint[0] = i
				startingPoint[1] = j
			}
		}
	}

	// we need to replace S with correct symbol based on what's next to it or it breaks things
	matrix[startingPoint[0]][startingPoint[1]] = determineStartingSymbol(matrix, startingPoint)
	fmt.Println("Starting at", startingPoint, matrix[startingPoint[0]][startingPoint[1]])

	reachedStartingPoint := false
	// we start at startingPoint and check all directions
	// as soon as we found a path we go that way and continue until we reach startingPoint again
	// we count steps and return half of it to get the max distance
	steps := 0
	current := [2]int{startingPoint[0], startingPoint[1]}

	var directionWeComeFrom Direction
	for !reachedStartingPoint {
		// fmt.Println("Steps taken", steps, "Current", current, matrix[current[0]][current[1]])
		// fmt.Println("checking direction")
		for _, direction := range directions {
			// we don't check direction we came from
			if direction[0] == directionWeComeFrom[0] && direction[1] == directionWeComeFrom[1] {
				continue
			}

			canConnectNext := canConnect(matrix, current, direction)
			// fmt.Println("Steps taken", steps, "Current", current, matrix[current[0]][current[1]], "Checking direction", direction, "next item", i, j, matrix[i][j], can)
			// we loop until we find a valid path
			if !canConnectNext {
				continue
			}
			// increase steps
			steps++
			// save where we came from so we don't go back
			directionWeComeFrom = [2]int{-direction[0], -direction[1]}
			// update current coordinates and exit loop
			current[0] = current[0] + direction[0]
			current[1] = current[1] + direction[1]

			break
		}
		// fmt.Println("After checking steps taken", steps, "Current", current, matrix[current[0]][current[1]])
		if current[0] == startingPoint[0] && current[1] == startingPoint[1] {
			reachedStartingPoint = true
		}
	}
	return steps / 2
}

func solution2(filename string) int {
	lines := utils.FileToArray(filename)

	matrix := make([][]string, len(lines))

	startingPoint := [2]int{}
	for i, line := range lines {
		matrix[i] = make([]string, len(line))
		for j, char := range line {
			matrix[i][j] = string(char)
			if string(char) == "S" {
				startingPoint[0] = i
				startingPoint[1] = j
			}
		}
	}

	// we need to replace S with correct symbol based on what's next to it or it breaks things
	matrix[startingPoint[0]][startingPoint[1]] = determineStartingSymbol(matrix, startingPoint)
	fmt.Println("Starting at", startingPoint, matrix[startingPoint[0]][startingPoint[1]])

	reachedStartingPoint := false
	// we start at startingPoint and check all directions
	// as soon as we found a path we go that way and continue until we reach startingPoint again
	// we count steps and return half of it to get the max distance
	steps := 0
	current := [2]int{startingPoint[0], startingPoint[1]}

	var pipePoints [][2]int
	var directionWeComeFrom Direction
	for !reachedStartingPoint {
		// fmt.Println("Steps taken", steps, "Current", current, matrix[current[0]][current[1]])
		// fmt.Println("checking direction")
		for _, direction := range directions {

			// we don't check direction we came from
			if direction[0] == directionWeComeFrom[0] && direction[1] == directionWeComeFrom[1] {
				continue
			}

			canConnectNext := canConnect(matrix, current, direction)
			// fmt.Println("Steps taken", steps, "Current", current, matrix[current[0]][current[1]], "Checking direction", direction, "next item", i, j, matrix[i][j], can)
			// we loop until we find a valid path
			if !canConnectNext {
				continue
			}
			// increase steps
			steps++
			// save where we came from so we don't go back
			directionWeComeFrom = [2]int{-direction[0], -direction[1]}
			// update current coordinates and exit loop
			current[0] = current[0] + direction[0]
			current[1] = current[1] + direction[1]

			pipePoints = append(pipePoints, current)
			break
		}

		// fmt.Println("After checking steps taken", steps, "Current", current, matrix[current[0]][current[1]])
		if current[0] == startingPoint[0] && current[1] == startingPoint[1] {
			reachedStartingPoint = true
		}
	}

	resultMap := make(map[int][]int)

	// Iterate through the original array and organize elements in the map
	for _, item := range pipePoints {
		firstInt := item[0]
		resultMap[firstInt] = append(resultMap[firstInt], item[1])
	}

	var area = 0
	for i := 0; i < len(matrix); i++ {
		if len(resultMap[i]) == 0 {
			continue
		}
		row := resultMap[i]

		sort.Slice(row, func(i, j int) bool {
			return row[i] < row[j]
		})

		encountered := 0

		// fmt.Println("row", i, row)
		for j := 0; j < len(row)-1; j += 1 {

			char := matrix[i][row[j]]
			if char == "|" {
				encountered++
				continue
			}
			if char == "─" {
				continue
			}

			next := matrix[i][row[j+1]]

			if char == "F" && next == "J" {
				encountered++
				continue
			}
			if char == "L" && next == "7" {
				encountered++
				continue
			}

			d := row[j+1] - row[j]

			if d > 1 {
				if encountered%2 != 0 {
					area += (d - 1)
				}
			}
		}

	}

	// fmt.Println("-- final area", area)
	return area
}

func matrixItem(matrix [][]string, coordinates [2]int) string {
	return matrix[coordinates[0]][coordinates[1]]
}

//493 is too high
