package main

// solutionValue := [][]int{}
// Use dynamic programming to find a solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func dynamicProgramming(items []Item, allowedWeight int) ([]Item, int, int) {
	numItems := len(items)
	solutionValue := make([][]Cell, numItems+1)
	// initialize solutionValue
	for i, _ := range solutionValue {
		itemNRow := make([]Cell, allowedWeight+1)
		for j, _ := range itemNRow {
			itemNRow[j] = Cell{value: 0, items: []Item{}}
		}
		solutionValue[i] = itemNRow
	}

	//begin work
	numberOfCalls := 0
	for i := 1; i <= numItems; i++ {
		currentItem := items[i-1]
		currentItem.isSelected = true
		for j := 1; j <= allowedWeight; j++ {
			numberOfCalls += 1
			if currentItem.weight <= j {
				remainingSpaceValue := 0
				var remaingingItems []Item = []Item{}
				if j-currentItem.weight >= 0 {
					cell := solutionValue[i-1][j-currentItem.weight]
					remainingSpaceValue = cell.value
					remaingingItems = cell.items
				}

				newValue := currentItem.value + remainingSpaceValue
				if newValue > solutionValue[i-1][j].value {
					solutionValue[i][j].value = newValue
					solutionValue[i][j].items = append(solutionValue[i][j].items, currentItem)
					solutionValue[i][j].items = append(remaingingItems, solutionValue[i][j].items...)
				} else {
					solutionValue[i][j] = solutionValue[i-1][j]
				}
			} else {
				solutionValue[i][j] = solutionValue[i-1][j]
			}
		}
	}

	solutionCell := solutionValue[numItems][allowedWeight]
	return solutionCell.items, solutionCell.value, numberOfCalls
}

type Cell struct {
	items []Item
	value int
}
