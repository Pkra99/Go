package main

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// aggregate is a higher order func
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	println("Add: ", aggregate(2, 4, 5, add))
	println("Multiply: ", aggregate(2, 4, 5, multiply))
}
