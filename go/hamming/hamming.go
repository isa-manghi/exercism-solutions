package hamming

import (
	"errors"
)

// Distance is a function that returns the integer hamming distance between two strings (or an error).
func Distance(a, b string) (int, error) {
	var hammingLength int
	if len(a) == len(b) {
		for x := 0; x < len(a); x++ {
			if a[x] != b[x] {
				hammingLength++
			}
		}
		return hammingLength, nil
	}
	return 0, errors.New("input strings are not equal lengths")
}

