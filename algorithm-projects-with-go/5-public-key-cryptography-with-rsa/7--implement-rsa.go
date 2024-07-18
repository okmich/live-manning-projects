package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

func gcd(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else {
		max := max(a, b)
		min := min(a, b)

		// q := max / min
		r := max % min
		return gcd(min, r)
	}
}

func lcm(a, b int) int {
	gcd := gcd(a, b)
	return (a * b) / gcd
}

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

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q int) int {
	return lcm(p-1, q-1)
}

// Pick a random exponent e in the range (2, λn)
// such that gcd(e, λn) = 1.
func randomExponent(λn int) int {
	for {
		e := randRange(2, λn)
		_gcd := gcd(e, λn)
		if _gcd == 1 {
			return e
		}
	}
}

// Calculate the inverse of a in the modulus.
// See https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Modular_integers
// Look at:
//
//	Section "Computing multiplicative inverses in modular structures"
//	Subsection "Modular integers"
func inverseMod(a, modulus int) int {
	t := 0
	newt := 1
	r := modulus
	newr := a

	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}

	if r > 1 {
		return -1
	}

	if t < 0 {
		t = t + modulus
	}
	return t
}

// copied from final solution
func main() {
	// Pick two random primes p and q.
	const numTests = 100
	p := findPrime(10000, 50000, numTests)
	q := findPrime(10000, 50000, numTests)

	// Calculate the public key modulus n.
	n := p * q

	// Calculate Carmichael's totient function λ(n).
	λn := totient(p, q)

	// Pick a random public key exponent e in the range [3, λn)
	// where gcd(e, λn) = 1.
	e := randomExponent(λn)

	// Find the inverse of e mod λn.
	d := inverseMod(e, λn)

	// Print out the important values.
	fmt.Printf("*** Public ***\n")
	fmt.Printf("Public key modulus:    %d\n", n)
	fmt.Printf("Public key exponent e: %d\n", e)
	fmt.Printf("\n*** Private ***\n")
	fmt.Printf("Primes:    %d, %d\n", p, q)
	fmt.Printf("λ(n):      %d\n", λn)
	fmt.Printf("d:         %d\n", d)

	for {
		var m int
		fmt.Printf("\nMessage:    ")
		fmt.Scan(&m)
		if m < 1 {
			break
		}

		ciphertext := fastExpMod(m, e, n)
		fmt.Printf("Ciphertext: %d\n", ciphertext)

		plaintext := fastExpMod(ciphertext, d, n)
		fmt.Printf("Plaintext:  %d\n", plaintext)
	}
}
