package main

import (
	"fmt"
	"sync"
)

func main() {
	list := []int{2, 4, 1, 5, 7, 2, 6, 1, 1, 6, 4, 10, 33, 5, 7, 23}
	ch := make(chan int)
	// defer close(ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mergeSort(list, ch)
	}()
	wg.Wait()
	sortedArr := make([]int, 0)
	for i := range ch {
		sortedArr = append(sortedArr, i)
	}

	fmt.Println(sortedArr)
}

func mergeSort(list []int, ch chan int) {
	listLength := len(list)

	if listLength == 1 {
		ch <- list[0]
		defer close(ch)
		return
	}

	midPoint := listLength / 2

	leftCh := make(chan int)
	rightCh := make(chan int)
	go mergeSort(list[:midPoint], leftCh)
	go mergeSort(list[midPoint:], rightCh)

	go merge(leftCh, rightCh, ch)
}

func merge(left, right, c chan int) {
	defer close(c)
	val, ok := <-left
	val2, ok2 := <-right

	for ok && ok2 {
		if val < val2 {
			c <- val
			val, ok = <-left
		} else {
			c <- val2
			val2, ok2 = <-right
		}
	}
	for ok {
		c <- val
		val, ok = <-left
	}
	for ok2 {
		c <- val2
		val2, ok2 = <-right
	}
}
