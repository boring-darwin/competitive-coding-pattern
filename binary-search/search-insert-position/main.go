package main

import "fmt"

// Given a sorted array of distinct integers and a target value,
// return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

// Constraints:

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums contains distinct values sorted in ascending order.
// -104 <= target <= 104

func main() {
	arr := []int{1, 3, 5}
	ans := binarySearch(arr, 1)
	fmt.Println(ans)
}

func binarySearch(arr []int, numToFind int) int {

	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2

		if arr[mid] == numToFind {
			return mid
		}

		if arr[mid] < numToFind {
			left = mid
		}
		if arr[mid] > numToFind {
			right = mid
		}

		if left+1 == right {
			break
		}
	}

	if arr[left] == numToFind {
		return left
	}
	if arr[right] == numToFind {
		return right
	}
	if left == right {
		if arr[left] > numToFind {
			return left + 1
		} else {
			return 0
		}
	} else {
		if numToFind > arr[right] {
			return right + 1
		} else if numToFind < arr[left] {
			return left
		} else {
			return right
		}
	}
}
