package main

import (
	"fmt"
	"github.com/NoireHub/Golang-L1/24/point"
)

func main() {
	point1:= point.NewPoint(2.0,3.1)
	point2:= point.NewPoint(5,-3)

	fmt.Println(point1.GetDistance(point2))
}
