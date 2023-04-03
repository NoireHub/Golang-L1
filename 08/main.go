package main

import (
    "fmt"
)

func setBit(num int64, pos uint, val bool) int64 {
    if val {
        num |= (1 << pos)
    } else {
        num &^= (1 << pos)
    }
    return num
}

func main() {
    var num int64 = 42
    fmt.Printf("Start: %b - %d\n", num, num)

    num = setBit(num, 2, true)
    fmt.Printf("Changed 2nd bit to 1: %b - %d\n", num, num)
    
    num = setBit(num, 1, false)
    fmt.Printf("Changed 1st bit to 0: %b - %d\n", num, num)
}


