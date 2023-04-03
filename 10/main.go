package main

import (
	"fmt"
	"math"
)


func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	rangeMap := make(map[int][]float64)

	for _, val := range data {
		key:= int(val - math.Mod(val, 10))
		rangeMap[key] = append(rangeMap[key], val)
	}

	for key, val := range rangeMap {
		fmt.Printf("%d:%v\n", key, val)
	}
}
