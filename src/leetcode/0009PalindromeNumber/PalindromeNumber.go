package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(171))
}

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0)  {
		return false
	}
	palindromeNum := 0
	input := x
	for x > 0 {
		temp := x % 10
		x = x / 10
		palindromeNum = palindromeNum * 10 + temp
	}
	return input == palindromeNum
}