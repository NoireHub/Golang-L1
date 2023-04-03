package main

import (
	"fmt"
	"time"
)

func main() {
	t:= time.Now()
	sleep(5 * time.Second)
	fmt.Println(time.Since(t))
}

func sleep(d time.Duration) {
	<-time.After(d)
}
