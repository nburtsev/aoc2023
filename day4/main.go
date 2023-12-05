package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

	for scanner.Scan() {
		sum += processLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func processLine(line string) int {

	var sum = 0
	game := strings.Split(line, ":")

	winning_numbers := strings.Split(strings.TrimSpace(strings.Split(game[1], "|")[0]), " ")
	game_numbers := strings.Split(strings.TrimSpace(strings.Split(game[1], "|")[1]), " ")

	var unique_game_numbers = []string{}
	m := map[string]bool{}

	for _, n := range game_numbers {
		if n == "" {
			continue
		}
		if !m[n] {
			m[n] = true
			unique_game_numbers = append(unique_game_numbers, n)
		}
	}

	var unique_winning_numbers = []string{}
	mw := map[string]bool{}

	for _, n := range winning_numbers {
		if n == "" {
			continue
		}
		if !mw[n] {
			mw[n] = true
			unique_winning_numbers = append(unique_winning_numbers, n)
		}
	}

	for i, n := range unique_winning_numbers {
		for _, g := range unique_game_numbers {
			if n == g {
				if sum == 0 {
					sum = 1
				} else {
					sum *= 2
				}
				winning_numbers[i] = ""
			}
		}
	}

	return sum
}

func incMap(m map[int]int, n int, v int) map[int]int {

	_, ok := m[n]

	if ok {
		m[n] += v
	} else {
		m[n] = v
	}
	return m
}

func solution2(fileName string) int {

	lines := map[int]int{}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		id, matches := processLine2(scanner.Text())
		lines = incMap(lines, id, 1)

		v, ok := lines[id]

		var cards = 1
		if ok {
			cards = v
		}

		for i := range make([]int, matches) {
			lines = incMap(lines, id+i+1, cards)
		}

	}

	var sum = 0
	for _, v := range lines {

		sum += v
	}

	return sum
}

func processLine2(line string) (int, int) {

	var sum = 0
	game := strings.Split(line, ":")

	game_id, _ := strconv.Atoi(strings.Split(game[0], " ")[len(strings.Split(game[0], " "))-1])
	winning_numbers := strings.Split(strings.TrimSpace(strings.Split(game[1], "|")[0]), " ")
	game_numbers := strings.Split(strings.TrimSpace(strings.Split(game[1], "|")[1]), " ")

	var unique_game_numbers = []string{}
	m := map[string]bool{}

	for _, n := range game_numbers {
		if n == "" {
			continue
		}
		if !m[n] {
			m[n] = true
			unique_game_numbers = append(unique_game_numbers, n)
		}
	}

	var unique_winning_numbers = []string{}
	mw := map[string]bool{}

	for _, n := range winning_numbers {
		if n == "" {
			continue
		}
		if !mw[n] {
			mw[n] = true
			unique_winning_numbers = append(unique_winning_numbers, n)
		}
	}

	// fmt.Println(unique_game_numbers)
	// fmt.Println(unique_winning_numbers)

	for _, n := range unique_winning_numbers {
		for _, g := range unique_game_numbers {
			if n == g {
				sum += 1
			}
		}
	}

	return game_id, sum
}
