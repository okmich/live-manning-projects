package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=6726s

// Write a function 'bestSum(targetSum, numbers)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing the shortest combineation of the elements that adds up to exactly the targetSum.
// If there is a tie for the shortest combination, you may return any of the shortest

// var howSumMemo = make(map[int][]int)

func bestSum(targetSum int, numbers []int) []int {
	return doBestSum(targetSum, numbers)
}

func doBestSum(targetSum int, numbers []int) []int {
	if targetSum < 0 {
		return nil
	}
	if targetSum == 0 {
		return []int{}
	}

	var currentBestSum []int
	for _, num := range numbers {
		remainder := targetSum - num
		remainderRes := doBestSum(remainder, numbers)
		if remainderRes != nil {
			res := append(remainderRes, num)
			if currentBestSum == nil || len(currentBestSum) > len(res) {
				currentBestSum = res
			}
		}
	}

	return currentBestSum
}

func doMemoizeBestSum(memo map[int][]int, targetSum int, numbers []int) []int {
	if value, ok := memo[targetSum]; ok {
		return value
	}
	if targetSum < 0 {
		return nil
	}
	if targetSum == 0 {
		return []int{}
	}

	var currentBestSum []int
	for _, num := range numbers {
		remainder := targetSum - num
		remainderRes := doMemoizeBestSum(memo, remainder, numbers)
		if remainderRes != nil {
			res := append(remainderRes, num)
			if currentBestSum == nil || len(currentBestSum) > len(res) {

				currentBestSum = res
			}
		}
	}

	memo[targetSum] = currentBestSum
	return currentBestSum
}

func main() {

	fmt.Println(bestSum(4, []int{1, 3}))
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(12, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(51, []int{5, 3, 4, 7, 15}))

	memo := make(map[int][]int)
	fmt.Println(doMemoizeBestSum(memo, 4, []int{1, 3}))
	fmt.Println(doMemoizeBestSum(memo, 7, []int{5, 3, 4, 7}))
	fmt.Println(doMemoizeBestSum(memo, 12, []int{5, 3, 4, 7}))
	fmt.Println(doMemoizeBestSum(memo, 51, []int{5, 3, 4, 7, 15}))
}
