package main

import "fmt"

// using simple recursion for calculating the fibonacci sequence
func fibonacciRecursion(n int) int {
	if(n <= 1) {
		return n
	}
	return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
}

// using memoization 
func fibonacciDp(n int, dp []int) int {
	if(n <= 1){
		return n
	}

	if(dp[n] != -1) {
		return dp[n]
	}

	return dp[n-1] + dp[n-2]
}

// using Tabulation
func fibonacciDpTabulation(n int) int {
	prev := 1
	secondPrev := 0

	for i := 2; i <= n; i++ {
		current := prev + secondPrev
		secondPrev = prev
		prev = current
	}
	return prev
}

func main() {
	n := 7
	fmt.Print("Usiing Recursion: ", fibonacciRecursion(n))	
	dp := make([]int, n + 1)
	for i := range dp {
		dp[i] = -1
	}
	fmt.Print("\nUsing DP with memoization: ", fibonacciDp(n, dp))
	fmt.Print("\nUsing DP with Tabulation: ", fibonacciDpTabulation(n))
}