package main

// https://www.youtube.com/watch?v=oBt53YbR9Kk&list=WL&index=26&t=2319s

// Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal is to travel to the bottom-right corner.
// Your may only move down or right.

// In how many ways can you travel to the goal on a grid with dimension m*n
// Write a function gridTraveler(m, n) that calculates this.

import (
	"fmt"
	"time"
)

var memo = make(map[string]int)

func gridTraveler(rows int, cols int) int {
	//base cases
	if rows == 1 && cols == 1 {
		return 1
	}
	if rows == 0 || cols == 0 {
		return 0
	}
	return gridTraveler(rows-1, cols) + gridTraveler(rows, cols-1)
}

func memoizedGridTraveler(rows int, cols int) int {
	key := fmt.Sprint("%v,%v", rows, cols)
	if value, ok := memo[key]; ok {
		return value
	} else {
		//base cases
		if rows == 1 && cols == 1 {
			return 1
		}
		if rows == 0 || cols == 0 {
			return 0
		}
		memo[key] = memoizedGridTraveler(rows-1, cols) + memoizedGridTraveler(rows, cols-1)
		return memo[key]
	}
}

func main() {
	rows := []int{1, 3, 5, 8, 13, 18, 25}
	cols := []int{1, 8, 13, 18, 25}

	for _, r := range rows {
		for _, c := range cols {
			start := time.Now()
			res := memoizedGridTraveler(r, c)
			end := time.Since(start)
			fmt.Printf("gridTraveler(%v, %v) ==>  %v. Time taken: %.4f secs.\n", r, c, res, end.Seconds())
		}
	}
}
