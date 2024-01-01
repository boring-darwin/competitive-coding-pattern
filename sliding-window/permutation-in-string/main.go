package main

import "fmt"

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.

// Example 1:
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

// Example 2:
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

// Constraints:
// 1 <= s1.length, s2.length <= 104
// s1 and s2 consist of lowercase English letters.

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Printf("%v", permutationString(s1, s2))
}

func permutationString(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count := make([]int, 26)
	s2Count := make([]int, 26)
	for i := range s1 {
		s1Count[s1[i]-'a'] += 1
		s2Count[s2[i]-'a'] += 1
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	start := 0
	end := len(s1)
	for end < len(s2) {
		if matches == 26 {
			return true
		}

		index := s2[end] - 'a'
		s2Count[index] += 1
		if s1Count[index] == s2Count[index] {
			matches += 1
		} else if s1Count[index]+1 == s2Count[index] {
			matches -= 1
		}

		index = s2[start] - 'a'
		s2Count[index] -= 1
		if s1Count[index] == s2Count[index] {
			matches += 1
		} else if s1Count[index]-1 == s2Count[index] {
			matches -= 1
		}

		start++
		end++
	}

	return matches == 26
}
