package main

import "fmt"

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=4003s

// Write a function 'canSum(targetSum, numbers)' that takes in a targetSum and an array of numbers as argument
// The function shoudl return a boolean indicating whether or not it is possible to generate the targetSum
// using numbers from the array

var memo = make(map[int]bool)

func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	} else if targetSum < 0 {
		return false
	}

	var remainder int
	var res bool
	for _, num := range numbers {
		remainder = targetSum - num
		res = canSum(remainder, numbers)
		if res {
			return true
		}
	}
	return false
}

func memoizeCanSum(targetSum int, numbers []int) bool {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return true
	} else if targetSum < 0 {
		return false
	}

	var remainder int
	var res bool
	for _, num := range numbers {
		remainder = targetSum - num
		res = memoizeCanSum(remainder, numbers)
		if res {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}

func main() {
	fmt.Println(canSum(2, []int{}))                //false
	fmt.Println(canSum(3, []int{1, 2}))            //true
	fmt.Println(canSum(4, []int{3, 2}))            //true
	fmt.Println(canSum(4, []int{1, 3}))            //true
	fmt.Println(canSum(10, []int{1, 3, 6}))        //true
	fmt.Println(canSum(7, []int{2, 4}))            //false
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))      //true
	fmt.Println(memoizeCanSum(3000, []int{7, 14})) //false
	fmt.Println(memoizeCanSum(8400, []int{7, 14})) //true
}
