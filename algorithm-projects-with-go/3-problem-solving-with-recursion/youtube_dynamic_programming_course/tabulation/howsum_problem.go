package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=13981s

// Write a function 'howSum(targetSum, numbers)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing any combineation of the elements that adds up to exactly the targetSum.
// If there is no combination that adds up to the targetSum, then return null

func howSum(targetSum int, numbers []int) []int {
	tab := make([][]int, targetSum+1)
	// initialize with emtpy array
	for i, _ := range tab {
		tab[i] = nil
	}
	// seed the first
	tab[0] = []int{}

	totalSum := 0
	// start the process
	for i, _ := range tab {
		for _, num := range numbers {
			if tab[i] != nil && i+num <= targetSum {
				tab[i+num] = append(tab[i], num)

				// check the sum
				for _, v := range tab[i+num] {
					totalSum += v
				}
				if totalSum == targetSum {
					return tab[i+num]
				} else {
					totalSum = 0
				}
			}
		}
	}

	return tab[targetSum]
}

func main() {
	fmt.Println(howSum(2, []int{}))
	fmt.Println(howSum(3, []int{1, 2}))
	fmt.Println(howSum(5, []int{3, 2}))
	fmt.Println(howSum(4, []int{3, 1}))
	fmt.Println(howSum(10, []int{4, 3, 6}))
	fmt.Println(howSum(7, []int{2, 4}))
	fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum(300, []int{7, 14}))
	fmt.Println(howSum(840, []int{7, 14}))
}
