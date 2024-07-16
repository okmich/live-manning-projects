package main

import (
	"fmt"
	"time"
)

func sieveOfEratosthenes(max int) []bool {
	if max <= 2 {
		panic("Wrong input")
	}

	sieve := make([]bool, max+1)
	for i, _ := range sieve {
		if i < 2 {
			sieve[i] = false
			continue
		}
		sieve[i] = i%2 == 1
	}

	index, i := 2, 0
	for index < len(sieve) {
		sieve[index] = true
		i = index + index
		for i < len(sieve) {
			sieve[i] = false
			i += index
		}

		j := index + 1
		for ; j < len(sieve); j++ {
			if sieve[j] {
				index = j
				break
			}
		}
		// ran out finding the next prime number
		if j >= len(sieve) {
			index = j
		}
	}
	return sieve
}

func printSieve(sieve []bool) {
	for i, _ := range sieve {
		if sieve[i] {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}

// Convert the sieve into a slice holding prime numbers.
func sieveToPrimes(sieve []bool) []int {
	res := make([]int, 0)
	for i, _ := range sieve {
		if sieve[i] {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve)

		primes := sieveToPrimes(sieve)
		fmt.Println(primes)
	}
}
