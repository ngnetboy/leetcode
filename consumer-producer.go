// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func consumerProducer() {
	ch := make(chan int, 10)
	var sum int64
	wgP := sync.WaitGroup{}
	wgC := sync.WaitGroup{}

	// producer
	for i := 0; i < 10; i++ {
		wgP.Add(1)
		go func(i int) {
			defer wgP.Done()
			ch <- i
		}(i)
	}

	// consumer
	for i := 0; i < 3; i++ {
		wgC.Add(1)
		go func(i int) {
			defer wgC.Done()
			fmt.Println("consumer", i)
			for value := range ch {
				atomic.AddInt64(&sum, int64(value))
			}
		}(i)
	}

	wgP.Wait()
	close(ch)

	wgC.Wait()
	fmt.Println(sum)
}
