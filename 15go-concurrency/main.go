package main

import (
	"fmt"
	"sync"
	"time"
)

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println("Letter: ", string(i))
		time.Sleep(time.Millisecond * 500)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("Number: ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {	
	var wg sync.WaitGroup
	wg.Add(2)
	go printLetters(&wg)
	go printNumbers(&wg)

	wg.Wait()
}
