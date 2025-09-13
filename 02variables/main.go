package main

import "fmt"

func main() {
	var username string = "Hello"
	fmt.Println(username)
	fmt.Printf("Variable of is type: %T \n", username)	// %T tell us about the datatype of variables

	var isUser bool = false
	fmt.Println(isUser)
	fmt.Printf("Variable of is type: %T \n", isUser)

	var smallInt uint8 = 255  // 256 will throw an error
	fmt.Println(smallInt)
	fmt.Printf("Variable of is type: %T \n", smallInt)
}
