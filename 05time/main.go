package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time Machine")

	presentTime := time.Now()

	formattedTime := presentTime.Format("15:04:05")
	formattedDate := presentTime.Format("01-02-2006")

	fmt.Println("The current time is:", formattedTime)
	fmt.Println("The current date is:", formattedDate)

	modifiedDate := time.Date(2019, time.April, 10, 23, 11, 0, 0, time.UTC)
	fmt.Println("Modified Date: ", modifiedDate)
}