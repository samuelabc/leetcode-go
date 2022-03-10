package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//double loop
// func maxArea(height []int) int {
// 	maxArea := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			smallerHeight := min(height[i], height[j])
// 			distance := j - i
// 			area := distance * smallerHeight
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 	}
// 	return maxArea
// }

//double pointer
func maxArea(height []int) int {
	maxArea := 0
	start := 0
	end := len(height) - 1
	for start < end {
		smallerHeight := 0
		distance := end - start
		if height[start] < height[end] {
			smallerHeight = height[start]
			start++
		} else {
			smallerHeight = height[end]
			end--
		}
		area := smallerHeight * distance
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	res := maxArea(height)
	fmt.Println(res)
}
