package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

func main() {
	nums := []int{1, 2, -2, -1}
	output := sumThree(nums)
	fmt.Println(output)
	// fmt.Println(sumTwo(nums, -3))
}

func sumThree(nums []int) [][]int {
	sort.Ints(nums)
	solution := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		sol := sumTwo(nums[i+1:], nums[i])
		if len(sol) > 0 {
			solution = append(solution, sol...)
		}
	}
	return solution
}

func sumTwo(nums []int, target int) [][]int {
	i := 0
	j := len(nums) - 1

	solution := make([][]int, 0)
	for i < j {
		sum := nums[i] + nums[j] + target
		if sum == 0 {
			solution = append(solution, []int{target, nums[i], nums[j]})
			j--
			for i < j && nums[j] == nums[j+1] {
				j--
			}
		}
		if sum > 0 {
			j--
		}
		if sum < 0 {
			i++
		}
	}
	return solution
}
