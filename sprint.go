package main

import (
	"regexp"
)

// NameWithoutPoints returns a clean name, without story points
func NameWithoutPoints(text string) string {
	r, err := regexp.Compile(`\(([0-9]*\.?[0-9]+)\)`)
	handleError(err)

	processed := r.ReplaceAllString(text, "")

	s := []rune(processed)

	final := processed

	if string(s[0]) == " " {
		final = string(s[1:])
	} else if string(s[len(processed)-1]) == " " {
		final = string(s[0 : len(processed)-1])
	}

	return final
}
