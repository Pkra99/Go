package main

import "fmt"

//totalStudent := 250		//can not user outside of the func, normal variables are allowed

const UserId string = "2345678"

func main() {
	var username string = "Hello"
	fmt.Println(username)
	fmt.Printf("Variable of is type: %T \n", username)	// %T tell us about the datatype of variables

	var isUser bool = false
	fmt.Println(isUser)
	fmt.Printf("Variable of is type: %T \n", isUser)

	// uint is unsingned(non-negative) 8-bit integer (0 to 255)
	var smallInt uint8 = 255  // 256 will throw an error
	fmt.Println(smallInt)
	fmt.Printf("Variable of is type: %T \n", smallInt)

	var smallFloat float32 = 255.3353535535535
	fmt.Println(smallFloat)
	fmt.Printf("Variable of is type: %T \n", smallFloat)

	var randomInt int		//does not give garbage value
	fmt.Println(randomInt)
	fmt.Printf("Variable of is type: %T \n", randomInt)

	var randomString string		//will give empty string
	fmt.Println(randomString)
	fmt.Printf("Variable of is type: %T \n", randomString)

	// implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable of is type: %T \n", website)
	//website = 3		//changing type is not allowed

	// no var style
	numberOfStudentRegistered := 200
	fmt.Println(numberOfStudentRegistered)

	fmt.Println(UserId)
	fmt.Printf("Variable of is type: %T \n", UserId)
}
