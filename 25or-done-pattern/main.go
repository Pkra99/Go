package main

import "fmt"

 
func generator() <- chan int {
	v := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending: %d\n", i)
			v <- i
		}
	}()
	return v
}

func orDone(done <-chan struct{}, c <-chan int) <-chan int {
	valStream :=  make(chan int)

	go func() {
		defer close(valStream)

		for{
			select{
			case <- done:
				return 
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <- done:
					return
				}
			}
		}
	}()
	return valStream
}

func main() {
	done := make(chan struct{})

	ch := generator()
	for v := range orDone(done, ch) {
		fmt.Printf("Recived: %d\n", v)

		if v == 2 {
			close(done)
		}
	}
}