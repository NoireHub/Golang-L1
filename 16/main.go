package main

import "fmt"

func quicksort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    var left, right []int
    pivot := arr[0]

    for _, val := range arr[1:] {
        if val < pivot {
            left = append(left, val)
        } else {
            right = append(right, val)
        }
    }

    left = quicksort(left)
    right = quicksort(right)

    return append(append(left, pivot), right...)
}

func main() {
    arr := []int{5, 1, 6, 2, 4, 3}
	
    fmt.Println("Unsorted array:", arr)
    arr = quicksort(arr)
    fmt.Println("Sorted array:", arr)
}
