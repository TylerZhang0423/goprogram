package main

import "fmt"

func test(n1 int)  {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

func testReturnMulti(i int, j int) (int, int) {
	return i+j, i-j
}

func main()  {
	//n1 := 10
	var n1 int = 10
	test(n1)
	fmt.Println("main() n1=", n1)
	fmt.Println(testReturnMulti(1, 2))
}