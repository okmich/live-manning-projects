package main

import "fmt"

type IntArray []int

func fib(n int) int {
	tab := make(IntArray, n+1)
	tab[1] = 1
	for i, _ := range tab {
		if i > 1 {
			tab[i] = tab[i-2] + tab[i-1]
		}
	}
	return tab[n]
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(8))
	fmt.Println(fib(50))
}
