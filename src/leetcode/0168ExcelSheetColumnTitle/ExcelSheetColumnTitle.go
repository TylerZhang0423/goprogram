package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(12))
}

func convertToTitle(columnNumber int) string {
	title := ""
	columnNumber = columnNumber - 1
	for columnNumber >= 26 {
		quotient := columnNumber / 26
		remainder := columnNumber % 26
		columnNumber = quotient - 1
		title =  string(remainder + 65) + title
	}
	title =  string(columnNumber + 65) + title
	return title
}
