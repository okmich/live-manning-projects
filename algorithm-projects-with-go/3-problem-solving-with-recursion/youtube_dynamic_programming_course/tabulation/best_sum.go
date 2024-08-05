package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=14841s

// Write a function 'bestSum(targetSum, numbers)' that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing the shortest combineation of the elements that adds up to exactly the targetSum.
// If there is a tie for the shortest combination, you may return any of the shortest

// var howSumMemo = make(map[int][]int)

func bestSum(targetSum int, numbers []int) []int {
	tab := make([][]int, targetSum+1)
	for i, _ := range tab {
		tab[i] = nil
	}
	tab[0] = []int{}

	for i, _ := range tab {
		if tab[i] != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					newCurrentTab := append(tab[i], []int{num}...)
					if tab[i+num] == nil {
						tab[i+num] = newCurrentTab
					} else if len(newCurrentTab) < len(tab[i+num]) {
						tab[i+num] = newCurrentTab
					}
				}
			}
		}
	}
	return tab[targetSum]
}

func main() {

	fmt.Println(bestSum(4, []int{1, 3}))
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	fmt.Println(bestSum(12, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(52, []int{5, 3, 4, 7, 15}))
}
