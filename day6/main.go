package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result1 := solution("input.txt")
	println(result1)

	result2 := solution2("input.txt")
	println(result2)
}

func solution(fileName string) int {

	input := readInput(fileName)

	durations, distances := input[0], input[1]

	var margin = 1

	for i, duration := range durations {

		distance := distances[i]
		var possible_durations = 0

		// I'm sure there is a math formula for this but I'm too lazy
		for d := 1; d <= duration; d++ {

			if ((duration - d) * d) > distance {
				possible_durations++
			}
		}

		margin = margin * possible_durations
	}

	return margin
}

func solution2(fileName string) int {

	input := readInput2(fileName)

	durations, distances := input[0], input[1]

	var margin = 1

	for i, duration := range durations {

		distance := distances[i]
		var possible_durations = 0

		// I'm sure there is a math formula for this but I'm too lazy
		for d := 1; d <= duration; d++ {

			if ((duration - d) * d) > distance {
				possible_durations++
			}
		}

		margin = margin * possible_durations
	}

	return margin
}

func readInput(fileName string) [][]int {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input = [][]int{{}, {}}
	var index = 0
	for scanner.Scan() {

		line := scanner.Text()
		var re = regexp.MustCompile(`(\d+)`)
		l := re.FindAllString(line, -1)

		for _, s := range l {
			i, _ := strconv.Atoi(s)
			input[index] = append(input[index], i)

		}
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func readInput2(fileName string) [][]int {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input = [][]int{{}, {}}
	var index = 0
	for scanner.Scan() {

		line := scanner.Text()

		line = strings.Replace(line, " ", "", -1)

		var re = regexp.MustCompile(`(\d+)`)
		l := re.FindAllString(line, -1)

		for _, s := range l {
			i, _ := strconv.Atoi(s)
			input[index] = append(input[index], i)

		}
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
