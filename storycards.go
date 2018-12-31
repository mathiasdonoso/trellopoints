package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	// ErrDoesntHavePoints is returned when a text has no story points
	ErrDoesntHavePoints = errors.New("it doesn't have story points")
)

// HasStoryPoint checks if a string has a story point assigned
func HasStoryPoint(text string) bool {
	// Regex for floatingpoint from https://www.regular-expressions.info/floatingpoint.html
	matched, err := regexp.MatchString(`\(([0-9]*\.?[0-9]+)\)(\w)*|(\w)*\(([0-9]*\.?[0-9]+)\)`, text)
	handleError(err)

	return matched
}

// GetStoryPoint returns the story points assigned
func GetStoryPoint(text string) (float64, error) {
	var points float64
	if !HasStoryPoint(text) {
		return points, ErrDoesntHavePoints
	}

	r, err := regexp.Compile(`\(([0-9]*\.?[0-9]+)\)`)
	if err != nil {
		return points, err
	}

	stringPoints := stripchars(r.FindString(text), "()")

	points, err = strconv.ParseFloat(stringPoints, 64)
	if err != nil {
		return float64(0), err
	}

	return points, nil
}

// @source https://www.rosettacode.org/wiki/Strip_a_set_of_characters_from_a_string#Go
func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
