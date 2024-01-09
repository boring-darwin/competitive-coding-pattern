package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 2}
	ans := subsetsWithDup(arr)
	fmt.Println(ans)
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	subset := make([]int, 0)
	sort.Ints(nums)
	var backtrack func(idx int)

	backtrack = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, subset...))
			return
		} else if idx > len(nums) {
			return
		}

		subset = append(subset, nums[idx])
		backtrack(idx + 1)

		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx++
		}

		subset = subset[:len(subset)-1]
		backtrack(idx + 1)
	}
	backtrack(0)
	return ans
}
