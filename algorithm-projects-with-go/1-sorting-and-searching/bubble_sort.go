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

func bubbleSort(slice []int) {
	//https://en.wikipedia.org/wiki/Bubble_sort
	n := len(slice)
	for n >= 1 {
		newn := 0
		for j := 1; j < n; j++ {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
				newn = j
			}
		}
		n = newn
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
