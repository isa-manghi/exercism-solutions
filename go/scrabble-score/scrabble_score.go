package scrabble

import "strings"

// Score maps the letters in the inputword and adds the score up for a good game of scrabble.
func Score(inputWord string) int {
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
	var totalScrabbleScore int
	for x := 0; x < len(inputWord); x++ {
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
