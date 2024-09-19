package parsers

import "main.go/internal/util"

func ParseSSA(content string) ([]byte, error) {
	// Check input data
	if err := ValidateNonEmpty(content); err != nil {
		return nil, err
	}

	NoBracket := util.BracketPattern.ReplaceAllString(content, "")

	return ParseRegex(NoBracket, util.ASS)
}
