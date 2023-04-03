package main

import (
	"fmt"
	"math/bits"
)

func createHugeString(size uint64) string {
	return string(make([]rune, size))
}


func main() {
	fmt.Println(bits.Len(5))
}