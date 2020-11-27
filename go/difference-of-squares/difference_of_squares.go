package diffsquares

// SquareOfSum calculates the square of the sum of n
// n is an input which determines how many natural numbers are summed.
// When the natural numbers determined by n are summed, it will then be squared, thus giving the SquareOfSum.
func SquareOfSum(n int) int {
	var SumOfNaturalNumbers = 0
	for x := 1; x <= n; x++ {
		SumOfNaturalNumbers += x
	}
	return SumOfNaturalNumbers * SumOfNaturalNumbers
}

// SumOfSquares calculates the
// the natural numbers defined by n are all squared individually.
// the squares of n are then summed together and give us the SumOfSquares
func SumOfSquares(n int) int {
	var SumOfSquareOfNaturalNumbers = 0
	for x := 1; x <= n; x++ {
		SumOfSquareOfNaturalNumbers += x * x
	}
	return SumOfSquareOfNaturalNumbers
}

// Difference calculates the difference between SumOfSquares and SquareOfSum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
