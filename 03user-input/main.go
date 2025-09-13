package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please provide ratings: ")

	// comma ok || err err

	input, _ := reader.ReadString('\n')		// we put "_" if we don't care about the error else we use "err"
	fmt.Println("Thanks for raring, ", input)
	fmt.Printf("Type of this rating is %T", input) // output will be string

	
}
