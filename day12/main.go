package main

import (
	"fmt"
	"regexp"
	"strings"
	"utils"
)

func main() {

	result := solution("input.txt")
	println(result)
	result = solution2("input.txt")
	println(result)

}

func solution(filename string) int {
	lines := utils.FileToArray(filename)
	combinations := 0
	for _, line := range lines {
		springs, groups := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		groupsArray := utils.StringToIntArray(groups, ",")
		combinations += step(springs, groupsArray, map[string]int{})
	}
	return combinations
}

func solution2(filename string) int {
	lines := utils.FileToArray(filename)
	combinations := 0
	for _, line := range lines {
		inputSprings, inputGroups := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		groupsArray := utils.StringToIntArray(inputGroups, ",")
		groups := []int{}
		springs := ""
		for i := 0; i < 5; i++ {
			groups = append(groups, groupsArray...)
		}

		for i := 0; i < 5; i++ {
			if i > 0 {
				springs += "?"
			}
			springs += inputSprings
		}
		combinations += step(springs, groups, map[string]int{})
	}
	return combinations
}

func step(spring string, groups []int, paths map[string]int) int {

	key := spring + fmt.Sprintf("%v", groups)

	if val, ok := paths[key]; ok {
		return val
	}

	paths[key] = 0

	valid, canContinue := true, true

	badGearsCount := strings.Count(spring, "#")

	// if we ran out of groups but still have bad gears it's not a valid variant but we can continue
	if badGearsCount > 0 && len(groups) == 0 {
		valid = false
		canContinue = true
	} else {

		re := regexp.MustCompile(`\#+`)
		badGears := re.FindAllString(spring, -1)
		noMoreQuestionMarks := strings.Count(spring, "?") == 0

		// if we have more groups than remaining bad gears, we can't add a group
		// if we don't have more question marks we can't continue
		if len(badGears) != len(groups) {
			valid = false
			canContinue = noMoreQuestionMarks
		} else {
			// if remaining groups don't fit in remaining bad gears, we can't add a group
			// if we don't have more question marks we can't continue
			for i, group := range groups {
				if group != len(badGears[i]) {
					valid = false
					canContinue = noMoreQuestionMarks
				}
			}
		}
	}

	if canContinue {
		if valid {
			paths[key] = 1
			return 1
		}
		return 0
	}

	if spring[0] == '.' {
		n := step(spring[1:], groups, paths)
		paths[key] = n
		return n
	}
	n := 0

	if spring[0] == '?' {
		n += step(spring[1:], groups, paths)
	}

	// check if group fits the string
	count := 0
	group := groups[0]

	// if string is shorter than group size it's not a valid variant
	if len(spring) >= group {

		// check if we have dots where group should be
		hasDots := false
		for i := 0; i < group; i++ {
			if spring[i] == '.' {
				hasDots = true
				break
			}
		}
		if !hasDots {
			// if group matches string then we got a match
			if len(spring) == group {
				count = group
			} else {
				// otherwise we check next character
				// if it is dot or question mark we skip it
				if spring[group] == '.' || spring[group] == '?' {
					count = group + 1
				}
			}
		}
	}

	if count == 0 {
		paths[key] = n
		return n
	}

	n += step(spring[count:], groups[1:], paths)
	paths[key] = n
	return n
}
