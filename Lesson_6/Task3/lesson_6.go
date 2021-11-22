package main

import (
	"fmt"
	"sync"
)

func Summer(sum *int) {
	var (
		//lock    sync.Mutex
		wg      sync.WaitGroup
		workers = make(chan struct{}, 10)
	)

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		workers <- struct{}{}

		go func() {
			//lock.Lock()
			*sum += 1
			defer func() {
				//lock.Unlock()
				<-workers
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
