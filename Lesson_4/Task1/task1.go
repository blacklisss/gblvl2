package main

import (
	"fmt"
	"sync"
	"time"
)

func Summer(sum *int) {
	var wg sync.WaitGroup
	var workers = make(chan struct{}, 1000)

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		workers <- struct{}{}

		time.Sleep(time.Millisecond * 10)
		go func() {
			defer func() {
				<-workers
				*sum += 1
				wg.Done()
			}()
		}()
	}

	wg.Wait()
}

func main() {
	var sum = new(int)
	*sum = 0

	Summer(sum)
	fmt.Println(*sum)
}
