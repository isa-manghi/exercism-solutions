package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the input is a Luhn number or not.
// Len(input) must be greater than one.
// First every second digit must be doubled starting from the right.
// If the doubled number is greater than 9, then 9 must be subtracted.
// Then all digits must be summed.
// if the sum is divisible by 10, it is a luhn number.
func Valid(input string) bool {
	var SumOfDigit int64
	input = strings.Replace(input, " ", "", -1) // strips spaces from the input string.
	if len(input) <= 1 {                        //if the length of the input is less than or equal to one, end.
		return false
	}
	// make new variable, increment by 1 from the right going to the left.
	for position := len(input) - 1; position >= 0; position-- {
		numsFromRightCount := len(input) - position // This helps identify the second digit from the right, by reversing the index.
		CharacterCodeAtCurrentPosition := input[position]
		Digit, err := strconv.ParseInt(string(CharacterCodeAtCurrentPosition), 10, 32) // converts CharacterCodeAtCurrentPosition into an integer
		if err != nil {
			return false // if CharacterCodeAtCurrentPosition is not an int, end.
		}
		if numsFromRightCount%2 == 0 { // if position is divisible by 2 (because it will identify every second digit)
			Digit *= 2 // then times Digit by 2 (double it) and assign new value to Digit.
		}
		if Digit > 9 { // if digit is more than 9
			Digit -= 9 // deduct 9 from Digit and assign new value to Digit
		}
		SumOfDigit += Digit // Add Digit to SumOfDigit
	}
	if SumOfDigit%10 == 0 { // if SumOfDigit is divisible by 10 then its a Luhn
		return true
	}
	return false
}
