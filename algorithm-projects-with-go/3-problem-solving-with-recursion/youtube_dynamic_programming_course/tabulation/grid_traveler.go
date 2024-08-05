package main

import (
	"fmt"
	"time"
)

// https://www.youtube.com/watch?v=oBt53YbR9Kk&t=12137s

// Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal is to travel to the bottom-right corner.
// Your may only move down or right.

// In how many ways can you travel to the goal on a grid with dimension m*n
// Write a function gridTraveler(m, n) that calculates this.

func gridTraveler(rows uint8, cols uint8) uint8 {
	tab := make([][]uint8, rows+1)
	// initialize with solution values
	for i := range tab {
		tab[i] = make([]uint8, cols+1)
		for j := range tab[i] {
			// if i and j are both 0, then 0, else
			// if i > 0 and j = 0, then i = tab[i-1][0]
			// else tab[i][j] = tab[i-1] + 1
			if i == 0 || j == 0 {
				tab[i][j] = 0
			} else {
				if i == 1 && j == 1 {
					tab[i][j] = 1
				} else {
					tab[i][j] = tab[i-1][j] + tab[i][j-1]
				}
			}
		}
	}
	return tab[rows][cols]
}

func main() {
	rows := []uint8{1, 3, 5, 8, 13, 18, 25}
	cols := []uint8{1, 8, 13, 18, 25}

	for _, r := range rows {
		for _, c := range cols {
			start := time.Now()
			res := gridTraveler(r, c)
			end := time.Since(start)
			fmt.Printf("gridTraveler(%v, %v) ==>  %v. Time taken: %.4f secs.\n", r, c, res, end.Seconds())
		}
	}
}
