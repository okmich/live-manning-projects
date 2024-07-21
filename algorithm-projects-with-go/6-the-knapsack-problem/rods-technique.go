package main

import "sort"

// Build the items' block lists.
func makeBlockLists(items []Item) {
	for i := range items {
		items[i].blockList = []int{}
		for j := range items {
			if i != j && items[i].value >= items[j].value && items[i].weight <= items[j].weight {
				items[i].blockList = append(items[i].blockList, items[j].id)
			}
		}
	}
}

// Block items on this item's blocks list.
func blockItems(source Item, items []Item) {
	for _, otherId := range source.blockList {
		if items[otherId].blockedBy < 0 {
			items[otherId].blockedBy = source.id
		}
	}
}

// Unblock items on this item's blocks list.
func unblockItems(source Item, items []Item) {
	for _, otherId := range source.blockList {
		if items[otherId].blockedBy == source.id {
			items[otherId].blockedBy = -1
		}
	}
}

// Use branch and bound to find a solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func rodsTechnique(items []Item, allowedWeight int) ([]Item, int, int) {
	makeBlockLists(items)

	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := sumValues(items, true)

	return doRodsTechnique(items, allowedWeight, 0,
		bestValue, currentValue, currentWeight, remainingValue)
}

// Use Rod's technique sorted to find a solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func rodsTechniqueSorted(items []Item, allowedWeight int) ([]Item, int, int) {
	makeBlockLists(items)

	// Sort so items with longer blocked lists come first.
	sort.Slice(items, func(i, j int) bool {
		return len(items[i].blockList) > len(items[j].blockList)
	})

	// Reset the items' IDs.
	for i := range items {
		items[i].id = i
	}

	// Rebuild the blocked lists with the new indices.
	makeBlockLists(items)

	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := sumValues(items, true)

	return doRodsTechnique(items, allowedWeight, 0,
		bestValue, currentValue, currentWeight, remainingValue)
}

func doRodsTechnique(items []Item, allowedWeight, nextIndex,
	bestValue, currentValue, currentWeight, remainingValue int) ([]Item, int, int) {
	// See if we have a full assignment.
	if nextIndex >= len(items) {
		// Make a copy of this assignment.
		solution := copyItems(items)

		// Return the assignment and its total value.
		return solution, currentValue, 1
	}

	// We do not have a full assignment.
	// See if we can improve this solution enough to be worth persuing.
	if currentValue+remainingValue <= bestValue {
		// We cannot improve on the best solution found so far.
		return nil, currentValue, 1
	}

	// Try adding the next item.
	var test1Solution []Item
	test1Solution = nil
	test1Value := 0
	test1Calls := 1

	// See if the item is blocked.
	if items[nextIndex].blockedBy < 0 {
		if currentWeight+items[nextIndex].weight <= allowedWeight {
			items[nextIndex].isSelected = true
			test1Solution, test1Value, test1Calls =
				doRodsTechnique(items, allowedWeight, nextIndex+1, bestValue,
					currentValue+items[nextIndex].value,
					currentWeight+items[nextIndex].weight,
					remainingValue-items[nextIndex].value)
			if test1Value > bestValue {
				bestValue = test1Value
			}
		}
	}

	// Try not adding the next item.
	var test2Solution []Item
	var test2Value int
	var test2Calls int
	// See if there is a chance of improvement without this item's value.
	if currentValue+remainingValue-items[nextIndex].value > bestValue {
		blockItems(items[nextIndex], items)
		test2Solution, test2Value, test2Calls =
			doRodsTechnique(items, allowedWeight, nextIndex+1, bestValue,
				currentValue, currentWeight,
				remainingValue-items[nextIndex].value)
		unblockItems(items[nextIndex], items)
	} else {
		test2Solution = nil
		test2Value = 0
		test2Calls = 1
	}

	// Return the solution that is better.
	if test1Value >= test2Value {
		return test1Solution, test1Value, test1Calls + test2Calls + 1
	} else {
		return test2Solution, test2Value, test1Calls + test2Calls + 1
	}
}
