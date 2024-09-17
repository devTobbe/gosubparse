package parsers

import (
	"encoding/json"
	"strings"

	"main.go/internal/models"
	"main.go/internal/util"
)

// TODO: Comment up parseSRT function
func ParseSRT(content string) (string, error) {
	if err := util.HelperEmptyString(content); err != nil {
		return "", err
	}

	regex := util.SRT
	var subtitles []models.Caption

	// NOTE: Regex checking limits parsing potential, for now. Look into this later.
	// blocks := strings.Split(content, "\n\n")

	//for _, block := range blocks {
	//	if err := util.HelperRegexMatch(regex, block); err != nil {
	//		return "", err
	//	}
	//}

	matches := regex.FindAllStringSubmatch(content, -1)

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

	// Marshal that data
	jsonData, err := json.MarshalIndent(subtitles, "", " ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
