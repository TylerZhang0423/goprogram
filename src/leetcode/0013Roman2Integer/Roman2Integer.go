package main

import "fmt"

func main() {
	fmt.Println(romanToInt("XI"))
}

func romanToInt(s string) int {
	sum :=0
	letter2NumMap := make(map[string]int)
	letter2NumMap["I"] = 1
	letter2NumMap["V"] = 5
	letter2NumMap["X"] = 10
	letter2NumMap["L"] = 50
	letter2NumMap["C"] = 100
	letter2NumMap["D"] = 500
	letter2NumMap["M"] = 1000
	for i := 0 ; i < len(s)-1; i++ {
		letter := string(s[i])
		nextLetter := string(s[i+1])
		num := letter2NumMap[letter]
		nextNum := letter2NumMap[nextLetter]
		if num < nextNum {
			sum = sum - num
		} else {
			sum = sum + num
		}
	}
	letter := string(s[len(s)-1])
	sum = sum + letter2NumMap[letter]
	return sum
}
