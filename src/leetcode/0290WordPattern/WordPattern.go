package main

import (
	"fmt"
	"strings"
)

func main() {
	var pattern = "abba"
	var str = "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))
}

func wordPattern(pattern string, s string) bool {
	var patternList = strings.Split(s, " ")
	if len(pattern) != len(patternList) {
		return false
	}
	var pattern2StringMap = make(map[string]string)
	var string2PatternMap = make(map[string]string)
	for i := 0; i < len(pattern); i++ {
		value, exist := pattern2StringMap[string(pattern[i])]
		value1, exist1 := string2PatternMap[patternList[i]]
		if exist {
			if value != patternList[i] {
				return false
			}
		} else {
			pattern2StringMap[string(pattern[i])] = patternList[i]
		}
		if exist1 {
			if value1 != string(pattern[i]) {
				return false
			}
		} else {
			string2PatternMap[patternList[i]] = string(pattern[i])
		}
	}
	return true
}