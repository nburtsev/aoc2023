package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	sum1 := solution("input.txt")

	println(sum1)

	sum2 := solution2("input.txt")
	println(sum2)

}

func solution(fileName string) int {

	var sum = 0
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers = [][]int{}
	var symbols = []int{}
	var line = 0
	for scanner.Scan() {
		line, numbers, symbols = processLine(scanner.Text(), numbers, symbols)
		sum += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func processLine(line string, previous_numbers [][]int, previous_symbols []int) (int, [][]int, []int) {
	var sum = 0
	number_pattern := regexp.MustCompile(`\d+`)
	number_match := number_pattern.FindAllStringIndex(line, -1)
	symbol_pattern := regexp.MustCompile(`[^\d.]`)
	symbol_match := symbol_pattern.FindAllStringIndex(line, -1)

	var this_line_symbols = []int{}
	for _, s := range symbol_match {
		this_line_symbols = append(this_line_symbols, s[0])
	}

	var numbers = [][]int{}
	for _, n := range number_match {
		number, _ := strconv.Atoi(line[n[0]:n[1]])
		numbers = append(numbers, []int{number, n[0], n[1]})
	}

	tnumbers := append(previous_numbers, numbers...)
	tsymbols := append(this_line_symbols, previous_symbols...)

	for _, s := range tsymbols {
		for i, n := range tnumbers {
			v := nextToSymbolOrZero(n, s)
			if v > 0 {
				tnumbers[i][0] = 0
				sum += v
			}
		}
	}

	return sum, numbers, this_line_symbols
}

func nextToSymbolOrZero(number []int, symbol int) int {
	n, number_start, number_end := number[0], number[1], number[2]
	if number_start == symbol || number_start == symbol-1 || number_start == symbol+1 {
		return n
	}
	if number_end == symbol || number_end == symbol+1 {
		return n
	}
	return 0
}

func solution2(fileName string) int {

	var lines = []string{}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	symbol_pattern := regexp.MustCompile(`[^\d.]`)
	number_pattern := regexp.MustCompile(`\d+`)

	var power = 0

	for i, line := range lines {
		symbol_match := symbol_pattern.FindAllStringIndex(line, -1)

		for _, s := range symbol_match {
			symbol := line[s[0]:s[1]]
			if symbol != "*" {
				continue
			}

			var lines_to_process = []string{}
			if i == 0 {
				lines_to_process = append(lines_to_process, lines[i], lines[i+1])
			} else if i == len(lines)-1 {
				lines_to_process = append(lines_to_process, lines[i], lines[i-1])
			} else {
				lines_to_process = append(lines_to_process, lines[i], lines[i+1], lines[i-1])
			}

			var numbers = [][]int{}

			for _, l := range lines_to_process {
				number_match := number_pattern.FindAllStringIndex(l, -1)

				for _, n := range number_match {
					number, _ := strconv.Atoi(l[n[0]:n[1]])
					numbers = append(numbers, []int{number, n[0], n[1]})
				}

			}
			var next_to_star = [][]int{}
			for _, n := range numbers {
				v := nextToSymbolOrZero(n, s[0])
				if v > 0 {
					next_to_star = append(next_to_star, []int{v, n[1], n[2]})
				}
			}

			if len(next_to_star) == 2 {
				power += next_to_star[0][0] * next_to_star[1][0]
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return power
}
