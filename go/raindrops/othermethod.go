package raindrops

import "strconv"

// Convert is a function that returns a string that contains raindrop sounds
// corresponding to certain potential factors.
func Convertv2 (input int) string {
	// Define the output string variable which will be built up by the rest of the logic
	var outputString = ""

	// Define a map of the factors -> strings according to the Exercism challenge
	var factorStrings = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	// Iterate through all of the values in the factorStrings map
	for key, value := range factorStrings {
		// If the input value modulo of this factor is 0, append the value string to outputString
		if input % key == 0 {
			outputString += value
		}
	}

	//var Pling = input % 3
	//var Plang = input % 5
	//var Plong = input % 7
	//
	//if Pling == 0 {
	//	outputString +=
	//}
	//if Plang == 0 {
	//	outputString += "Plang"
	//}
	//if Plong == 0 {
	//	outputString += "Plong"
	//}


	if outputString == "" {
		return strconv.Itoa(input)
	}
	return outputString
}
