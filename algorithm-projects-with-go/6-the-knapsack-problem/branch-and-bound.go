package main

// Use branch and bound to find a solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func branchAndBound(items []Item, allowedWeight int) ([]Item, int, int) {
	bestValue := 0
	currentValue := 0
	currentWeight := 0
	remainingValue := sumValues(items, true)

	return doBranchAndBound(items, allowedWeight, 0,
		bestValue, currentValue, currentWeight, remainingValue)
}

func doBranchAndBound(items []Item, allowedWeight, nextIndex,
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
	var test1Value int
	var test1Calls int
	if currentWeight+items[nextIndex].weight <= allowedWeight {
		items[nextIndex].isSelected = true
		test1Solution, test1Value, test1Calls =
			doBranchAndBound(items, allowedWeight, nextIndex+1, bestValue,
				currentValue+items[nextIndex].value,
				currentWeight+items[nextIndex].weight,
				remainingValue-items[nextIndex].value)
		if test1Value > bestValue {
			bestValue = test1Value
		}
	} else {
		test1Solution = nil
		test1Value = 0
		test1Calls = 1
	}

	// Try not adding the next item.
	var test2Solution []Item
	var test2Value int
	var test2Calls int
	// See if there is a chance of improvement without this item's value.
	if currentValue+remainingValue-items[nextIndex].value > bestValue {
		items[nextIndex].isSelected = false
		test2Solution, test2Value, test2Calls =
			doBranchAndBound(items, allowedWeight, nextIndex+1, bestValue,
				currentValue, currentWeight,
				remainingValue-items[nextIndex].value)
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
