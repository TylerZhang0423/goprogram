package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aabbccd"))
}

func longestPalindrome(s string) int {
	letterMap := make(map[string]int)
	flag := 0
	result := 0
	for i:=0; i<len(s); i++ {
		letter := string(s[i])
		value, exist := letterMap[letter]
		if exist {
			letterMap[letter] = value + 1
		} else {
			letterMap[letter] = 1
		}
	}
	for letter := range letterMap {
		if letterMap[letter] % 2 == 0 {
			result = result + letterMap[letter]
		} else if letterMap[letter] % 2 != 0 && letterMap[letter] > 1{
			result = result + letterMap[letter] - 1
			flag = 1
		} else {
			flag = 1
		}
	}
	return result + flag
}
