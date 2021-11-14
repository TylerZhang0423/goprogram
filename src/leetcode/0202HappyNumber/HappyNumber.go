package main

import "fmt"

func main() {
	fmt.Println(15)
}

func isHappy(n int) bool {
	numMap := make(map[int]int)
	numMap[2] = 2
	numMap[4] = 4
	numMap[16] = 16
	numMap[37] = 37
	numMap[58] = 58
	numMap[89] = 89
	numMap[145] = 145
	numMap[42] = 42
	for n != 1 {
		_, exist := numMap[n]
		if exist {
			return false
		}
		n = calSum(n)
	}
	return true
}

func calSum(n int) int {
	result := 0
	for n > 0 {
		result = result + (n % 10) * (n % 10)
		n = n / 10
	}
	return result
}
