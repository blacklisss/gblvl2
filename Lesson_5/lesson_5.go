package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	var workers = make(chan struct{}, 10)
	var sum = new(int)
	*sum = 0

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		workers <- struct{}{}

		go func(job int) {
			lock.Lock()
			*sum += 1
			defer func() {
				lock.Unlock()
				<-workers
				wg.Done()

			}()
		}(i)
	}

	wg.Wait()
	fmt.Println(*sum)
}
