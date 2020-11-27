package scrabble

import "strings"

func Scorev1(inputWord string) int {
	inputWord = strings.ToUpper(inputWord)
	var scrabbleScoresMap = map[int][]string{
		1:  []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  []string{"D", "G"},
		3:  []string{"B", "C", "M", "P"},
		4:  []string{"F", "H", "V", "W", "Y"},
		5:  []string{"K"},
		8:  []string{"J", "X"},
		10: []string{"Q", "Z"},
	}
	// the for loop needs to loop through each letter of the string scrabble word
	// look for the letter in the table,
	// for that letter look for the corresponding value.
	// jot down the value. repeat until the word is complete.
	// add all values up.
	var totalScrabbleScore int
	for x := len(inputWord) - 1; x >= 0; x-- {
		for scrabbleScore, listOfLettersWithThisScrabbleScore := range scrabbleScoresMap {
			for _, actualLetter := range listOfLettersWithThisScrabbleScore {
				if string(inputWord[x]) == actualLetter {
					totalScrabbleScore += scrabbleScore
				}
			}
		}
	}
	return totalScrabbleScore
}
