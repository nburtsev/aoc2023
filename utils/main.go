package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
					t := b
					b = a % b
					a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
					result = LCM(result, integers[i])
	}

	return result
}


func FileToArray(filename string) []string {
	lines := []string{}

	file, err := os.Open(filename)
	if err != nil {
					panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
					lines = append(lines, scanner.Text())
	}

	return lines
}

func StringToIntArray(input string, separator string) []int {
	output := []int{}
	for _, v := range strings.Split(input, separator) {
		i, _ := strconv.Atoi(string(v))
		output = append(output, i)
	}
	return output
}

func SumIntArray(input []int) int {
	output := 0
	for _, v := range input {
		output += v
	}
	return output
}
