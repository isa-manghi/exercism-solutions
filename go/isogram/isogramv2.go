package isogram

import "strings"

//isItAnIsogram determines if the input is an isogram.
//look at the word, look at all of the letters in the word
// check if the letters repeat
func IsIsogramv2(input string) bool {
	input = strings.ToUpper(input)
	isogramMap := make(map[string]bool)
	for _, letter := range input { //lookup in map
		if isogramMap[string(letter)] == true && letter < 90 { //ok checks if the value exists in the map
			return false
		}
		isogramMap[string(letter)] = true // assigns the value of letter to isogram map
	}
	return true
}

// ANDREW
// A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
//
// input = "ANDREW"
// isogramMap = ["A": true, "B": false, "C": false, "D": true, "E": true]
//
