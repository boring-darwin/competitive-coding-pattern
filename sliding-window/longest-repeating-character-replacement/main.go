package main

import "fmt"

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.

func main() {
	s := "AABABBA"
	k := 1
	t := longestRepeatingCharReplacement(s, k)
	fmt.Println(t)

}

func longestRepeatingCharReplacement(s string, k int) int {
	start, end, maxRec, max := 0, 0, 0, 0
	freq := make(map[byte]int)
	for end < len(s) {
		freq[s[end]] = freq[s[end]] + 1

		if freq[s[end]] > maxRec {
			maxRec = freq[s[end]]
		}

		if end-start+1-maxRec > k {
			freq[s[start]] = freq[s[start]] - 1
			start++
		}

		if end-start+1 > max {
			max = end - start + 1
		}
		end++
	}

	return max
}
