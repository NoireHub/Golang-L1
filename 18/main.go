package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Value int
	mu    sync.Mutex
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	c.mu.Lock()
	c.Value++
	c.mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	counter:= Counter{
		Value: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter.Increment(&wg)
	}

	wg.Wait()
	fmt.Println("final value =", counter.Value)
}
