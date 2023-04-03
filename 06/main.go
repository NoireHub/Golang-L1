package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Первый способ – закрыть канал и выйти из цикла чтения канала
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("[Goroutine 1]: Done")
				return
			}
			fmt.Println("in the loop #1")
		}
	}()

	ch <- 1
	ch <- 1
	ch <- 1
	close(ch)
	wg.Wait()

	// Второй способ - создать канал, для остановки выполнения 
	quit := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("[Goroutine 2]: Done")
				return
			default:
				fmt.Println("in the loop #2")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- true
	wg.Wait()

	// Третий способ - остановка осуществляется через контекст
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
    go func(ctx context.Context) {
		defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                fmt.Println("[Goroutine 3]: Done")
                return
            default:
                fmt.Println("in the loop #3")
				time.Sleep(500 * time.Millisecond)
            }   
        }
    }(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()

}
