package main

import (
	"fmt"
	"utils"
)

func main() {

	result := solution("input2.txt")
	println(result)

	result = solution2("input2.txt")
	println(result)

}

func processPattern(pattern []string, skipLine int) (int, int) {
	mirrorLine := lookForMirror(pattern, skipLine)
	if mirrorLine > 0 {
		return mirrorLine, (mirrorLine + 1) * 100
	}

	pattern = transposePattern(pattern)
	mirrorLine = lookForMirror(pattern, skipLine)
	if mirrorLine > 0 {
		return mirrorLine, mirrorLine + 1
	}
	return -2, 0
}

func solution(filename string) int {
	fmt.Println("solution1 -- ")
	lines := utils.FileToArray(filename)

	patterns := [][]string{}

	tmp := []string{}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, tmp)
			tmp = []string{}
			continue
		}
		tmp = append(tmp, line)
	}
	patterns = append(patterns, tmp)

	sum := 0
	for _, pattern := range patterns {

		_, patternResult := processPattern(pattern, -1)
		sum += patternResult
	}
	return sum
}

func solution2(filename string) int {
	fmt.Println("solution2 -- ")
	lines := utils.FileToArray(filename)

	patterns := [][]string{}

	tmp := []string{}
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, tmp)
			tmp = []string{}
			continue
		}
		tmp = append(tmp, line)
	}
	patterns = append(patterns, tmp)

	sum := 0
	for _, pattern := range patterns {

		_, patternResult := processPattern2(pattern, -1)
		sum += patternResult
	}
	return sum
}

func getLine(pattern []string, index int) string {
	if index < 0 || index >= len(pattern) {
		return ""
	}
	return pattern[index]
}

func compareLines(pattern []string, index, index2 int) bool {

	line1 := getLine(pattern, index)
	line2 := getLine(pattern, index2)

	if line1 == "" || line2 == "" {
		return true
	}
	return getLine(pattern, index) == getLine(pattern, index2)
}

func transposePattern(input []string) []string {
	length := len(input[0])
	transposed := make([]string, length)
	for i := 0; i < length; i++ {
		for _, str := range input {
			transposed[i] += string(str[i])
		}
	}
	return transposed
}

// func processPattern2(pattern []string, originalLine int) (int, int) {

// 	originalReflection, originalSolution := processPattern(pattern, -1)
// 	println("-- original reflection found at ", originalReflection)
// 	// fixedReflection := 0
// 	for i := 0; i < len(pattern); i++ {
// 		for j := 0; j < len(pattern[i]); j++ {
// 			// fmt.Println("----------------flipping", i, j, "-", string(pattern[i][j]))
// 			newPattern := flipCharacter(pattern, i, j)
// 			// f, _ := processPattern(newPattern, originalSolution)
// 			fixedReflection, _ := processPattern2(newPattern, originalSolution)

// 			// if i == 8 && j == 7 {
// 			// 	println(fixedReflection)
// 			// 	println("XXX", string(newPattern[i][j]))

// 			// 	println("XXX", string(transposePattern(newPattern)[j][i]))
// 			// }

// 			// if new solution same as old one we continue looking
// 			// if fixedReflection == originalReflection {
// 			// 	println("-- original reflection found at ", fixedReflection, i, j)
// 			// 	continue
// 			// }
// 			// if we found a solution we return it
// 			if fixedReflection > 0 {
// 				println("-- fixed reflection found at ", fixedReflection, i, j)
// 				return fixedReflection, i
// 			}
// 		}
// 	}

// 	// println("-- cannot find fixed reflection ")
// 	// println("-- original reflection found at ", originalReflection)
// 	// printPattern(pattern)
// 	// println("-------------------")
// 	// printPattern(transposePattern(pattern))

// 	return 0, -2
// }

func processPattern2(pattern []string, skipLine int) (int, int) {
	originalLine, originalScore := processPattern(pattern, -1)
	fmt.Println("original line", originalLine, "score", originalScore)
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			fmt.Println("----------------flipping", i, j, "-", string(pattern[i][j]))
			newPattern := flipCharacter(pattern, i, j)
			mirrorLine := lookForMirror(newPattern, originalLine)
			if mirrorLine > 0 {
				fmt.Println("found new horizontal", mirrorLine, "score", (mirrorLine+1)*100)
				printPattern(newPattern)
				return mirrorLine, (mirrorLine + 1) * 100
			}

			newPattern = transposePattern(newPattern)
			mirrorLine = lookForMirror(newPattern, originalLine)
			if mirrorLine > 0 {
				fmt.Println("found new vertical", mirrorLine, "score", (mirrorLine + 1))
				printPattern(newPattern)
				return mirrorLine, mirrorLine + 1
			}
		}
	}

	println("-- cannot find fixed reflection ")
	printPattern(pattern)
	println("-------------------")
	printPattern(transposePattern(pattern))
	return -2, 0
}

func printPattern(pattern []string) {
	for i := 0; i < len(pattern); i++ {
		println(i, pattern[i])
	}
}

func lookForMirror(pattern []string, skipLine int) int {
	mirror := false
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern); j++ {
			// when we are on the index, we are checking line between it and next index
			up := i - j
			down := i + 1 + j
			// if we reach end of pattern in either direction we won't find more mirrors
			if up < 0 || down >= len(pattern) {
				break
			}
			mirror = compareLines(pattern, up, down)
			if !mirror {
				break
			}
		}
		if mirror {
			return i
		}
	}
	return 0
}

func flipCharacter(pattern []string, i, j int) []string {
	result := []string{}
	for k, line := range pattern {
		if k == i {
			result = append(result, flipCharacterInLine(line, j))
		} else {
			result = append(result, line)
		}
	}
	return result
}

func flipCharacterInLine(line string, i int) string {
	if string(line[i]) == "#" {
		return line[:i] + "." + line[i+1:]
	}
	return line[:i] + "#" + line[i+1:]
}

// 22k too low
// 26392 wrong
// 40k too high
// 40692 wrong

// 33174 wrong

// 33614 maybe ?

// 14474 too low
// 9833 too low

// 21980 wrong
// 22147 wrong
