package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid checks if the input is a Luhn number or not.
// First every second digit must be doubled.
// If the doubled number is greater than 9, then 9 must be subtracted.
// Then all digits must be summed.
// if the sum is divisible by 10, it is a luhn number.
func Validv1(input string) bool {
	input = strings.TrimSpace(input)
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("err: ", err)
	}
	var SecondDigit int
	var FirstDigit int
	// Andrew: Ok, so you're looping through each character in the input string here, but what is your intention to achieve within this loop?
	// Andrew: Like, which bit of your algorithm do you intend this bit to be?
	// Andrew: It is probably worth writing a comment here above this loop explaining what /this bit/ is supposed to be doing.
	for x := 1; x <= len(input); x++ {
		// Andrew: This if statement is checking whether the length of the input string is less than/equal to 1 character long?
		// Andrew: Is that what you intend? If so, why does this need to be inside your for loop?
		// Andrew: The conditional will return the same regardless of which loop iteration you're on, so right now it will be redundantly repeated.
		if len(input) <= 1 {
			return false
		}
		// Andrew: So, if x is an even number (where x is the variable you chose to be your loop iteration index)...
		if x%2 == 0 {
			// Andrew: ...set the value of the SecondDigit variable to your loop iterator, squared?
			// Andrew: I suspect this isn't what you mean.
			// Andrew: If you intended this to be the part of your algorithm which says you should double every second digit,
			// Andrew: you probably meant to use your "x" variable as an index to /reference/ a specific character in your "input" variable
			// Andrew: and you probably meant to /update/ the character in your input string to be the result of doubling that value,
			// Andrew: which would result in your "input" variable being modified
			SecondDigit = x * 2
		} else {
			// Andrew: Not sure what this was meant to be? FirstDigit is an empty int variable here and has never been assigned a value.
			// Andrew: x is your for loop iterator, so setting that to a different value is probably never what you want anyway.
			// Andrew: Which part of your algorithm was this else block supposed to be achieving?
			// Andrew: I think some comments, from yourself to yourself, explaining your intentions with every line you write, would go a long way.
			x = FirstDigit
		}
	}
	// Andrew: So, since this is outside your for loop, you're saying was the /last value/ of the "SecondDigit" variable greater than 9?
	// Andrew: you probably meant for this to be inside your loop, so the logic would be applied after your attempted doubling of those second digits.
	// Andrew: As a feedback point; indentation is key, and it's really important that you pay more attention to those curly braces,
	// Andrew: your indentation and program flow so you think more about where bits of your logic should be.
	// Andrew: Blank lines between blocks, and comments above things explaining intention will help too!
	if SecondDigit > 9 {
		SecondDigit = SecondDigit - 9
	}
	// Andrew: Even if these variables actually had values in them which matched the variable names (they don't, due to other logic flaws highlighted above);
	// Andrew: I don't believe this if statement corresponds to any part of your algorithm statement at the top of the function.
	// Andrew: if this is supposed to be "if the sum is divisible by 10" - where's the "all digits must be summed"? This would only ever be summing two numbers?
	if FirstDigit+SecondDigit%10 == 0 {
		return true
	}
	return false
}
