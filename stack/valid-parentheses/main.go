package main

import "fmt"

func main() {
	s := "(){}[](}"
	output := validParentheses(s)
	fmt.Println(output)
}

func validParentheses(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		lenght := len(stack)
		if lenght == 0 {
			stack = append(stack, char)
			continue
		}

		if stack[lenght-1] == '{' && char == '}' {
			stack = stack[:lenght-1]
		} else if stack[lenght-1] == '[' && char == ']' {
			stack = stack[:lenght-1]
		} else if stack[lenght-1] == '(' && char == ')' {
			stack = stack[:lenght-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
