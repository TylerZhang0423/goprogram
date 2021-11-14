package main

import "fmt"

func main() {
	nums := []int{2,7,11,15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value, exist := numMap[target - nums[i]]
		if exist {
			return []int{value, i}
		} else {
			numMap [nums[i]] = i
		}
	}
	return []int{}
}
