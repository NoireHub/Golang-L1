package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	data := []int{2, 4, 6, 8, 10}
	result:= 0

	for i := range data {
		wg.Add(1)
		go func (num int)  {
			mu.Lock()
			result += num * num
			mu.Unlock()
			wg.Done()
		}(data[i])
	}

	wg.Wait()
	fmt.Println(result)
}
