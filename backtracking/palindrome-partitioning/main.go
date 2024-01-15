package main

import "fmt"

func main() {
	s := "aab"
	ans := palindromePartioning(s)
	fmt.Println(ans)
}

func palindromePartioning(s string) [][]string {
	output := make([][]string, 0)
	subset := make([]string, 0)

	var backtrack func(idx int)

	backtrack = func(idx int) {
		if idx == len(s) {
			output = append(output, append([]string{}, subset...))
			return
		}
		for i := idx; i < len(s); i++ {
			s := s[idx : i+1]
			if isPalindrome(s) {
				subset = append(subset, s)
				backtrack(i + 1)
				subset = subset[:len(subset)-1]
			}
		}
	}
	backtrack(0)
	return output
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
