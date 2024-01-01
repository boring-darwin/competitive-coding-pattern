package main

import "fmt"

// Given a string s, find the length of the longest
// substring
// without repeating characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.

func main() {
	s := "bbbbb"
	output := longestSubString(s)
	fmt.Println(output)
}

func longestSubString(s string) int {
	charsMap := make(map[byte]int, 0)
	output := 0
	start, end := 0, 0

	for end < len(s) {
		en := s[end]
		charsMap[en] = charsMap[en] + 1

		for charsMap[en] > 1 {
			st := s[start]
			charsMap[st] = charsMap[st] - 1
			start++
		}

		if end-start+1 > output {
			output = end - start + 1
		}
		end++
	}
	return output
}
