package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions in GO")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)
	// channels
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	var score = []int{0}

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		ch1 <- true
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		<-ch1
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		ch2 <- true
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		<-ch2
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		ch3 <- true
	}(wg, mut)

	<-ch3
	wg.Wait()

	mut.Lock()
	fmt.Println("Score: ", score)
	mut.Unlock()
}
