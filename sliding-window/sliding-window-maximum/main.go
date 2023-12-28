package main

import "fmt"

// You are given an array of integers nums, there is a sliding window of size k
// which is moving from the very left of the array to the very right.
// You can only see the k numbers in the window. Each time the sliding window
// moves right by one position.

// Return the max sliding window.

// Example 1:
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]

// type dequeue []int

// func (d dequeue) add() {

// }

func main() {
	k := 3
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	output := make([]int, 0)
	start, end := 0, 0
	dequeue := make([]int, 0)

	for end < len(nums) {
		for len(dequeue) > 0 {
			if nums[dequeue[len(dequeue)-1]] < nums[end] {
				dequeue = dequeue[:len(dequeue)-1]
			} else {
				break
			}
		}
		dequeue = append(dequeue, end)

		if start > dequeue[0] {
			dequeue = dequeue[1:]
		}

		if end+1 >= k {
			output = append(output, nums[dequeue[0]])
			start++
		}
		end++
	}
	fmt.Println(output)
}
