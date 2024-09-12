package main

import "fmt"

func main() {
	matrix := [][]int{{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}

	target := 1

	ans := searchMatrix(matrix, target)
	fmt.Println(ans)

}

func searchMatrix(matrix [][]int, target int) bool {
	sizeOfArr := len(matrix[0]) - 1
	start := 0
	end := len(matrix) - 1

	for start <= end {
		mid := (start + end) / 2

		if target >= matrix[mid][0] && target <= matrix[mid][sizeOfArr] {
			return binarySearch(matrix[mid], target)
		}

		if matrix[mid][sizeOfArr] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
