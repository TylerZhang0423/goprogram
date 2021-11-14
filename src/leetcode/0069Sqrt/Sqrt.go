package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(14))
}

func mySqrt(x int) int {
	var temp float64 = float64(x)
	var res float64 = temp
	for res * res - temp > 0.00001 || res * res - temp < -0.00001 {
		res = res * 0.5 + temp / (res * 2)
	}
	return int(math.Floor(math.Abs(res)))
}