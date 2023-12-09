package main

import (
	"utils"
)
func main() {

	result := solution("input.txt")
	println(result[0])
	println(result[1])
}

func solution(filename string) []int {
	next_number := 0 
	previous_number := 0 
	lines:= utils.FileToArray(filename)
	for _, line := range lines {
		line := utils.StringToIntArray(line, " ")
		diffs := [][]int{}
		diffs = append(diffs, line) 

		diff_index := 1
		all_zero := false
		for all_zero == false {
			new_diff := []int{}
			previous_diff := diffs[diff_index-1]
			for i:=1; i<len(previous_diff); i++ {
				d:= previous_diff[i] - previous_diff[i-1]
				new_diff = append(new_diff, d)
			}

			diffs = append(diffs, new_diff)

			all_zero = true
			for _, v := range new_diff {
				if v != 0 {
					all_zero = false
				}
			}
			diff_index ++ 
		}

		next_prediction := 0
		previous_prediction := 0 
		for i:=len(diffs); i>0; i-- {
			current_diff:= diffs[i-1]
			new_last:= current_diff[len(current_diff)-1]
			new_first:= current_diff[0]
			previous_prediction = new_first - previous_prediction
			next_prediction += new_last
		}

		next_number += next_prediction
		previous_number += previous_prediction
	}
	return []int{next_number, previous_number}
}
