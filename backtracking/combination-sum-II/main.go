package main

import (
	"fmt"
	"sort"
)

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sum to target.

// Each number in candidates may only be used once in the combination.

// Note: The solution set must not contain duplicate combinations.

// Example 1:
// Input: candidates = [10,1,2,7,6,1,5], target = 8
// Output:
// [[1,1,6],[1,2,5],[1,7],[2,6]]

// Example 2:
// Input: candidates = [2,5,2,1,2], target = 5
// Output:
// [[1,2,2],[5]]

// Constraints:
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30

func main() {
	arr := []int{2, 5, 2, 1, 2}
	target := 5
	output := combinationSum2(arr, target)
	fmt.Println(output)

}

func combinationSum2(arr []int, target int) [][]int {
	output := make([][]int, 0)
	sort.Ints(arr)
	var combination func(v, s int, subset []int)
	combination = func(t, currSum int, s []int) {
		if currSum == target {
			output = append(output, append([]int{}, s...))
			return
		} else if currSum > target {
			return
		}

		for i := t; i < len(arr); i++ {
			if i > t && arr[i-1] == arr[i] {
				continue
			}
			s = append(s, arr[i])
			combination(i+1, currSum+arr[i], s)
			s = s[:len(s)-1]
		}
	}
	subset := make([]int, 0)
	combination(0, 0, subset)
	return output
}
