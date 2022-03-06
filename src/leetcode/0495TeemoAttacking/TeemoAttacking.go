package main

import "fmt"

func main() {
	timeSeries := []int {1, 2}
	duration := 1
	fmt.Println(timeSeries, duration)
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	poisonTime := 0
	for i:=0; i<len(timeSeries); i++ {
		if i+1 < len(timeSeries) {
			if timeSeries[i] + duration <= timeSeries[i+1] {
				poisonTime = poisonTime + duration
			} else {
				poisonTime = poisonTime + timeSeries[i+1] - timeSeries[i]
			}
		} else {
			poisonTime = poisonTime + duration
		}
	}
	return poisonTime
}
