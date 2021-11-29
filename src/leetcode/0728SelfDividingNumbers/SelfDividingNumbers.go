package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 20))
}

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i:=left; i<=right; i++ {
		if isSelfDividingNumber(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDividingNumber(num int) bool {
	if num >= 0 && num <= 9 {
		return true
	}
	temp := num
	for temp > 0 {
		quotient := temp % 10
		if quotient == 0 || num % (temp % 10) != 0 {
			return false
		}
		temp = temp / 10
	}
	return true
}
