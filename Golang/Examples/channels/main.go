package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	myChan := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isOpen := <-ch
		fmt.Println(isOpen)

		fmt.Println(val)
		fmt.Println(<-ch)
		fmt.Println(<-ch)

		defer wg.Done()
	}(myChan, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		ch <- 7
		close(ch)

		defer wg.Done()
	}(myChan, wg)
	wg.Wait()
}
