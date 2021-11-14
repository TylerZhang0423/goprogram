package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("AAB"))
}

func titleToNumber(columnTitle string) int {
	sum := 0
	length := len(columnTitle)
	for i := 0; i < length; i++ {
		sum = sum + int(float64(columnTitle[i] - 64) * math.Pow(26, float64(length-1-i)))
	}
	return sum
}