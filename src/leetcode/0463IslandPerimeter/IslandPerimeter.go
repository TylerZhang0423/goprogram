package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0 , 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := range grid {
		for j := range grid[i] {
			perimeter = perimeter + getPerimeter(grid, i, j)
		}
	}
	return perimeter
}

func getPerimeter(grid [][]int, i int, j int) int {
	subPerimeter := 0
	if grid[i][j] == 1 {
		subPerimeter = 4
		if i-1 >= 0 && grid[i-1][j] == 1 {
			subPerimeter = subPerimeter - 1
		}
		if j-1 >= 0 && grid[i][j-1] == 1 {
			subPerimeter = subPerimeter - 1
		}
		if i+1 < len(grid) && grid[i+1][j] == 1 {
			subPerimeter = subPerimeter - 1
		}
		if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
			subPerimeter = subPerimeter - 1
		}
	}
	return subPerimeter
}
