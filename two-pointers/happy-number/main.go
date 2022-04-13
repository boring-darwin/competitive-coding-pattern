package main

import "fmt"

func main() {

	ok := true
	var slow = 0
	var fast = 0
	for slow != fast || ok {

		if ok {
			slow = 13
			fast = 13
		}
		slow = calSquareSum(slow)
		fast = calSquareSum(calSquareSum(fast))
	}

	fmt.Println(slow == 1)

}

func calSquareSum(num int) int {

	var sum = 0
	for num > 0 {
		d := num % 10
		sum = sum + (d * d)
		num /= num
	}

	return sum
}
