package main

import (
	"fmt"
	"time"
)

func delta(startTime time.Time) int {
	return int(time.Since(startTime).Seconds())
}


func main() {

	startTime := time.Now()
	ch := make(chan int)

	n := 0
	fmt.Printf("Enter time in seconds: ")
	fmt.Scan(&n)

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
	}()

	for i := 0; delta(startTime) < n; i++ {
		ch <- i
	}
	close(ch)
}
