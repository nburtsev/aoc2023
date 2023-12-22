package utils

import (
	"bufio"
	"fmt"
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

func StringToStringArray(input string, separator string) []string {
	output := []string{}
	for _, v := range strings.Split(input, separator) {
		output = append(output, string(v))
	}
	return output
}

func TransposeMatrix(matrix [][]string) [][]string {
	result := [][]string{}
	for colIndex := range matrix[0] {
		row := []string{}
		for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
			row = append(row, matrix[rowIndex][colIndex])
		}
		result = append(result, row)
	}
	return result
}

func PrintMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("------------------")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
