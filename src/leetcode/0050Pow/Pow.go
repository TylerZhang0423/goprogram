package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 10))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1/getPowResult(x, -n)
	} else {
		return getPowResult(x, n)
	}
}

func getPowResult(x float64, n int) float64 {
	if n==1 {
		return x
	}
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		var temp = getPowResult(x, n/2)
		return temp * temp
	} else {
		var temp = getPowResult(x, n/2)
		return x * temp * temp
	}
}