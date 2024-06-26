package main

import (
	"fmt"
	"strconv"
)

var fibonacciValues []int64

func fibonacciOnTheFly(n int64) int64 {
	if n == 0 || n == 1 || int64(len(fibonacciValues)) < n {
		return fibonacciValues[n]
	} else {
		res := fibonacciOnTheFly(n-1) + fibonacciOnTheFly(n-2)
		fibonacciValues = append(fibonacciValues, res)
		return res
	}
}

func initializeSlice() {
	fibonacciValues = make([]int64, 93)

	fibonacciValues[0] = 0
	fibonacciValues[1] = 1
	for i := 2; i < 93; i++ {
		fibonacciValues[i] = fibonacciValues[i-1] + fibonacciValues[i-2]
	}
}

func fibonacciPrefilled(n int64) int64 {
	if n < 0 || int(n) > len(fibonacciValues) {
		return -1
	}
	return fibonacciValues[n]
}

func fibonacciBottomUp(n int64) int64 {
	if n <= 1 {
		return int64(n)
	}

	var fibI, fibIMinus1, fibIMinus2 int64
	fibIMinus2 = 0
	fibIMinus1 = 1
	fibI = fibIMinus1 + fibIMinus2
	for i := int64(1); i < n; i++ {
		// Calculate this value of fibI.
		fibI = fibIMinus1 + fibIMinus2

		// Set fibIMinus2 and fibIMinus1 for the next value.
		fibIMinus2 = fibIMinus1
		fibIMinus1 = fibI
	}
	return fibI
}

func main() {
	// Fill-on-the-fly.
	fibonacciValues = make([]int64, 2)
	fibonacciValues[0] = 0
	fibonacciValues[1] = 1

	// Prefilled.
	initializeSlice()

	for {
		// Get n as a string.
		var nString string
		fmt.Printf("N: ")
		fmt.Scanln(&nString)

		// If the n string is blank, break out of the loop.
		if len(nString) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(nString, 10, 64)

		// Uncomment one of the following.
		fmt.Printf("fibonacciOnTheFly(%d) = %d\n", n, fibonacciOnTheFly(n))
		// fmt.Printf("fibonacciPrefilled(%d)  = %d\n", n, fibonacciPrefilled(n))
		// fmt.Printf("fibonacciBottomUp(%d)  = %d\n", n, fibonacciBottomUp(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacciValues); i++ {
		fmt.Printf("%d: %d\n", i, fibonacciValues[i])
	}
}
