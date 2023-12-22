package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {

	fmt.Println("Solution1", solution1("input.txt"))
	fmt.Println("Solution2", solution2("input.txt"))

}

type Point struct {
	x, y, z int
}
type Brick struct {
	id    int
	start Point // x,y,z
	end   Point // x,y,z
}

var VERBOSE = false

// each level of the stack is Z coordinate with all bricks that start on that level	in a slice
type Stack map[int][]Brick

func (b *Brick) String() string {
	return fmt.Sprintf("%v -> %v", b.start, b.end)
}

func (b *Brick) IsSupporting(other Brick) bool {

	// if other brick is not directly above us we cannot support it
	if other.start.z-b.end.z != 1 {
		return false
	}

	x := (b.start.x <= other.end.x && b.end.x >= other.start.x)
	y := (b.start.y <= other.end.y && b.end.y >= other.start.y)

	return x && y

}

// to find if brick can move down, we need to check if any of the bricks below it are supporting it
// if they do not we can move it down one step and check again until we reach a brick that supports it

func (b *Brick) canMoveDown(stack Stack) bool {

	// bottom of the stack
	if b.start.z == 1 {
		return false
	}

	for z := 1; z < b.start.z; z++ {
		// for all bricks in all levels below us we check if they support us
		for _, brickBelow := range stack[z] {
			if brickBelow.IsSupporting(*b) {
				return false
			}
		}
	}

	return true
}

func (b *Brick) moveDown(stack Stack) {

	// remove from previous level
	for i, brick := range stack[b.start.z] {
		if brick == *b {
			stack[b.start.z] = append(stack[b.start.z][:i], stack[b.start.z][i+1:]...)
			break
		}
	}

	// reduce z coordinates
	b.start.z--
	b.end.z--

	// add to new level
	stack[b.start.z] = append(stack[b.start.z], *b)

}

func (b *Brick) canBeDisintegrated(stack Stack) bool {

	thisBrickSupports := []Brick{}
	// if brick does not support anything above it - it can be disintegrated
	if layer, ok := stack[b.end.z+1]; !ok {
		// nothing above it
		return true
	} else {
		for _, brick := range layer {
			if b.IsSupporting(brick) {
				thisBrickSupports = append(thisBrickSupports, brick)
			}
		}

		if len(thisBrickSupports) == 0 {
			return true // does not support anything above it
		}

		// for each brick that is supported by current brick, we need to look if any of the bricks below it support it
		for _, brickToCheck := range thisBrickSupports {
			// we start with each brick being supported by current brick
			needsSupport := true
			// we go through all levels below
			for z := 1; z < brickToCheck.start.z; z++ {
				// for all bricks in all levels
				for _, brickBelow := range stack[z] {
					// we look for a brick that supports brick we are checking
					if brickBelow == *b {
						continue // we already checked this
					}
					if brickBelow.IsSupporting(brickToCheck) {
						// if we found one we stop checking the layer
						// and we set needsSupport to false
						needsSupport = false
						break
					}
				}
				// if we found a brick that supports the brick we are checking we can stop checking other layers
				if !needsSupport {
					break
				}
			}
			// if after we checked all the layers this brick still needs support we can stop checking other bricks
			if needsSupport {
				return false
			}
		}

	}

	return true
}

func solution1(input string) int {
	lines := utils.FileToArray(input)

	stack := make(Stack, 0)

	highest := 0
	for i, line := range lines {

		b := utils.StringToIntArray(strings.Replace(line, "~", ",", -1), ",")

		brick := Brick{
			id:    i,
			start: Point{b[0], b[1], b[2]},
			end:   Point{b[3], b[4], b[5]},
		}

		if brick.start.z > highest {
			highest = brick.start.z
		}

		stack[brick.start.z] = append(stack[brick.start.z], brick)
	}

	for z := 1; z <= highest; z++ {
		tmp := make([]Brick, len(stack[z]))
		copy(tmp, stack[z])

		for _, brick := range tmp {
			for brick.canMoveDown(stack) {
				brick.moveDown(stack)
			}
		}
	}

	count := 0

	for z := 1; z <= highest; z++ {
		for _, brick := range stack[z] {
			if brick.canBeDisintegrated(stack) {
				count++
			}
		}
	}

	return count

}

// part 2 requires counting how many bricks will fall if we remove one brick for all bricks in the solution
// with small input bruteforce for the win
func solution2(input string) int {
	lines := utils.FileToArray(input)

	stack := make(Stack, 0)

	highest := 0
	for i, line := range lines {

		b := utils.StringToIntArray(strings.Replace(line, "~", ",", -1), ",")

		brick := Brick{
			id:    i,
			start: Point{b[0], b[1], b[2]},
			end:   Point{b[3], b[4], b[5]},
		}

		if brick.start.z > highest {
			highest = brick.start.z
		}

		stack[brick.start.z] = append(stack[brick.start.z], brick)
	}

	for z := 1; z <= highest; z++ {
		tmp := make([]Brick, len(stack[z]))
		copy(tmp, stack[z])

		for _, brick := range tmp {
			for brick.canMoveDown(stack) {
				brick.moveDown(stack)
			}
		}
	}

	count := 0

	for z := 1; z <= highest; z++ {
		for _, brick := range stack[z] {

			// we remove the brick from the stack
			newStack := copyStack(stack)
			newStack = removeBrick(newStack, brick)

			// we let all bricks fall down and count how many bricks fall
			counter := map[int]bool{}
			for z := 1; z <= highest; z++ {
				tmp := make([]Brick, len(newStack[z]))
				copy(tmp, newStack[z])

				for _, brick := range tmp {
					for brick.canMoveDown(newStack) {
						if _, ok := counter[brick.id]; !ok {
							counter[brick.id] = true

						}
						brick.moveDown(newStack)
					}
				}
			}

			count += len(counter)
		}
	}

	return count

}

func copyStack(stack Stack) Stack {
	newStack := make(Stack, 0)
	for z, bricks := range stack {
		newStack[z] = make([]Brick, len(bricks))
		copy(newStack[z], bricks)
	}
	return newStack
}

func removeBrick(stack Stack, brick Brick) Stack {

	for i, b := range stack[brick.start.z] {
		if brick == b {
			stack[brick.start.z] = append(stack[brick.start.z][:i], stack[brick.start.z][i+1:]...)
			break
		}
	}

	return stack
}
