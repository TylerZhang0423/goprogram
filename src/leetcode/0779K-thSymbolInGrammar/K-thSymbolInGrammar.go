package main

import (
	"fmt"
)

func main() {
	fmt.Println(kthGrammar(4, 5))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if kthGrammar(n-1, (k+1)/2) == 1 {
		if k%2 == 0 {
			return 0
		} else {
			return 1
		}
	} else {
		if k%2 == 0 {
			return 1
		} else {
			return 0
		}
	}
}