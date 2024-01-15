package main

import "fmt"

func main() {

	matrix := [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	word := "SEE"

	output := exist(matrix, word)
	fmt.Println(output)
}

// func searchWord(matrix [][]string, word string) bool {
// 	// visited := map[int]struct{}{}
// 	var backtrack func(i, j, curr int) bool
// 	ROW := len(matrix) - 1
// 	COL := len(matrix[0]) - 1
// 	backtrack = func(i, j, curr int) bool {
// 		if curr == len(word)-1 {
// 			return true
// 		}

// 		// _, exist := visited[i*10+j]
// 		if i > ROW || j > COL || i < 0 || j < 0 || matrix[i][j] == "*" || string(word[curr]) != matrix[i][j] || curr == len(word) {
// 			return false
// 		}

// 		tmp := matrix[i][j]
// 		matrix[i][j] = "*"
// 		// visited[i*10+j] = struct{}{}

// 		res := backtrack(i+1, j, curr+1) || backtrack(i-1, j, curr+1) || backtrack(i, j+1, curr+1) || backtrack(i, j-1, curr+1)
// 		matrix[i][j] = tmp
// 		return res
// 	}

// 	for i := 0; i < ROW; i++ {
// 		for j := 0; j < COL; j++ {
// 			if backtrack(i, j, 0) {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

func exist(board [][]string, word string) bool {
	n := len(board)
	m := len(board[0])

	var dfs func(i int, j int, curr int) bool
	dfs = func(i int, j int, curr int) bool {
		if i < 0 || j < 0 || i >= n || j >= m || curr == len(word) {
			return false
		}
		if board[i][j] != string(word[curr]) || board[i][j] == "*" {
			return false
		}
		if curr == len(word)-1 {
			return true
		}

		tmp := board[i][j]
		board[i][j] = "*"

		res := dfs(i+1, j, curr+1) || dfs(i-1, j, curr+1) || dfs(i, j-1, curr+1) || dfs(i, j+1, curr+1)

		board[i][j] = tmp

		return res
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}