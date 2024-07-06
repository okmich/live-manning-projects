package main

import (
	"fmt"
	"math/rand"
)

type Customer struct {
	id           string
	numPurchases int
}

// Make a slice containing pseudorandom numbers in [0, max).
func makeRandomCustomerSlice(numItems, max int) []Customer {
	var randomSlice []Customer
	randomSlice = make([]Customer, max)
	for i := 0; i < max; i++ {
		randomSlice[i] = Customer{id: fmt.Sprintf("C%v", i), numPurchases: rand.Intn(max)}
	}
	return randomSlice
}

// Print at most numItems items.
func printCustomerSlice(slice []Customer, numItems int) {
	var numberToPrint int = min(len(slice), numItems)
	for i := 0; i < numberToPrint; i++ {
		fmt.Printf("%v ", slice[i])
	}
	fmt.Println()
}

// Verify that the slice is sorted.
func checkCustomerSorted(slice []Customer) {
	size := len(slice)
	isSorted := true
	for i := 1; i < size; i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
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

func countingSort(slice []Customer, max int) []Customer {
	// https://www.geeksforgeeks.org/counting-sort/
	var arraySize int = len(slice)
	var outputArr []Customer = make([]Customer, arraySize)
	var countArray []int = make([]int, max+1)
	for i := 0; i < max+1; i++ {
		countArray[i] = 0
	}

	for i := 0; i < arraySize; i++ {
		countArray[slice[i].numPurchases] = countArray[slice[i].numPurchases] + 1
	}

	for i := 1; i < max+1; i++ {
		countArray[i] = countArray[i-1] + countArray[i]
	}

	for i := arraySize - 1; i >= 0; i-- {
		outputArr[countArray[slice[i].numPurchases]-1] = slice[i]
		countArray[slice[i].numPurchases] = countArray[slice[i].numPurchases] - 1
	}

	return outputArr
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomCustomerSlice(numItems, max)
	printCustomerSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printCustomerSlice(sorted, 40)

	// Verify that it's sorted.
	checkCustomerSorted(sorted)
}
