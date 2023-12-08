package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

func main() {

	result1 := solution("input.txt")
	println(result1)

	result2 := solution3("input.txt")
	println(result2)
}

func solution(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	instructions = strings.Replace(instructions, "L", "0", -1)
	instructions = strings.Replace(instructions, "R", "1", -1)
	steps := map[string][]string{}

	pattern := regexp.MustCompile(`[A-Z]+`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllString(line, -1)
		steps[matches[0]] = matches[1:]
	}

	next:= 0
	next_location:= "AAA"
	steps_taken := 0
	for next<len(instructions) {
		direction, _ := strconv.Atoi(string(instructions[next]))
		next_steps := steps[next_location]
		next_location = next_steps[direction]
		steps_taken++
		next++
		if (next_location == "ZZZ") {
			break
		}
		if next == len(instructions) {
			next = 0
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return steps_taken
}

// people much smarter than me told me this will finish in weeks
func solution2(filename string) int {


	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	
	instructions = strings.Replace(instructions, "L", "0", -1)
	instructions = strings.Replace(instructions, "R", "1", -1)


	steps := map[string][]string{}
	pattern := regexp.MustCompile(`[0-9A-Z]{3,}`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllString(line, -1)
		steps[matches[0]] = matches[1:]
	}

	next:= 0
	steps_taken := 0

	next_locations := []string{}

	for k := range steps {
		if (strings.HasSuffix(k,"A")) {
			next_locations = append(next_locations, k)
		}
	}

	for next<len(instructions) {
		direction, _ := strconv.Atoi(string(instructions[next]))

		
		finished:= true 

		for i, v := range next_locations {
			next_step := steps[v]
			next_location := next_step[direction]
			next_locations[i] = next_location

			if !strings.HasSuffix(next_location, "Z") {
				finished = false
			}
		}

		steps_taken++
		next++
		if (steps_taken % 1000000 == 0) {
			fmt.Printf("\033[2K\r%d", steps_taken)
		}
		
		if finished {
			break
		}
		
		if next == len(instructions) {
			next = 0
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return steps_taken
}

func solution3(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	
	instructions = strings.Replace(instructions, "L", "0", -1)
	instructions = strings.Replace(instructions, "R", "1", -1)


	steps := map[string][]string{}
	pattern := regexp.MustCompile(`[0-9A-Z]{3,}`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllString(line, -1)
		steps[matches[0]] = matches[1:]
	}

	locations := []string{}
	
	for k := range steps {
		if (strings.HasSuffix(k,"A")) {
			locations = append(locations, k)
		}
	}

	lengths := []int{}
	for _, v := range locations {

		next:= 0
		steps_taken := 0

		for next<len(instructions) {
			direction, _ := strconv.Atoi(string(instructions[next]))
			next_step := steps[v]
			v = next_step[direction]
			steps_taken++
			next++
			if next == len(instructions) {
				next = 0
			}
			if strings.HasSuffix(v, "Z") {
				lengths = append(lengths, steps_taken)
				break
			}
		}
	}	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// test case
	if len(lengths) == 2 {
		return utils.LCM(lengths[0], lengths[1])
	}
	return utils.LCM(lengths[0], lengths[1], lengths[2:]...)

}
