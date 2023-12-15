package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

func main() {

	result1 := solution1("input.txt")
	println(result1)
	result2 := solution2("input.txt")
	println(result2)

}

func solution1(filename string) int {
	lines := utils.FileToArray(filename)

	steps := strings.Split(lines[0], ",")

	sum := 0
	for i := 0; i < len(steps); i++ {
		sum += hash(steps[i])
	}

	return sum
}

func solution2(filename string) int {

	lines := utils.FileToArray(filename)

	steps := strings.Split(lines[0], ",")

	boxes := make([]Lense, 256)
	for i := 0; i < len(steps); i++ {

		step := parseStep(steps[i])
		if step.operation == "=" {
			addLense(&boxes[step.box], step.label, step.focalLength)
		} else {
			removeLense(&boxes[step.box], step.label)
		}
	}

	totalPower := 0
	for i := 0; i < len(boxes); i++ {
		if boxes[i].label != "" {
			slot := 1
			for box := &boxes[i]; box != nil; box = box.next {
				fp := (i + 1) * slot * box.focalLength
				totalPower += fp
				slot++
			}
		}
	}

	return totalPower

}

func printBox(box *Lense) {
	if box.label != "" {
		for box := box; box != nil; box = box.next {
			fmt.Printf(" -> [%s %d]", box.label, box.focalLength)
		}
		fmt.Print("\n")
	}
}

func printBoxes(boxes []Lense) {
	for i := 0; i < len(boxes); i++ {
		if boxes[i].label != "" {
			fmt.Print(i)
			for box := &boxes[i]; box != nil; box = box.next {
				fmt.Printf(" -> [%s %d]", box.label, box.focalLength)
			}
			fmt.Print("\n")
		}
	}
	println("____________________________")
}

// linked list of lenses
type Lense struct {
	label       string
	focalLength int
	next        *Lense
}

type Step struct {
	label       string
	box         int
	operation   string
	focalLength int
}

func parseStep(step string) Step {

	re := regexp.MustCompile(`(\w+)(=|-)(\d+)?`)

	components := re.FindAllStringSubmatch(step, -1)

	focalLength := 0

	if len(components[0]) > 3 {
		focalLength, _ = strconv.Atoi(components[0][3])
	}
	return Step{
		label:       components[0][1],
		box:         hash(components[0][1]),
		operation:   components[0][2],
		focalLength: focalLength,
	}
}

func addLense(lense *Lense, label string, focalLength int) {

	// empty box
	if lense.label == "" {
		lense.label = label
		lense.focalLength = focalLength
		return
	}
	// go through boxes until we find the label
	for lense != nil {
		if lense.label == label {
			lense.focalLength = focalLength
			return
		}
		// we stop on the last one
		if lense.next == nil {
			break
		}
		lense = lense.next
	}
	// add new lense
	lense.next = &Lense{
		label:       label,
		focalLength: focalLength,
	}
}

func removeLense(lense *Lense, label string) {
	// empty box
	if lense.label == "" {
		return
	}
	// one lense in the box we delete it
	if lense.label == label && lense.next == nil {
		lense.label = ""
		return
	}
	prev := lense
	// go through boxes until we find the label
	for lense != nil {
		if lense.label == label {
			// when we find the label, we replace it with the next item if it exists
			// and if it does not exist we need to set previous next to nil
			if lense.next != nil {
				lense.label = lense.next.label
				lense.focalLength = lense.next.focalLength
				lense.next = lense.next.next
			} else {
				prev.next = nil
				// lense = nil
			}
			return
		}
		prev = lense
		lense = lense.next
	}
}

func hash(s string) int {
	hash := 0
	for _, s := range s {
		hash = hash + int(s)
		hash = hash * 17
		hash = hash % 256
	}
	return hash
}
