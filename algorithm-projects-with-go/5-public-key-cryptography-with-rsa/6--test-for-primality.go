package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const numTests = 20

// from milestone 2
func fastExpMod(num, pow, mod int) int {
	var result int = 1
	current_pow := num
	exp := pow
	for exp >= 1 {
		if exp%2 == 1 {
			result = (result * current_pow) % mod
		}
		exp /= 2
		current_pow = (current_pow * current_pow) % mod
	}
	return result
}

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

// Return a pseudo random number in the range [min, max).
func randRange(min int, max int) int {
	return min + random.Intn(max-min)
}

// Perform tests to see if a number is (probably) prime.
func isProbablyPrime(p int, numTests int) bool {
	var res bool = true
	var i int = 0
	for ; i < numTests; i++ {
		n := randRange(1, p-1)
		expMod := fastExpMod(n, p-1, p)
		res = res && expMod == 1
		if !res {
			return false
		}
	}
	return true
}

// Probabilistically find a prime number within the range [min, max).
func findPrime(min, max, numTests int) int {
	for {
		num := randRange(min, max)
		if num%2 == 0 {
			continue
		}
		isPrime := isProbablyPrime(num, numTests)
		if isPrime {
			return num
		}
	}
	return 0
}

func testKnownValues() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Println("Probability: 99.999905%\n")
	fmt.Println("Primes:")
	for _, v := range primes {
		if isProbablyPrime(v, numTests) {
			fmt.Printf("%v \t Prime\n", v)
		} else {
			fmt.Printf("%v \t Composite \n", v)
		}
	}
	fmt.Println("\nComposites:")
	for _, v := range composites {
		if isProbablyPrime(v, numTests) {
			fmt.Printf("%v \t Prime\n", v)
		} else {
			fmt.Printf("%v \t Composite \n", v)
		}
	}
}

func main() {
	fmt.Println(time.Now())

	// Test some known primes and composites.
	testKnownValues()

	// Generate random primes.
	for {
		// Get the number of digits.
		var numDigits int
		fmt.Printf("\n# Digits: ")
		fmt.Scan(&numDigits)
		if numDigits < 1 {
			break
		}

		// Calculate minimum and maximum values.
		min := int(math.Pow(10.0, float64(numDigits-1)))
		max := 10 * min
		if min == 1 {
			min = 2
		} // 1 is not prime.

		// Find a prime.
		fmt.Printf("Prime: %d\n", findPrime(min, max, numTests))
	}
}
