package main

import (
    "fmt"
)

func intersection(a, b []int) []int {
    setA := make(map[int]bool)
	res := []int{}

    for _, x := range a {
        setA[x] = true
    }

    for _, y := range b {
        if setA[y] {
            res = append(res, y)
        }
    }

    return res
}

func main() {
    setA := []int{1, 2, 3, 4, 5}
    setB := []int{3, 4, 5, 6, 7}

    fmt.Println(intersection(setA, setB))
}
