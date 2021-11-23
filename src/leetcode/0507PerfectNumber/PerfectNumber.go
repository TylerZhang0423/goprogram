package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(28))
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 0
	for i:=1; i<=mySqrt(num); i++ {
		if num % i == 0 {
			sum = sum + i + num / i
		}
		//fmt.Println(sum)
	}
	sum = sum - num
	if sum == num {
		return true
	}
	return false
}

func mySqrt(x int) int {
	var temp float64 = float64(x)
	var res float64 = temp
	for res * res - temp > 0.00001 || res * res - temp < -0.00001 {
		res = res * 0.5 + temp / (res * 2)
	}
	return int(math.Floor(math.Abs(res)))
}