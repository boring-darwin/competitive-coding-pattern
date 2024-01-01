package main

import "fmt"

func main() {
	list := []int{2, 4, 1, 5, 7, 2, 6, 1, 1, 6, 4, 10, 33, 5, 7, 23}
	output := mergeSort(list)
	fmt.Println(output)
}

func mergeSort(list []int) []int {
	listLength := len(list)

	if listLength == 1 {
		return list
	}

	midPoint := listLength / 2

	leftHalf := mergeSort(list[:midPoint])
	rightHalf := mergeSort(list[midPoint:])

	return merge(leftHalf, rightHalf)
}

func merge(leftHalf, rightHalf []int) []int {
	output := make([]int, 0)
	i, j := 0, 0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] < rightHalf[j] {
			output = append(output, leftHalf[i])
			i += 1
		} else {
			output = append(output, rightHalf[j])
			j += 1
		}
	}
	output = append(output, leftHalf[i:]...)
	output = append(output, rightHalf[j:]...)

	return output
}
