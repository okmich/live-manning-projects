package main

import "fmt"

// Write a function 'allSums(targetSum, numbers)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing all combineations of the elements that adds up to exactly the targetSum.

// var howSumMemo = make(map[int][]int)

func allSums(targetSum int, numbers []int) [][]int {
	if targetSum < 0 {
		return nil
	}
	if targetSum == 0 {
		return [][]int{[]int{}}
	}

	result := [][]int{}
	for _, num := range numbers {
		remainder := targetSum - num
		allSumWays := allSums(remainder, numbers)
		if allSumWays != nil {
			for i, _ := range allSumWays {
				allSumWays[i] = append(allSumWays[i], num)
			}
		}
		result = append(result, allSumWays...)
	}
	return result
}

func main() {
	fmt.Println(allSums(3, []int{1, 2}))
	fmt.Println(allSums(4, []int{1, 3}))
	fmt.Println(allSums(7, []int{5, 3, 4, 7}))
	fmt.Println(allSums(12, []int{5, 3, 4, 7}))
	fmt.Println(allSums(51, []int{5, 3, 4, 7, 15}))
}
