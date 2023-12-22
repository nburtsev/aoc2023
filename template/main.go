package main

import (
	"fmt"
	"utils"
)

func main() {

	fmt.Println("Solution1", solution1("input.txt"))
	fmt.Println("Solution2", solution2("input.txt"))

}

func solution1(input string) int {
	lines := utils.FileToArray(input)
	return len(lines)

}

func solution2(input string) int {
	lines := utils.FileToArray(input)
	return len(lines)
}
