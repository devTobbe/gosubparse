package util

import (
	"errors"
	"regexp"
)

// HelperEmptyString checks if the provided string is empty and returns an error if it is.
func HelperEmptyString(s string) error {
	if len(s) > 0 {
		return nil
	}
	return errors.New("Error: The content is empty. Please provide valid content.")
}

// HelperRegexMatch checks if the provided string matches the given regex pattern. Only checks one pattern so
// data might have to be prepared before using.
func HelperRegexMatch(regex *regexp.Regexp, s string) error {
	if regex.MatchString(s) {
		return nil
	}
	return errors.New("Error: The input string does not match the required format")
}
