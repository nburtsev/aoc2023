package main

import (
	"utils"
)

const TIMES = 1_000_000_000

func main() {

	result1 := solution1("input.txt")
	println(result1)
	result2 := solution2("input.txt", TIMES)
	println(result2)

}

func solution1(filename string) int {

	lines := utils.FileToArray(filename)

	m := [][]string{}

	for line := range lines {
		m = append(m, utils.StringToStringArray(lines[line], ""))
	}
	m = rollNorth(m)

	totalLoad := calculateNorthLoad(m)

	return totalLoad
}

func solution2(filename string, times int) int {

	lines := utils.FileToArray(filename)

	m := [][]string{}

	stateMap := map[string]int{}

	for line := range lines {
		m = append(m, utils.StringToStringArray(lines[line], ""))
	}

	// breakpoint for the loop below
	b := times
	for i := 0; i < times; i++ {
		// cycle the map
		m = cycleField(m)
		// join it so we can use it as key
		h := joinMatrix(m)
		// if we haven't changed the breakpoint yet
		if b == times {
			// check if we have seen this state before
			if prev, ok := stateMap[h]; ok {
				// if we have, we can calculate the breakpoint
				// current step + remaining steps % cycle_length - 1
				b = i + (times-prev)%(i-prev) - 1
			} else {
				stateMap[h] = i
			}
		}

		if i == b {
			println("Cycle start", stateMap[h], "cycle length", i-stateMap[h])
			break
		}

	}

	totalLoad := calculateNorthLoad(m)

	return totalLoad
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

func cycleField(m [][]string) [][]string {

	m = rollNorth(m)
	m = rollWest(m)
	m = rollSouth(m)
	m = rollEast(m)

	return m
}

func rollNorth(m [][]string) [][]string {
	// go until we reach O
	// start swapping it with previous item until it is #, O or we reach border
	for j := 0; j < len(m[0]); j++ {
		for i := 0; i < len(m); i++ {

			object := m[i][j]
			if object == "." || object == "#" {
				continue
			}

			for k := i - 1; k >= 0; k-- {
				if m[k][j] == "#" || m[k][j] == "O" {
					break
				}
				m[k][j], m[k+1][j] = m[k+1][j], m[k][j]
			}
		}
	}
	return m
}

func rollSouth(m [][]string) [][]string {
	// go until we reach O
	// start swapping it with next item until it is #, O or we reach border
	for j := len(m[0]) - 1; j >= 0; j-- {
		for i := len(m) - 1; i >= 0; i-- {

			object := m[i][j]
			if object == "." || object == "#" {
				continue
			}

			for k := i + 1; k < len(m); k++ {
				if m[k][j] == "#" || m[k][j] == "O" {
					break
				}
				m[k][j], m[k-1][j] = m[k-1][j], m[k][j]
			}
		}
	}
	return m
}

func rollEast(m [][]string) [][]string {
	// go until we reach O
	// start swapping it with next item until it is #, O or we reach border
	for i := len(m) - 1; i >= 0; i-- {
		for j := len(m[0]) - 1; j >= 0; j-- {

			object := m[i][j]
			if object == "." || object == "#" {
				continue
			}

			for k := j + 1; k < len(m[0]); k++ {
				if m[i][k] == "#" || m[i][k] == "O" {
					break
				}
				m[i][k], m[i][k-1] = m[i][k-1], m[i][k]
			}
		}
	}
	return m
}

func rollWest(m [][]string) [][]string {
	// go until we reach O
	// start swapping it with previous item until it is #, O or we reach border

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			object := m[i][j]
			if object == "." || object == "#" {
				continue
			}

			for k := j - 1; k >= 0; k-- {
				if m[i][k] == "#" || m[i][k] == "O" {
					break
				}
				m[i][k], m[i][k+1] = m[i][k+1], m[i][k]
			}
		}
	}

	return m
}

func calculateNorthLoad(m [][]string) int {
	totalLoad := 0
	max := len(m)
	for i, row := range m {
		sumRow := 0
		for _, item := range row {
			if item == "O" {
				sumRow++
			}
		}
		totalLoad += sumRow * (max - i)
	}
	return totalLoad
}
