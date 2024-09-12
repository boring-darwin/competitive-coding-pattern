package main

import "fmt"

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
	}
	i := 0
	checkType(i)
	ans := numberOfIslands(grid)
	fmt.Println(ans)
}

func checkType(x interface{}) {

	switch v := x.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	case float64:
		fmt.Println("It's a float64:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func numberOfIslands(grid [][]byte) int {
	islands := 0
	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r < 0 || c < 0 {
			return
		}
		if r >= len(grid) || c >= len(grid[0]) {
			return
		}
		if grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				islands += 1
				dfs(r, c)
			}
		}
	}
	return islands
}