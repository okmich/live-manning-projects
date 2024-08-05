package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=5369s

// Write a function 'howSum(targetSum, numbers)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing any combineation of the elements that adds up to exactly the targetSum.
// If there is no combination that adds up to the targetSum, then return null

var howSumMemo = make(map[int][]int)

func howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		remResult := howSum(remainder, numbers)
		if remResult != nil {
			return append(remResult, num)
		}
	}
	return nil
}

func memoizeHowSum(targetSum int, numbers []int) []int {
	if val, ok := howSumMemo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		howSumMemo[remainder] = memoizeHowSum(remainder, numbers)

		if howSumMemo[remainder] != nil {
			return append(howSumMemo[remainder], num)
		}
	}

	howSumMemo[targetSum] = nil
	return nil
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
	fmt.Println(memoizeHowSum(300, []int{7, 14}))
	fmt.Println(memoizeHowSum(840, []int{7, 14}))
}
