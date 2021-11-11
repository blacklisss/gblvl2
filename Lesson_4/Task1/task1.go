package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var workers = make(chan struct{}, 1000)
	var sum = new(int)
	*sum = 0

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		workers <- struct{}{}

		go func(job int) {
			defer func() {
				<-workers
				*sum += 1
				wg.Done()
			}()
		}(i)
	}

	wg.Wait()
	fmt.Println(*sum)
}
