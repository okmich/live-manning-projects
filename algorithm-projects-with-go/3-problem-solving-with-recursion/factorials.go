package main

import "fmt"

func iterativeFactorial(num int64) int64 {
	result := int64(1)
	for i := int64(2); i <= int64(num); i++ {
		result *= i
	}

	return result
}

func factorial(num int64) int64 {
	if num <= 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}
