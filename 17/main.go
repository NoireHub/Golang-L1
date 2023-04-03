package main

import "fmt"

func binarySearch(slice []int, target int) int {
	low, high := 0, len(slice)-1
	for low <= high {
		mid := (low + high) / 2
		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	slice := []int{1, 3, 5, 7, 9}
	target := 5

	index := binarySearch(slice, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Target %d not found\n", target)
	}
}
