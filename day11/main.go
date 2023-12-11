package main

import (
	"strings"
	"utils"
)

func main() {

	result := solution("input.txt", 2)
	println(result)
	result = solution("input.txt", 1_000_000)
	println(result)
}

func solution(filename string, expansionTimes int) int {

	lines := utils.FileToArray(filename)

	matrix := [][]string{}

	for _, line := range lines {
		matrix = append(matrix, utils.StringToStringArray(line, ""))
	}
	// i am too lazy to do any of this in a more clever way
	matrix = ExpandMatrix(matrix)
	matrix = utils.TransposeMatrix(matrix)
	matrix = ExpandMatrix(matrix)
	matrix = utils.TransposeMatrix(matrix)

	points := [][]int{}

	xCoord := 0
	for _, row := range matrix {
		if row[0] == "X" {
			xCoord += expansionTimes - 1
		}
		yCoord := 0
		for _, value := range row {
			if value == "X" {
				yCoord += expansionTimes - 1
			}
			if value == "#" {
				points = append(points, []int{xCoord, yCoord})

			}
			yCoord += 1

		}
		xCoord += 1
	}

	distancesSum := 0
	for len(points) > 0 {
		point := points[0]
		points = points[1:]
		for _, v := range points {
			distancesSum += DistanceBetweenPoints(point[0], point[1], v[0], v[1])
		}
	}

	return distancesSum
}

func DistanceBetweenPoints(x1, y1, x2, y2 int) int {
	return utils.Abs(x2-x1) + utils.Abs(y2-y1)
}

func ExpandMatrix(matrix [][]string) [][]string {
	result := [][]string{}

	blankRow := utils.StringToStringArray(strings.Repeat("X", len(matrix[0])), "")

	for _, row := range matrix {
		empty := true
		for _, v := range row {
			if v == "#" {
				empty = false
				break
			}
		}
		if empty {
			result = append(result, blankRow)
		} else {
			result = append(result, row)
		}
	}

	return result
}
