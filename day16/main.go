package main

import (
	"fmt"
	"utils"

	"github.com/fatih/color"
)

func main() {
	println(solution1("input.txt"))
	println(solution2("input.txt"))
}

// all these could be slices I guess but it is easier to read like this :)
type Cell struct {
	symbol  rune
	visited int // totally unnecessary, could be a bool, but I wanted to print different colors
}

type Field [][]Cell

type Beam struct {
	direction string // "up", "down", "left", "right" i'm sure I could use an enum or something, but go makes things complicated
	x         int
	y         int
}

func solution1(filename string) int {
	lines := utils.FileToArray(filename)

	field := make(Field, len(lines))

	for y := 0; y < len(lines); y++ {
		field[y] = make([]Cell, len(lines))
		for x := 0; x < len(lines[0]); x++ {
			field[y][x] = Cell{rune(lines[y][x]), 0}
		}
	}

	c, _ := energizedCells(field, []*Beam{{direction: "right", x: -1, y: 0}})
	// printFieldColor(f)
	return c
}

func solution2(filename string) int {
	lines := utils.FileToArray(filename)

	field := make(Field, len(lines))

	for y := 0; y < len(lines); y++ {
		field[y] = make([]Cell, len(lines))
		for x := 0; x < len(lines[0]); x++ {
			field[y][x] = Cell{rune(lines[y][x]), 0}
		}
	}

	max := 0
	// maxF := Field{}

	for y := 0; y < len(field); y++ {

		cpy := make(Field, len(field))
		copy(cpy, field)
		m1, _ := energizedCells(cpy, []*Beam{{direction: "right", x: -1, y: y}})
		if m1 > max {
			max = m1
			// maxF = f
		}
		copy(cpy, field)
		m2, _ := energizedCells(cpy, []*Beam{{direction: "left", x: len(field[0]), y: y}})
		if m2 > max {
			max = m2
			// maxF = f
		}
	}

	for x := 0; x < len(field[0]); x++ {
		cpy := make(Field, len(field))
		copy(cpy, field)
		m1, _ := energizedCells(cpy, []*Beam{{direction: "down", x: x, y: -1}})
		if m1 > max {
			max = m1
			// maxF = f
		}
		copy(cpy, field)
		m2, _ := energizedCells(cpy, []*Beam{{direction: "up", x: x, y: len(field)}})
		if m2 > max {
			max = m2
			// maxF = f
		}
	}

	// printFieldColor(maxF)
	return max
}

func energizedCells(input Field, beams []*Beam) (int, Field) {

	// I don't know why go makes this so unnecessarily complicated
	field := make(Field, len(input))
	for i := range input {
		field[i] = make([]Cell, len(input[i]))
		copy(field[i], input[i])
	}
	for len(beams) > 0 {

		newBeams := []*Beam{}
		for _, b := range beams {
			newBeams = append(newBeams, stepBeam(field, b)...)
		}

		beams = newBeams
	}

	count := 0
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			if field[y][x].visited > 0 {
				count++
			}
		}
	}

	return count, field
}

func stepBeam(field Field, b *Beam) []*Beam {

	switch b.direction {
	case "up":
		b.y--
	case "down":
		b.y++
	case "left":
		b.x--
	case "right":
		b.x++
	}
	// if we are out of bounds this beam is done
	if b.x < 0 || b.x >= len(field[0]) || b.y < 0 || b.y >= len(field) {
		return []*Beam{}
	}

	positionAfterStep := &field[b.y][b.x]

	positionAfterStep.visited += 1

	switch positionAfterStep.symbol {
	// we mark dot as energized and continue
	case '.':
		switch b.direction {
		case "up":
			positionAfterStep.symbol = '^'
		case "down":
			positionAfterStep.symbol = 'v'
		case "left":
			positionAfterStep.symbol = '<'
		case "right":
			positionAfterStep.symbol = '>'
		}
	// we either split the beam into 2 new ones or do nothing
	case '|':
		if b.direction == "left" || b.direction == "right" {
			return []*Beam{
				{direction: "up", x: b.x, y: b.y},
				{direction: "down", x: b.x, y: b.y},
			}
		}
	// we either split the beam into 2 new ones or do nothing
	case '-':
		if b.direction == "up" || b.direction == "down" {
			return []*Beam{
				{direction: "left", x: b.x, y: b.y},
				{direction: "right", x: b.x, y: b.y},
			}
		}
	// we turn current beam and continue
	case '/':
		switch b.direction {
		case "up":
			b.direction = "right"
		case "down":
			b.direction = "left"
		case "left":
			b.direction = "down"
		case "right":
			b.direction = "up"
		}

	case '\\':
		switch b.direction {
		case "up":
			b.direction = "left"
		case "down":
			b.direction = "right"
		case "left":
			b.direction = "up"
		case "right":
			b.direction = "down"
		}

	default: // we hit already energized empty cell, if it matches our direction we are done
		if (b.direction == "up" && positionAfterStep.symbol == '^') ||
			(b.direction == "down" && positionAfterStep.symbol == 'v') ||
			(b.direction == "left" && positionAfterStep.symbol == '<') ||
			(b.direction == "right" && positionAfterStep.symbol == '>') {
			return []*Beam{}
		}
	}

	return []*Beam{b}
}

func printFieldPlain(field Field) {

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			fmt.Print(string(field[y][x].symbol))
		}
		fmt.Println()
	}
}

func printFieldColors(field Field) {
	green := color.New(color.FgGreen)
	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[0]); x++ {
			if field[y][x].visited > 1 {
				if field[y][x].symbol == '^' || field[y][x].symbol == 'v' || field[y][x].symbol == '<' || field[y][x].symbol == '>' {
					red.Print(field[y][x].visited)
				} else {
					cyan.Print(string(field[y][x].symbol))
				}
			} else if field[y][x].visited > 0 {
				green.Print(string(field[y][x].symbol))
			} else {
				print(" ")
				// this would print unvisited cells
				// print(string(field[y][x].symbol))
			}
		}
		fmt.Println()
	}
}

func printBeam(beam *Beam) {
	fmt.Print("{", beam.x, beam.y, ", ", beam.direction, "} ")
}
