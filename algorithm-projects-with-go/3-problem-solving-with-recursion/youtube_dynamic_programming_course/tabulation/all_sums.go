package main

import "fmt"

// Write a function 'allSums(targetSum, numbers)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing all combineations of the elements that adds up to exactly the targetSum.

func allSums(targetSum int, numbers []int) [][]int {
	tabs := make([][][]int, targetSum+1)
	for i, _ := range tabs {
		tabs[i] = nil
	}
	tabs[0] = [][]int{}

	for i, _ := range tabs {
		if tabs[i] != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					if tabs[i+num] == nil && len(tabs[i]) == 0 {
						tabs[i+num] = [][]int{[]int{num}}
					} else {
						for k, _ := range tabs[i] {
							tabs[i+num] = append(tabs[i+num], append(tabs[i][k], num))
						}
					}
				}
			}
		}
	}
	return tabs[targetSum]
}

func main() {
	fmt.Println(3, " == ", allSums(3, []int{1, 2}))
	fmt.Println(4, " == ", allSums(4, []int{1, 3}))
	fmt.Println(7, " == ", allSums(7, []int{5, 3, 4, 7}))
	fmt.Println(12, " == ", allSums(12, []int{5, 3, 4, 7}))
	// fmt.Println(51, " == ", allSums(51, []int{5, 3, 4, 7, 15}))
}
