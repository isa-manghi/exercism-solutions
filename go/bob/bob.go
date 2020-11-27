// Package bob should have a package comment that summarizes what it's about.
package bob

import "strings"

// Hey returns a string dependant on what the remark is.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var isInputRemarkAllCaps = remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
	var isInputRemarkAQuestion = strings.HasSuffix(remark, "?")
	if isInputRemarkAllCaps && isInputRemarkAQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isInputRemarkAllCaps {
		return "Whoa, chill out!"
	}
	if isInputRemarkAQuestion {
		return "Sure."
	}
	if remark == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
