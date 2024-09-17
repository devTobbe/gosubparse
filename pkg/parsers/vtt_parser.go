package parsers

import (
	"html"
	"strings"

	"main.go/internal/util"
)

// TODO: Comment up parseVTT function
func ParseVTT(content string) (string, error) {
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
	return ParseSRT(s)
}
