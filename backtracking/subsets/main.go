package main

import "fmt"

// Given an integer array nums of unique elements, return all possible
// subsets
//  (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// Example 2:

// Input: nums = [0]
// Output: [[],[0]]

// Constraints:

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
func main() {
	arr := []int{1, 2, 3}
	ans := subSet(arr)
	fmt.Println(ans)
}

func subSet(arr []int) [][]int {
	output := make([][]int, 0)
	subSet := make([]int, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {

		if idx == len(arr) {
			output = append(output, append([]int{}, subSet...))
			return
		}

		subSet = append(subSet, arr[idx])
		backtrack(idx + 1)

		subSet = subSet[:len(subSet)-1]
		backtrack(idx + 1)
	}
	backtrack(0)
	return output
}
