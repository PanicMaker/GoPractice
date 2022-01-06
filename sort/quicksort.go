package main

import "fmt"

func main() {
	list := []int{5, 8, 2, 9, 1, 0, 4, 7, 3, 6}
	list = quicksort(list, 0, len(list)-1)
	fmt.Println(list)
}

func quicksort(arr []int, start, end int) []int {
	if end <= start {
		return nil
	}
	pivot := partition(arr, start, end)
	quicksort(arr, start, pivot-1)
	quicksort(arr, pivot+1, end)
	return arr
}

func partition(arr []int, left, right int) int {

	pivot := arr[left]
	index := left + 1

	for i := index; i < right; i++ {
		if pivot > arr[i] {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[left], arr[index] = arr[index], pivot
	return index
}
