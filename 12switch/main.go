package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Welcome to ludo game...")
	
	diceNum := rand.IntN(6) + 1

	fmt.Println("Got dice number: ", diceNum)
	
	switch diceNum {
	case 1: 
		fmt.Println("Dice value is 1 you can can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		//fallthrough								// can move to next case auto
	case 5:
		fmt.Println("You can move 5 spot")
		//fallthrough
	case 6:
		fmt.Println("You can move 6 spot + can roll dice again")
	}
}