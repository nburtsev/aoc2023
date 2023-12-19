package main

import (
	"fmt"
	"strconv"
	"utils"

	"github.com/gammazero/deque"
)

// I've never finished this, I'm pretty sure I need to fix the conditions for stopping a path
// but it's beyond me at the moment
func main() {

	result1 := solution1("input.txt")
	println(result1)
	result2 := solution2("input.txt")
	println(result2)

}

type Path struct {
	direction string  // "up", "down", "left", "right"
	x         int     // current position
	y         int     // current position
	loss      int     // accumulated loss
	visited   [][]int // list of visited points
}

type Field [][]int

func solution1(filename string) int {

	lines := utils.FileToArray(filename)

	field := make(Field, len(lines))

	maxLoss := 0
	for y := 0; y < len(lines); y++ {
		field[y] = make([]int, len(lines))
		for x := 0; x < len(lines[0]); x++ {
			cellValue, _ := strconv.Atoi(string(lines[y][x]))
			field[y][x] = cellValue
			maxLoss += cellValue
		}
	}

	var q deque.Deque[Path]

	q.PushBack(Path{"down", 0, 1, field[1][0], [][]int{{0, 0}, {0, 1}}})
	q.PushBack(Path{"down", 0, 2, field[1][0] + field[2][0], [][]int{{0, 0}, {0, 2}}})
	q.PushBack(Path{"down", 0, 3, field[1][0] + field[2][0] + field[3][0], [][]int{{0, 0}, {0, 3}}})
	q.PushBack(Path{"right", 1, 0, field[0][1], [][]int{{0, 0}, {1, 0}}})
	q.PushBack(Path{"right", 2, 0, field[0][1] + field[0][2], [][]int{{0, 0}, {2, 0}}})
	q.PushBack(Path{"right", 3, 0, field[0][1] + field[0][2] + field[0][3], [][]int{{0, 0}, {3, 0}}})

	// p := q.Front()
	// fmt.Println("X", p.x, "Y", p.y, p.loss, "going", p.direction, p.visited)
	// println("----------")
	// for q.Len() > 0 {
	// 	p := q.PopFront()
	// 	fmt.Println("Y", p.y, "X", p.x, p.loss, "going", p.direction, p.visited)
	// }
	finishedPaths := []Path{}
	outOfBoundsPaths := []Path{}
	alreadyVisitedPaths := []Path{}
	steps := 0
	for q.Len() > 0 {
		p := q.PopFront()

		// we take next path and step from it to all possible directions
		// for possible distances
		for _, direction := range nextDirections(p.direction) {
			for i := 1; i <= 3; i++ {

				np := stepPath(p, direction, i, field)

				if np.x == -1 && np.y == -1 {
					outOfBoundsPaths = append(outOfBoundsPaths, np)
					continue
				}

				if np.x == -2 && np.y == -2 {
					alreadyVisitedPaths = append(alreadyVisitedPaths, np)
				}
				// if we exceeded maxLoss we don't add it back to the queue
				if np.loss >= maxLoss {
					continue
				}

				// if we step out of bounds we don't add it back to the queue
				if np.x == len(field[0])-1 && np.y == len(field)-1 {
					finishedPaths = append(finishedPaths, np)
					continue
				}

				q.PushBack(np)
			}
		}
		steps++

		// fmt.Println("queue", q)
		// if steps > 10 {
		// 	break
		// }
		if steps%1000000 == 0 {
			fmt.Println("Queue size", q.Len())
			fmt.Println(q.Back())
			fmt.Println("Paths that finished", len(finishedPaths))
			fmt.Println("Paths out of bounds", len(outOfBoundsPaths))
			fmt.Println("Paths alreadyVisitedPaths", len(alreadyVisitedPaths))
		}
	}
	// for q.Len() > 0 {
	// 	p := q.PopFront()
	// 	fmt.Println("X", p.x, "Y", p.y, p.loss, "going", p.direction, p.visited)
	// }

	// fmt.Println("Finished paths", len(finishedPaths))
	// return paths
	return 0
}

func stepPath(p Path, direction string, distance int, field Field) Path {

	path := Path{direction: direction, x: p.x, y: p.y, loss: p.loss, visited: p.visited}
	switch direction {
	case "up":
		path.y -= distance
	case "down":
		path.y += distance
	case "left":
		path.x -= distance
	case "right":
		path.x += distance
	default:
		panic("Unknown direction")
	}

	if path.x < 0 || path.x >= len(field[0]) || path.y < 0 || path.y >= len(field) {
		return Path{"", -1, -1, 999999999999999999, [][]int{}}
	}

	switch direction {
	case "up":
		for i := 1; i <= distance; i++ {
			path.loss += field[path.y+i][path.x]
		}
	case "down":
		for i := 1; i <= distance; i++ {
			path.loss += field[path.y-i][path.x]
		}
	case "left":
		for i := 1; i <= distance; i++ {
			path.loss += field[path.y][path.x+i]
		}
	case "right":
		for i := 1; i <= distance; i++ {
			path.loss += field[path.y][path.x-i]
		}
	default:
		panic("Unknown direction")
	}

	for _, v := range p.visited {
		if v[0] == path.x && v[1] == path.y {
			return Path{"", -2, -2, 999999999999999999, [][]int{}}
		}
	}
	path.visited = append(path.visited, []int{path.x, path.y})
	return path
}

// func calculateLoss(field Field, paths []*Path) int {

// 	endPaths := []*Path{}
// 	for len(paths) > 0 {
// 		p, e := stepPaths(field, paths)

// 		paths = p

// 		endPaths = append(endPaths, e...)
// 	}

// 	minLoss := 999999999999999999
// 	for _, p := range endPaths {
// 		if p.loss < minLoss {
// 			minLoss = p.loss
// 		}
// 	}

// 	return minLoss
// }

// we step path in all available directions (e.g. same direction, left or right)

func nextDirections(direction string) []string {

	switch direction {
	case "up":
		return []string{"left", "right"}
	case "down":
		return []string{"left", "right"}
	case "left":
		return []string{"up", "down"}
	case "right":
		return []string{"up", "down"}
	}
	return []string{}
}

// // we do this for all possible distances
// // if we step out of bounds we stop the path
// func stepPaths(field Field, paths []*Path) ([]*Path, []*Path) {

// 	printPaths(paths)
// 	newPaths := []*Path{}
// 	endPaths := []*Path{}
// 	for _, path := range paths {
// 		for _, direction := range nextDirections(path.direction) {
// 			for i := 1; i <= 3; i++ {
// 				np := stepPath(field, i, direction, path)
// 				for _, p := range np {
// 					if p.x == len(field[0])-1 && p.y == len(field)-1 {
// 						endPaths = append(endPaths, p)
// 					} else {
// 						newPaths = append(newPaths, p)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return newPaths, endPaths
// }

// func stepPath(field Field, distance int, direction string, p *Path) []*Path {

// 	path := &Path{direction: direction, x: p.x, y: p.y, loss: p.loss}
// 	fmt.Println("Y", path.y, "X", path.x, path.loss, "going", direction, distance)
// 	switch direction {
// 	case "up":
// 		path.y -= distance
// 		// fmt.Println("After step", "Y", path.y, "X", path.x, path.loss, "going", direction, distance)
// 		if path.y < 0 {
// 			return []*Path{}
// 		}
// 		for i := 1; i <= distance; i++ {
// 			path.loss += field[path.y+i][path.x].loss
// 		}
// 	case "down":
// 		path.y = path.y + distance
// 		if path.y >= len(field) {
// 			return []*Path{}
// 		}
// 		for i := 1; i <= distance; i++ {
// 			path.loss += field[path.y-i][path.x].loss
// 		}
// 	case "left":
// 		path.x = path.x - distance
// 		if path.x >= len(field[0]) {
// 			return []*Path{}
// 		}
// 		for i := 1; i <= distance; i++ {
// 			path.loss += field[path.y][path.x+i].loss
// 		}
// 	case "right":
// 		path.x = path.x + distance
// 		if path.x >= len(field[0]) {
// 			return []*Path{}
// 		}
// 		for i := 1; i <= distance; i++ {
// 			path.loss += field[path.y][path.x-i].loss
// 		}
// 	}

// 	// println("After everything", "Y", path.y, "X", path.x, path.loss, "going", direction, distance)

// 	return []*Path{path}
// }

func solution2(filename string) int {

	lines := utils.FileToArray(filename)

	sum := 0
	for i := 0; i < len(lines); i++ {
		sum += 1
	}

	return sum

}
