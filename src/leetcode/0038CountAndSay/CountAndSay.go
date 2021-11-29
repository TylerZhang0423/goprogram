package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(10))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		return count(countAndSay(n-1))
	}
}

func count(s string) string {
	length := len(s)
	currentLetter := string(s[0])
	currentCount := 1
	result := ""
	for i:=1; i<length; i++ {
		if string(s[i]) == currentLetter {
			currentCount = currentCount + 1
		} else {
			result = result + strconv.Itoa(currentCount) + currentLetter
			currentLetter = string(s[i])
			currentCount = 1
		}
	}
	result = result + strconv.Itoa(currentCount) + currentLetter
	return result
}
