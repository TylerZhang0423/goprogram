package main

import "fmt"

func main() {
	var n = 5
	fmt.Println(arrangeCoins(n))
}

func arrangeCoins(n int) int {
	var total = n
	var index = 1
	for total >= index {
		total = total - index
		index = index + 1
	}
	return index - 1
}
