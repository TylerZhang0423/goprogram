package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}

func firstBadVersion(n int) int {
	var front int = 1
	var end int = n
	if n == 1 {
		return n
	}
	for front <= end {
		if end - front == 1 {
			if isBadVersion(front) {
				return front
			} else {
				return end
			}
		}
		var mid int = (front + end) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			front = mid
		}
	}
	return end
}

func isBadVersion(num int) bool {
	if num >= 4 {
		return true
	}
	return false
}
