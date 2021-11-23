package main

import "fmt"

func main() {
	nums := []int{3,2,3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	length := len(nums)
	for i:=0; i<length; i++ {
		_, exist := numMap[nums[i]]
		if exist {
			numMap[nums[i]] = numMap[nums[i]] + 1
			if numMap[nums[i]] > length/2 {
				return nums[i]
			}
		} else {
			numMap[nums[i]] = 1
		}
	}
	return nums[0]
}
