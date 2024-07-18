package main

// Recursively assign values in or out of the solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func exhaustiveSearch(items []Item, allowedWeight int) ([]Item, int, int) {
	return doExhaustiveSearch(items, allowedWeight, 0)
}

func doExhaustiveSearch(items []Item, allowedWeight, nextIndex int) ([]Item, int, int) {
	if nextIndex >= len(items) {
		copiedItems := copyItems(items)
		solutionValue := solutionValue(copiedItems, allowedWeight)
		return copiedItems, solutionValue, 1
	}
	// We do not have a full assignment.
	// Try adding the next item.
	items[nextIndex].isSelected = true

	// Recursively call the function.
	items1, test1Value, noFuncCall1 := doExhaustiveSearch(items, allowedWeight, nextIndex+1)

	// Try not adding the next item.
	items[nextIndex].isSelected = false

	// Recursively call the function.
	items2, test2Value, noFuncCall2 := doExhaustiveSearch(items, allowedWeight, nextIndex+1)

	// Return the solution that is better.
	if test1Value >= test2Value {
		return items1, test1Value, noFuncCall1 + noFuncCall2 + 1
	} else {
		return items2, test2Value, noFuncCall1 + noFuncCall2 + 1
	}
}
