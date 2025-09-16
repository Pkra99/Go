package main

import "fmt"

func main() {
	myNum := 20
	var ptr = &myNum		//referencing

	fmt.Println(ptr)
	fmt.Println(*ptr)		//dereferencing		//*ptr is actual value not the memory address

	*ptr = *ptr + 5
	fmt.Println("New value of myNum: ", myNum)
}
