package grains

import (
	"errors"
	"math"
)

// Square returns how many grains are on a specific square of a chessboard.
// the number of grains doubles with each successive square.
func Square(square int) (uint64, error) {
	if square > 0 && square <= 64 {
		return uint64(math.Pow(2, float64(square-1))), nil
	}
	return 0, errors.New("error")
}

// Total returns the total number of grains on a chessboard.
// double the amount of grains with each successive square
// add all of that up
func Total() uint64 {
	var TotalGrainsOnTheChessboard uint64
	for i := 1; i <= 64; i++ {
		AddToTot, _ := Square(i)
		TotalGrainsOnTheChessboard += AddToTot
	}
	return TotalGrainsOnTheChessboard
}
