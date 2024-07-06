package main

import (
	"fmt"
	"math/rand"
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

// Make a slice containing pseudorandom numbers in [0, max).
func printSlice(slice []int, numItems int) {
	var numberToPrint int = min(len(slice), numItems)
	for i := 0; i < numberToPrint; i++ {
		fmt.Printf("%v ", slice[i])
	}
	fmt.Println()
}

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	size := len(slice)
	isSorted := true
	for i := 1; i < size; i++ {
		if slice[i-1] > slice[i] {
			isSorted = false
			break
		}
	}
	if isSorted {
		fmt.Println("The slice is sorted")
	} else {
		fmt.Println("The slice is NOT sorted!")
	}
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

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
