package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=13079s

// Write a function 'canSum(targetSum, numbers)' that takes in a targetSum and an array of numbers as argument
// The function shoudl return a boolean indicating whether or not it is possible to generate the targetSum
// using numbers from the array

func canSum(targetSum int, numbers []int) bool {
	tab := make([]bool, targetSum+1)
	tab[0] = true

	for i, v := range tab {
		if v {
			for _, num := range numbers {
				if i+num <= targetSum {
					tab[i+num] = true
				}
			}
		}

		if tab[targetSum] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canSum(2, []int{}))           //false
	fmt.Println(canSum(3, []int{1, 2}))       //true
	fmt.Println(canSum(4, []int{3, 2}))       //true
	fmt.Println(canSum(4, []int{1, 3}))       //true
	fmt.Println(canSum(10, []int{1, 3, 6}))   //true
	fmt.Println(canSum(7, []int{2, 4}))       //false
	fmt.Println(canSum(7, []int{5, 3, 4, 7})) //true
	fmt.Println(canSum(300, []int{7, 14}))    //false
	fmt.Println(canSum(840, []int{7, 14}))    //true
}
