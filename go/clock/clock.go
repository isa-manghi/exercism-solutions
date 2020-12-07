package clock

import (
	"fmt"
)

// Clock is a struct that defines hours and minutes as integers
// these are later defined within a 24h period
type Clock struct {
	hours   int
	minutes int
}

// New is a constructor function.
// The purpose of constructor is to initialize the object of a class.
// The purpose of a method is to perform a task in Go.
func New(hours, minutes int) Clock {
	// - if minutes are greater than or equal to 60, must roll into hours
	// - if hours are equal or greater than 24, must start from 00:00
	// - if minutes or hours are negative then they must be converted into positive time

	for minutes < 0 { // if minutes is less than zero, then add 60 mins to make it positive
		minutes += 60 // add 60 mins until it is more than 0
		hours--       // minuses one hour with each iteration as it is a negative value
	}
	for hours < 0 { // if hours is less than zero, then add 24 to make it positive
		hours += 24
	}
	for minutes >= 60 {
		minutes -= 60 // minus 60 in each loop until it is less than 60
		hours++       // adds one to each iteration
	}
	if hours >= 24 {
		hours %= 24 //remainder to roll into hours
	}
	return Clock{hours: hours, minutes: minutes}
}

// String representation of Time. The "tostring" method.
// "get me a version of this object that becomes a string"
func (c Clock) String() string {
	return fmt.Sprintf("%02[1]d:%02[2]d", c.hours, c.minutes) // "%02d" gives zero padding
	// given one clock object, return a string 07:40
	// zero padding, so it is a two digit number
}

// Add minutes to time
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	return New(c.hours, c.minutes)
}

// Subtract minutes from time
func (c Clock) Subtract(minutes int) Clock {
	c.minutes -= minutes
	return New(c.hours, c.minutes)
}
