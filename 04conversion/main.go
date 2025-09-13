package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Please provide rating between 1 to 5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ :=reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding +1 to rating: ", numRating + 1)
	}

}
