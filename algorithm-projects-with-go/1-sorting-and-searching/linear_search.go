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

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	for index, value := range slice {
		numTests++
		if value == target {
			return index, numTests
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
		index, numTest := linearSearch(slice, intValue)
		if index >= 0 {
			fmt.Printf("values[%v] = %v, %v tests\n", index, slice[index], numTest)
		} else {
			fmt.Printf("Target %v not found, %v tests\n", intValue, numTest)
		}
	}
}
