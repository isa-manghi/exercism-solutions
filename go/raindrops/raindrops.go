package raindrops

import "strconv"

// Convert is a function that returns a string that contains raindrop sounds
// corresponding to certain potential factors.
func Convert (input int) string {

	var Pling = input % 3
	var Plang = input % 5
	var Plong = input % 7
	var outputString = ""

 	if Pling == 0 {
		outputString += "Pling"
	}
	if Plang == 0 {
		outputString += "Plang"
	}
	if Plong == 0 {
		outputString += "Plong"
	}
	if outputString == "" {
		return strconv.Itoa(input)
	}
	return outputString
}