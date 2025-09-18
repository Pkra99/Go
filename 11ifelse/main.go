package main

import (
	"fmt"
)

func main() {
	num := 15

	if num > 10 {
		fmt.Println("Number is larger than 10")
	} else if num == 30 {
		fmt.Println("Number is equal to 30")
	} else {
		fmt.Println("Hey! there")
	}

	if 11 % 2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num := 2; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}