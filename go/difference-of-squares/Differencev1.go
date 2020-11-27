package diffsquares

// SquareOfSum calculates the square of the sum of n
// n is an input which determines how many natural numbers are summed.
// When the natural numbers determined by n are summed, it will then be squared, thus giving the SquareOfSum.
func SquareOfSumv1(n int) int {
	// look at n
	// write out each number within n (from one to n)
	// add them all together individually, one by one.
	// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 + 11 + 12 + 13 + ...
	//   3   6    10   15 ...
	// add all of the natural numbers defined by n.
	var SumOfNaturalNumbers = 0
	for x := 1; x <= n; x++ {
		SumOfNaturalNumbers += x
	}
	return SumOfNaturalNumbers * SumOfNaturalNumbers
}

// SumOfSquares calculates the
// the natural numbers defined by n are all squared individually.
// the squares of n are then summed together and give us the SumOfSquares
func SumOfSquaresv1(n int) int {
	// write out each number within n (from one to n)
	// square them individually
	// add up all the squares
	var SumOfSquareOfNaturalNumbers = 0
	for x := 1; x <= n; x++ {
		SumOfSquareOfNaturalNumbers += x * x
	}
	return SumOfSquareOfNaturalNumbers
}

// Difference calculates the difference between SumOfSquares and SquareOfSum
func Differencev1(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
