package main

import "fmt"

func main() {
	fmt.Println(reverse(19))

}

func reverse(x int) int {
	isPositive := true
	if x < 0 {
		isPositive = false
		x = -x
	}
	reverseNum := 0
	for x > 0 {
		temp := x % 10
		x = x / 10
		reverseNum = reverseNum * 10 + temp
		if reverseNum > 2147483647 || reverseNum < -2147483647 {
			return 0
		}
	}
	if isPositive {
		return reverseNum
	} else {
		return -reverseNum
	}
}
