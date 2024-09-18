package parsers

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"main.go/internal/models"
)

// TODO: DOCS: Comment up
// Define regex patterns as global variables
// ParseWithRegex parses the input content using the provided regex pattern
func ParseRegex(content string, re *regexp.Regexp) ([]byte, error) {
	var subtitles []models.Caption
	matches := re.FindAllStringSubmatch(content, -1)

	for i, match := range matches {
		index := uint16(i + 1)
		endTime := match[3]
		startTime := match[2]
		text := strings.TrimSpace(match[4])

		subtitles = append(subtitles, models.Caption{
			StartTime: startTime,
			EndTime:   endTime,
			Index:     index,
			Text:      text,
		})
	}

	// Marshal the data to JSON
	jsonData, err := json.MarshalIndent(subtitles, "", " ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// HelperEmptyString checks if the provided string is empty and returns an error if it is.
func ValidateNonEmpty(s string) error {
	if len(s) > 0 {
		return nil
	}
	return errors.New("Error: The content is empty. Please provide valid content.")
}

// HelperRegexMatch checks if the provided string matches the given regex pattern. Only checks one pattern so
// data might have to be prepared before using.
func ValidateRegexPatterns(regex *regexp.Regexp, s string) error {
	if regex.MatchString(s) {
		return nil
	}
	return errors.New("Error: The input string does not match the required format")
}