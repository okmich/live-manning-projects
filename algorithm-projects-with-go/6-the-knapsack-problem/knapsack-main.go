package main

import (
	"fmt"
	"math/rand"
	"time"
)

// const numItems = 20 // A reasonable value for exhaustive search.
var numItems int = 45 // A reasonable value for branch-and-bound search.

const minValue = 1
const maxValue = 10
const minWeight = 4
const maxWeight = 10

var allowedWeight int

type Item struct {
	value, weight int
	isSelected    bool
	id, blockedBy int
	blockList     []int // Other items that this one blocks.
}

// func makeItems(numItems, minValue, maxValue, minWeight, maxWeight int) []Item {
// 	items := make([]Item, numItems)

// 	// random := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	random := rand.New(rand.NewSource(1337))
// 	for i := 0; i < numItems; i++ {
// 		value := random.Intn(maxValue-minValue+1) + minValue
// 		weight := random.Intn(maxWeight-minWeight+1) + minWeight

// 		items[i] = Item{value: value, weight: weight, isSelected: false}
// 	}
// 	return items
// }

// Make some random items.
func makeItems(numItems, minValue, maxValue, minWeight, maxWeight int) []Item {
	// Initialize a pseudorandom number generator.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	//random := rand.New(rand.NewSource(1337)) // Initialize with a fixed seed

	items := make([]Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = Item{
			id: i, blockedBy: -1, blockList: nil,
			value:      random.Intn(maxValue-minValue+1) + minValue,
			weight:     random.Intn(maxWeight-minWeight+1) + minWeight,
			isSelected: false,
		}
	}
	return items
}

// Return a copy of the items slice.
func copyItems(items []Item) []Item {
	newItems := make([]Item, len(items))
	copy(newItems, items)
	return newItems
}

// Return the total value of the items.
// If addAll is false, only add up the selected items.
func sumValues(items []Item, addAll bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if addAll || items[i].isSelected {
			total += items[i].value
		}
	}
	return total
}

// Return the total weight of the items.
// If addAll is false, only add up the selected items.
func sumWeights(items []Item, addAll bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if addAll || items[i].isSelected {
			total += items[i].weight
		}
	}
	return total
}

// Return the value of this solution.
// If the solution is too heavy, return -1 so we prefer an empty solution.
func solutionValue(items []Item, allowedWeight int) int {
	// If the solution's total weight > allowedWeight,
	// return 0 so we won't use this solution.
	if sumWeights(items, false) > allowedWeight {
		return -1
	}

	// Return the sum of the selected values.
	return sumValues(items, false)
}

// Print the selected items.
func printSelected(items []Item) {
	numPrinted := 0
	for i, item := range items {
		if item.isSelected {
			fmt.Printf("%d(%d, %d) ", i, item.value, item.weight)
		}
		numPrinted += 1
		if numPrinted > 100 {
			fmt.Println("...")
			return
		}
	}
	fmt.Println()
}

// Run the algorithm. Display the elapsed time and solution.
func runAlgorithm(alg func([]Item, int) ([]Item, int, int), items []Item, allowedWeight int) {
	// Copy the items so the run isn't influenced by a previous run.
	testItems := copyItems(items)

	start := time.Now()

	// Run the algorithm.
	solution, totalValue, functionCalls := alg(testItems, allowedWeight)

	elapsed := time.Since(start)

	fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
	printSelected(solution)
	fmt.Printf("Value: %d, Weight: %d, Calls: %d\n",
		totalValue, sumWeights(solution, false), functionCalls)
	fmt.Println()
}

func makeTestItems() []Item {
	items := make([]Item, 4)

	items[0] = Item{value: 50, weight: 8}
	items[1] = Item{value: 150, weight: 2}
	items[2] = Item{value: 210, weight: 6}
	items[3] = Item{value: 30, weight: 1}

	return items
}

func main() {
	fmt.Printf("Enter number of items: ")
	fmt.Scanln(&numItems)

	// items := makeTestItems()
	items := makeItems(numItems, minValue, maxValue, minWeight, maxWeight)
	allowedWeight = sumWeights(items, true) / 2

	// Display basic parameters.
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items count: %d\n", numItems)
	fmt.Printf("Total value: %d\n", sumValues(items, true))
	fmt.Printf("Total weight: %d\n", sumWeights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowedWeight)
	fmt.Println()

	// Exhaustive search
	if numItems <= 23 { // Only run exhaustive search if numItems <= 23.
		fmt.Println("*** Exhaustive Search ***")
		runAlgorithm(exhaustiveSearch, items, allowedWeight)
	}

	if numItems <= 45 {
		fmt.Println("*** Branch & Bound Search ***")
		runAlgorithm(branchAndBound, items, allowedWeight)
	}

	if numItems <= 85 {
		fmt.Println("*** Rod's technique ***")
		runAlgorithm(rodsTechnique, items, allowedWeight)
	}

	if numItems <= 350 {
		fmt.Println("*** Rod's sorted technique ***")
		runAlgorithm(rodsTechniqueSorted, items, allowedWeight)
	}

	fmt.Println("*** Dynamic programming ***")
	runAlgorithm(dynamicProgramming, items, allowedWeight)
}
