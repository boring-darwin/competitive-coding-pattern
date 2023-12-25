package main

import (
	"fmt"
)

func generatePalindrome(N, K int) string {
	if N%2 == 0 && K < 2 {
		fmt.Println("Invalid input: Even length requires at least 2 distinct letters.")
		return ""
	}

	if K < 1 {
		fmt.Println("Invalid input: At least 1 distinct letter is required.")
		return ""
	}

	// Initialize an array with distinct letters
	letters := make([]rune, K)
	for i := 0; i < K; i++ {
		letters[i] = rune('a' + i)
	}

	// Build the first half of the palindrome
	palindrome := make([]rune, N/2)
	for i := 0; i < N/2; i++ {
		palindrome[i] = letters[i%K]
	}

	// Create the second half by reversing the first half
	secondHalf := make([]rune, N/2)
	copy(secondHalf, palindrome)
	reverseSlice(secondHalf)

	var result string
	// Combine the first and second half to form the palindrome
	if N%2 == 0 {
		result = string(append(palindrome, secondHalf...))
	} else {
		palindrome = append(palindrome, letters[len(letters)-1])
		result = string(append(palindrome, secondHalf...))
	}

	return result
}

func reverseSlice(s []rune) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

func main() {
	N := 35
	K := 3

	palindrome := generatePalindrome(N, K)

	if palindrome != "" {
		fmt.Printf("Generated Palindrome: %s\n", palindrome)
	}
}
