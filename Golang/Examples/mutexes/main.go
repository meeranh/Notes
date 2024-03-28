package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	var score = []int{0}

	waitGroup.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Func One")
		mut.RLock()
		score = append(score, 1)
		mut.RUnlock()
		defer wg.Done()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Func Two")
		mut.RLock()
		score = append(score, 2)
		mut.RUnlock()
		defer wg.Done()
	}(waitGroup, mutex)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Func Three")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		defer wg.Done()
	}(waitGroup, mutex)

	waitGroup.Wait()

	fmt.Println(score)
}
