package main

import "fmt"

func main() {
	fmt.Println("Aminals List: ")

	var animalList [4]string

	animalList[0] = "cats"
	animalList[1] = "dogs"
	animalList[3] = "cows"

	fmt.Println("list: ",animalList)

	var birdList = [3]string{"parrot", "peigon", "crow"}
	fmt.Println(birdList)
} 