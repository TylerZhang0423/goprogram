package main

import "fmt"

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	length := len(height)
	front := 0
	end := length-1
	area := 0
	for front < end {
		currentArea := (end - front) * minHeight(height[front], height[end])
		if currentArea > area {
			area = currentArea
		}
		if height[front] > height[end] {
			end = end - 1
		} else {
			front = front + 1
		}
	}
	return area
}

func minHeight(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
