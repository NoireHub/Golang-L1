package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	sampleMap:= make(map[string]int)

	keys := []string{"key1", "key2", "key3", "key4", "key5"}

	for _, k := range keys {
		wg.Add(1)
		go func(key string) {
			mu.Lock()
			sampleMap[key] = rand.Intn(10)
			mu.Unlock()
			wg.Done()
		}(k)
	}

	wg.Wait()
	fmt.Println(sampleMap)
}
