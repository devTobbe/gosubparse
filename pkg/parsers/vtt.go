package parsers

import (
	"html"
	"strings"

	"main.go/internal/util"
)

// TODO: DOCS: Comment up parseVTT function
func ParseVTT(content string) ([]byte, error) {
	// VTTRegex should be defined similarly
	// Check input data
	if err := ValidateNonEmpty(content); err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")

	// If there are less than 3 lines, return an empty string
	if len(lines) <= 2 {
		return ParseSRT("")
	}

	noTags := util.TagPatternVTT.ReplaceAllString(strings.Join(lines[2:], "\n"), "")

	// Revert coded HTML items
	noEntities := html.UnescapeString(noTags)

	// Join the lines back together, skipping the first two
	s := noEntities

	VTTRegex := util.SRT
	return ParseRegex(s, VTTRegex)
}

//func ParseVTT(content string) (string, error) {
//	lines := strings.Split(content, "\n")
//
//	// If there are less than 3 lines, return an empty string
//	if len(lines) <= 2 {
//		return ParseSRT("")
//	}
//
//	noTags := util.TagPatternVTT.ReplaceAllString(strings.Join(lines[2:], "\n"), "")
//
//  //Revert coded HTML items
//	noEntities := html.UnescapeString(noTags)
//
//	// Join the lines back together, skipping the first two
//	s := noEntities
//	return ParseSRT(s)
//}
