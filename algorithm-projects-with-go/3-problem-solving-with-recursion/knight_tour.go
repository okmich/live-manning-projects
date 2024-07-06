package main

import (
	"fmt"
	"time"
)

// The board dimensions.
const numRows = 8
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

var moveOffsets []Offset

var numCalls int64

// Fill the Offset slice.
func initializeOffsets() {
	moveOffsets = []Offset{
		Offset{-2, -1},
		Offset{-1, -2},
		Offset{+2, -1},
		Offset{+1, -2},
		Offset{-2, +1},
		Offset{-1, +2},
		Offset{+2, +1},
		Offset{+1, +2},
	}
}

// Make a board filled with -1s.
func makeBoard(numRows, numCols int) [][]int {
	board := make([][]int, numRows)
	for r := range board {
		board[r] = make([]int, numCols)
		for c := 0; c < numCols; c++ {
			board[r][c] = unvisited
		}
	}
	return board
}

func dumpBoard(board [][]int) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%02d ", board[r][c])
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

// Try to extend a knight's tour starting at (curRow, curCol).
// Return true or false to indicate whether we have found a solution.
func findTour(board [][]int, numRows, numCols, curRow, curCol, numVisited int) bool {
	numCalls++
	if numVisited == numRows*numCols {
		if !requireClosedTour {
			return true
		} else {
			for _, offset := range moveOffsets {
				r := curRow + offset.dr
				c := curCol + offset.dc

				if r < 0 || r >= numRows || c < 0 || c >= numCols {
					continue
				}

				if board[r][c] == 0 {
					return true
				}
			}
			return false
		}
	}

	for _, offset := range moveOffsets {
		// Get the move.
		r := curRow + offset.dr
		c := curCol + offset.dc

		// See if this move is on the board or if we have already visited this position.
		if r < 0 || r >= numRows || c < 0 || c >= numCols || board[r][c] != unvisited {
			continue
		}

		// The move to [r][c] is viable.
		board[r][c] = numVisited

		// If we succeed, return true.
		if findTour(board, numRows, numCols, r, c, numVisited+1) {
			return true
		}

		// We did not find a tour with this move. Unmake this move.
		board[r][c] = unvisited
	}

	return false
}

func main() {
	numCalls = 0

	// Initialize the move offsets.
	initializeOffsets()

	// Create the blank board.
	board := makeBoard(numRows, numCols)

	// Try to find a tour.
	start := time.Now()
	board[0][0] = 0
	if findTour(board, numRows, numCols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dumpBoard(board)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", numCalls)
}
