package isogram

import "strings"

// IsIsogram determines if the input is an isogram.
func IsIsogram(input string) bool {
	input = strings.ToUpper(input)
	isogramMap := make(map[string]bool)
	for _, letter := range input {
		if isogramMap[string(letter)] == true && letter <= 90 && letter >= 65 {
			return false
		}
		isogramMap[string(letter)] = true
	}
	return true
}
