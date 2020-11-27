package reverse

// Reverse reverses the inputted string using a for loop.
func Reverse(inputString string) string {
	var OutputString = ""

	// Iterate over the letter in the inputString backwards
	for x := len(inputString) - 1; x >= 0; x-- {
		// Concatenate each letter of the inputString onto OutputString
		OutputString += string(inputString[x])
	}
	return OutputString
}
// cabbage = inoutsting
// len inputstring = 7
// x = 6
