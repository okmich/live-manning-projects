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

func eulersSieve(max int) []bool {
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
	sieve[2] = true

	for p := 3; p <= max; p += 2 {
		if sieve[p] {
			maxQ := int(max / p)
			if maxQ%2 == 0 {
				maxQ -= 1
			}
			for q := maxQ; q >= p; q -= 2 {
				if sieve[q] {
					sieve[q*p] = false
				}
			}
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
	eulerStart := time.Now()
	eulerSieve := eulersSieve(max)
	eulerElapsed := time.Since(eulerStart)

	if max <= 1000 {
		fmt.Printf("Eratosthenes Elapsed: %f seconds\n", elapsed.Seconds())
		printSieve(sieve)
		fmt.Println(sieveToPrimes(sieve))
		fmt.Println()

		fmt.Printf("Eulers Elapsed: %f seconds\n", eulerElapsed.Seconds())
		printSieve(eulerSieve)
		fmt.Println(sieveToPrimes(eulerSieve))
		fmt.Println()
	}
}
