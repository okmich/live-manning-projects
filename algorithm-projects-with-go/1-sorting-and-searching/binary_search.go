package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	var randomSlice []int
	randomSlice = make([]int, max)
	for i := 0; i < max; i++ {
		randomSlice[i] = rand.Intn(max)
	}
	return randomSlice
}

func printSlice(slice []int, numItems int) {
	var numberToPrint int = min(len(slice), numItems)
	for i := 0; i < numberToPrint; i++ {
		fmt.Printf("%v ", slice[i])
	}
	fmt.Println()
}

func partition(slice []int) int {
	hi := len(slice)
	pivot := slice[hi-1]

	tempIndex := -1
	for i := 0; i < hi-1; i++ {
		if slice[i] <= pivot {
			tempIndex = tempIndex + 1
			slice[i], slice[tempIndex] = slice[tempIndex], slice[i]
		}
	}
	tempIndex = tempIndex + 1
	slice[tempIndex], slice[hi-1] = slice[hi-1], slice[tempIndex]
	return tempIndex
}

func quicksort(slice []int) {
	//https://en.wikipedia.org/wiki/Quicksort
	if len(slice) <= 1 {
		return
	}
	p := partition(slice[0:len(slice)])
	if p > 1 {
		quicksort(slice[0:p])
	}
	if p <= len(slice)-1 {
		quicksort(slice[p+1 : len(slice)])
	}
}

// Perform binary search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearch(slice []int, target int) (index, numTests int) {
	start := 0
	end := len(slice) - 1

	for start <= end {
		mid := start + (end-start)/2
		numTests++
		if slice[mid] > target {
			end = mid - 1
		} else if slice[mid] < target {
			start = mid + 1
		} else {
			return mid, numTests
		}
	}
	return -1, numTests
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	quicksort(slice)
	printSlice(slice, 40)
	fmt.Println()

	var userInput string
	fmt.Println("\n\n***Searching begins***")
	for {
		userInput = ""
		fmt.Printf("Target: ")
		fmt.Scanln(&userInput)
		if userInput == "" {
			break
		}
		intValue, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(err)
			break
		}
		index, numTest := binarySearch2(slice, intValue)
		if index >= 0 {
			fmt.Printf("values[%v] = %v, %v tests\n", index, slice[index], numTest)
		} else {
			fmt.Printf("Target %v not found, %v tests\n", intValue, numTest)
		}
	}
}
