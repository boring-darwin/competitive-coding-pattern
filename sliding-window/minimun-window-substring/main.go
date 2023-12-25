// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring
//  of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

// The testcases will be generated such that the answer is unique.

// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.

// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

// Constraints:
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s and t consist of uppercase and lowercase English letters.

package main

import "fmt"

func main() {
	str := "acbbaca"
	t := "aba"
	output := minWindow(str, t)
	fmt.Println(output)
}

func minWindow(s string, t string) string {
	output := ""
	tMap := make(map[uint8]int)
	windowMap := make(map[uint8]int)
	have := 0

	if len(t) == 0 || len(s) < len(t) {
		return ""
	}
	for i := range t {
		tMap[t[i]]++
	}
	start, end := 0, 0
	for end < len(s) {
		windowMap[s[end]]++
		if tMap[s[end]] != 0 && windowMap[s[end]] == tMap[s[end]] {
			have += 1
		}
		for have == len(tMap) {
			if output == "" {
				output = s[start : end+1]
			}
			if end-start+1 < len(output) {
				output = s[start : end+1]
			}

			windowMap[s[start]] = windowMap[s[start]] - 1
			if windowMap[s[start]] < tMap[s[start]] {
				have -= 1
			}
			start++
		}
		end++
	}
	return output
}
