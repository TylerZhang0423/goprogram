package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "2", "C"}
	fmt.Println(calPoints(ops))
}

func calPoints(ops []string) int {
	var numSlice []int
	var sum = 0
	for _, item := range ops{
		index := len(numSlice) - 1
		if item == "C" {
			sum = sum - numSlice[index]
			numSlice = numSlice[:len(numSlice)-1]
		} else if item == "D" {
			numSlice = append(numSlice, numSlice[index]*2)
			sum = sum + numSlice[index+1]
		} else if item == "+" {
			numSlice = append(numSlice, numSlice[index]+numSlice[index-1])
			sum = sum + numSlice[index+1]
		} else {
			num, _ := strconv.Atoi(item)
			numSlice = append(numSlice, num)
			sum = sum + num
		}
	}
	return sum
}
