package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abd", "abda"))
}

func findTheDifference(s string, t string) byte {
	sum1 := 0
	sum2 := 0
	for _, c := range s {
		sum1 = sum1 + int(c)
	}
	for _, c := range t {
		sum2 = sum2 + int(c)
	}
	return byte(sum2 - sum1)
}
