package main

import (
	"fmt"
	"strconv"
)

func fibonacci(num int64) int64 {
	if num == 0 || num == 1 {
		return num
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
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
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}
