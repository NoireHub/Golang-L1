package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func WorkerFunc(wg *sync.WaitGroup, i int, ch chan int) {
	defer wg.Done()

	fmt.Printf("Worker %d Started\n", i)

	for data := range ch {
		fmt.Printf("Worker %d get data %d\n", i, data)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d ready again\n", i)
	}

	fmt.Printf("Worker %d Done\n", i)
}

func SenderFunc(ctx context.Context, workCh chan int, insCh chan int) {
	for {
		select {
		case data := <-insCh:
			workCh <- data
		case <-ctx.Done():
			close(workCh)
			fmt.Println("channel closed")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	n := 0
	fmt.Printf("Number of workers: ")
	fmt.Scan(&n)
	workerChan := make(chan int, n)
	insertChan := make(chan int, 1)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go SenderFunc(ctx, workerChan, insertChan)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go WorkerFunc(&wg, i, workerChan)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("*********************************\nShutdown signal received\n*********************************")
		cancelFunc()
		wg.Wait()
		fmt.Println("All workers done, shutting down!")
		os.Exit(0)
	}()

	for {
		data:= rand.Intn(100)
		insertChan <- data
	}

}
