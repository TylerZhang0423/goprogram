package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i:=1; i<=num/2; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
