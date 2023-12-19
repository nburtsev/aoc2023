package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input_test.txt")
	re := regexp.MustCompile(`(.) (.*) \(#(.*)(.)\)`)
	delta := map[string]image.Point{
		"R": {1, 0}, "D": {0, 1}, "L": {-1, 0}, "U": {0, -1},
		"0": {1, 0}, "1": {0, 1}, "2": {-1, 0}, "3": {0, -1},
	}

	// thank you reddit, I was stuck on this for a while
	// now I know about shoelaces :)
	run := func(directionMatchIndex, digLengthMatchIndex, base int) int {
		dig, a2 := image.Point{0, 0}, 0
		// we go through all things that match the regexp
		for _, m := range re.FindAllStringSubmatch(string(input), -1) {
			// we parse the length using base 10 or 16
			l, _ := strconv.ParseInt(m[digLengthMatchIndex], base, strconv.IntSize)
			// we add the delta to the dig multiplied by the length
			n := dig.Add(delta[m[directionMatchIndex]].Mul(int(l)))
			fmt.Println(dig, n)
			fmt.Println(dig.X*n.Y-dig.Y*n.X, int(l))
			// we add the cross product of the two points to the area
			a2 += dig.X*n.Y - dig.Y*n.X + int(l)
			dig = n
		}
		return a2/2 + 1
	}

	fmt.Println(run(1, 2, 10))
}

// type Cell struct {
// 	symbol string
// 	color  string
// }

// func solution1(filename string) int {

// 	lines := utils.FileToArray(filename)

// 	// naive bruteforce with much copying should do it I think

// 	field := [][]Cell{{{"#", "#FFFFF"}}}

// 	position := [2]int{0, 0}

// 	re := regexp.MustCompile(`(U|D|R|L) (\d+) \((.*)\)`)

// 	for _, line := range lines {
// 		// fmt.Println(line)
// 		m := re.FindStringSubmatch(line)
// 		direction := m[1]
// 		distance, _ := strconv.Atoi(m[2])
// 		color := m[3]
// 		field, position = step(field, position, direction, distance, color)
// 	}

// 	// printField(field)

// 	count := 0

// 	reg := regexp.MustCompile(`#+( +)?#+`)
// 	for _, line := range field {

// 		s := ""

// 		for _, cell := range line {
// 			if cell.symbol == "#" {
// 				count++
// 				s += cell.symbol
// 			} else {
// 				s += " "
// 			}

// 		}

// 		// fmt.Println(s)
// 		// fmt.Println(strings.TrimSpace(s))
// 		for _, m := range reg.FindAllStringSubmatch(s, -1) {
// 			for _, c := range m[1:] {
// 				count += len(c)
// 			}
// 		}
// 	}

// 	return count
// }

// // 47798 too high
// // 44056 too low

// func printField(field [][]Cell) {
// 	color := "#FFFFFF"
// 	for _, line := range field {
// 		for _, cell := range line {
// 			if cell.symbol == "" {
// 				fmt.Print(gchalk.WithHex(color).Bold(" "))
// 			} else {
// 				color = cell.color
// 				fmt.Print(gchalk.WithHex(color).Bold(cell.symbol))
// 			}

// 		}
// 		fmt.Print("\n")
// 	}
// }

// func step(field [][]Cell, position [2]int, direction string, distance int, color string) ([][]Cell, [2]int) {

// 	switch direction {
// 	case "U":
// 		start := position
// 		for i := 1; i <= distance; i++ {
// 			newY := start[0] - i
// 			if newY < 0 {
// 				newLine := [][]Cell{}
// 				newLine = append(newLine, make([]Cell, len(field[0])))
// 				field = append(newLine, field...)
// 				// since this new top line we add to Y 0
// 				field[0][start[1]] = Cell{"#", color}
// 				position = [2]int{0, start[1]}
// 				continue
// 			}

// 			field[newY][start[1]] = Cell{"#", color}
// 			position = [2]int{newY, start[1]}
// 		}

// 	case "D":

// 		start := position
// 		for i := 1; i <= distance; i++ {
// 			newY := start[0] + i
// 			if newY >= len(field) {
// 				newLine := [][]Cell{}
// 				newLine = append(newLine, make([]Cell, len(field[0])))
// 				field = append(field, newLine...)
// 			}
// 			field[newY][start[1]] = Cell{"#", color}
// 		}
// 		position = [2]int{position[0] + distance, position[1]}

// 	case "R":
// 		for i := 1; i <= distance; i++ {
// 			newX := position[1] + i
// 			if newX >= len(field[0]) {
// 				for k, line := range field {
// 					field[k] = append(line, Cell{})
// 				}
// 			}
// 			field[position[0]][newX] = Cell{"#", color}
// 		}
// 		position = [2]int{position[0], position[1] + distance}

// 	case "L":
// 		start := position
// 		for i := 1; i <= distance; i++ {
// 			newX := start[1] - i

// 			if newX < 0 {
// 				for k, line := range field {
// 					field[k] = append([]Cell{{}}, line...)
// 				}
// 				field[start[0]][0] = Cell{"#", color}
// 				position = [2]int{start[0], 0}
// 				continue
// 			}

// 			field[start[0]][newX] = Cell{"#", color}
// 			position = [2]int{start[0], newX}

// 		}
// 	}
// 	return field, position
// }

// func solution2(filename string) int {
// 	return 0
// }
