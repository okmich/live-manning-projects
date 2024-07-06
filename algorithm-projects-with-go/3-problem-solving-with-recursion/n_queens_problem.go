package main

import (
	"fmt"
	"time"
)

// Make a board filled with periods.
func makeBoard(numRows int) [][]string {
	numCols := numRows
	board := make([][]string, numRows)
	for r := range board {
		board[r] = make([]string, numCols)
		for c := 0; c < numCols; c++ {
			board[r][c] = "."
		}
	}
	return board
}

// Display the board.
func dumpBoard(board [][]string) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%s ", board[r][c])
		}
		fmt.Println()
	}
}

// Return true if this series of squares contains at most one queen.
func seriesIsLegal(board [][]string, r0, c0, dr, dc int) bool {
	numRows := len(board)
	numCols := numRows
	hasQueen := false

	r := r0
	c := c0
	for {
		if board[r][c] == "Q" {
			// If we already have a queen on this row,
			// then this board is not legal.
			if hasQueen {
				return false
			}

			// Remember that we have a queen on this row.
			hasQueen = true
		}

		// Move to the next square in the series.
		r += dr
		c += dc

		// If we fall off the board, then the series is legal.
		if r >= numRows || c >= numCols || r < 0 || c < 0 {
			return true
		}
	}
}

// Return true if the board is legal.
func boardIsLegal(board [][]string) bool {
	numRows := len(board)

	// See if each row is legal.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, 0, 0, 1) {
			return false
		}
	}

	// See if each column is legal.
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, 0) {
			return false
		}
	}

	// See if diagonals down to the right are legal.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, 0, 1, 1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, 1) {
			return false
		}
	}

	// See if diagonals down to the left are legal.
	for r := 0; r < numRows; r++ {
		if !seriesIsLegal(board, r, numRows-1, 1, -1) {
			return false
		}
	}
	for c := 0; c < numRows; c++ {
		if !seriesIsLegal(board, 0, c, 1, -1) {
			return false
		}
	}

	// If we survived this long, then the board is legal.
	return true
}

// Return true if the board is legal and a solution.
func boardIsASolution(board [][]string) bool {
	// See if it is legal.
	if !boardIsLegal(board) {
		return false
	}

	// See if the board contains exactly numRows queens.
	numRows := len(board)
	numQueens := 0
	for r := 0; r < numRows; r++ {
		for c := 0; c < numRows; c++ {
			if board[r][c] == "Q" {
				numQueens += 1
			}
		}
	}
	return numQueens == numRows
}

func placeQueens1(board [][]string, numRows, r, c int) bool {

	return false
}

func main() {
	const numRows = 5
	board := makeBoard(numRows)

	start := time.Now()
	success := placeQueens1(board, numRows, 0, 0)
	//success := placeQueens2(board, numRows, 0, 0, 0)
	// success := placeQueens3(board, numRows, 0, 0, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dumpBoard(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
