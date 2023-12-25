package main

import "fmt"

func main() {
	fmt.Println(FindOccurance("aawwwfffaa"))
}

func FindOccurance(S string) string {
	occurance := make([]int, 26)

	for _, s := range S {
		occurance[s-'a'] = occurance[s-'a'] + 1
	}
	fmt.Println(occurance)
	best_char := rune(S[0])
	best_times := 0
	for i := 97; i < 122; i++ {
		if occurance[i-'a'] >= best_times {
			if best_times == occurance[i-'a'] && comesBefore(best_char, rune(occurance[i-'a']), S) {
				continue
			}
			best_char = rune(i)
			best_times = occurance[i-'a']
		}
	}
	return string(best_char)
}

func comesBefore(a, b rune, str string) bool {
	for _, s := range str {
		if s == a {
			return true
		}
		if s == b {
			return false
		}
	}
	return false
}
