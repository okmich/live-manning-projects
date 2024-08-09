package main

import (
	"fmt"
	"math"
)

// https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/
// Given an array arr[] of size N, the task is to find the length of the Longest Increasing Subsequence (LIS) i.e.,
// the longest possible subsequence in which the elements of the subsequence are sorted in increasing order.

func longestIncreasingSubsequence(target []int) []int {
	myTarget := append(target, math.MaxInt)
	targetLength := len(myTarget)
	table := make([][]int, targetLength)
	for i, _ := range table {
		table[i] = []int{}
	}

	for i, _ := range table {
		for j := i + 1; j < len(table); j++ {
			if myTarget[i] < myTarget[j] {
				temp := append(table[i], myTarget[i])
				if len(temp) > len(table[j]) {
					table[j] = temp
				}
			}
		}
	}
	return table[targetLength-1]
}

func runMain(target []int) {
	result := longestIncreasingSubsequence(target)
	fmt.Printf("%v => %v\n", target, result)
}

func main() {
	runMain([]int{3, 10, 2, 1, 20})
	runMain([]int{50, 3, 10, 7, 40, 80})
	runMain([]int{5, 2, 8, 6, 3, 6, 9, 7})
}
