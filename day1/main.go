package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var sum2 = 0
	var sum1 = 0
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum1 += processLine(scanner.Text())
		sum2 += processLine2(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum1, sum2)

}

func processLine(line string) int {
	var a, b int = -1, -1

	for _, c := range line {
		if c >= '0' && c <= '9' {
			if a == -1 {
				a = int(c - '0')
			}
			b = int(c - '0')
		}
	}

	return (a*10 + b)
}

var NUMBERS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func toNumber(s string) int {
	for k, v := range NUMBERS {
		if strings.HasPrefix(s, k) {
			return v
		}
	}
	return -1
}

func processLine2(line string) int {
	var a, b int = -1, -1

	for i, c := range line {
		if c >= '0' && c <= '9' {
			int_c := int(c - '0')
			if a == -1 {
				a = int_c
			}
			b = int_c
			continue
		}

		var t = toNumber(line[i:])
		if t != -1 {
			if a == -1 {
				a = t
			}
			b = t
		}

	}
	return a*10 + b
}
