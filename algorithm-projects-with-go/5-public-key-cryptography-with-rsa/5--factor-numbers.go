package main

import (
	"fmt"
	"time"
)

var primes []int

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

func findFactors(num int) []int {
	factors := make([]int, 0)
	for num%2 == 0 {
		factors = append(factors, 2)
		num /= 2
	}

	for factor := 3; factor <= num; {
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		} else {
			factor += 2
		}
	}
	if num > 1 {
		factors = append(factors, num)
	}
	return factors
}

func findFactorsSieve(num int) []int {
	factors := make([]int, 0)
	for num%2 == 0 {
		factors = append(factors, 2)
		num /= 2
	}
	for _, factor := range primes {
		for num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
			if num == 1 {
				break
			}
		}
		if factor*factor > num {
			factors = append(factors, num)
			break
		}
	}
	return factors
}

func multiplySlice(nums []int) int {
	product := 1
	for _, v := range nums {
		product *= v
	}
	return product
}

func main() {
	var num int

	//preload the sieve
	primes = sieveToPrimes(eulersSieve(2000000000))

	//run the loop
	for {
		fmt.Printf("Number: ")
		fmt.Scan(&num)

		if num < 2 {
			break
		}

		// Find the factors the slow way.
		start := time.Now()
		factors := findFactors(num)
		elapsed := time.Since(start)
		fmt.Printf("findFactors:       %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = findFactorsSieve(num)
		elapsed = time.Since(start)
		fmt.Printf("findFactorsSieve: %f seconds\n", elapsed.Seconds())
		fmt.Println(multiplySlice(factors))
		fmt.Println(factors)
		fmt.Println()
	}

}
