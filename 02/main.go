package main

import (
	"fmt"
	"sync"
)

func GetSquare(wg *sync.WaitGroup, number *int) {
	*number = *number * *number
	wg.Done()
}

func main() {
	data:= []int{2,4,6,8,10}
	var wg sync.WaitGroup

	for i:= range data {
		wg.Add(1)
		go GetSquare(&wg,&data[i])
		
	}
	wg.Wait()

	for _,val:= range data {
		fmt.Println(val)
	}
}
