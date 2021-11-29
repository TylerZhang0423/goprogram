package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbguc"))
}

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen > tLen {
		return false
	}
	start := 0
	for i:=0; i<sLen; i++ {
		sLetter := string(s[i])
		for j:=start; j<tLen; j++ {
			tLetter := string(t[j])
			if sLetter == tLetter {
				start = j + 1
				break
			} else if sLetter != tLetter && j == tLen-1{
				return false
			}
		}
		if start>tLen-1 && i<sLen-1 {
			return false
		}
	}
	return true
}
