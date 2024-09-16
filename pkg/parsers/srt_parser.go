package parsers

import (
	"encoding/json"
	"strconv"
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

	valdationContent := strings.Split(content, "\n\n")

	for _, vals := range valdationContent {
		if err := util.HelperRegexMatch(regex, vals); err != nil {
			return "", err
		}
	}

	matches := regex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		index := match[1]
		endTime := match[3]
		startTime := match[2]
		text := strings.TrimSpace(match[4])

		intIndex, err := strconv.Atoi(index)
		if err != nil {
			return "", err
		}

		subtitles = append(subtitles, models.Caption{
			StartTime: startTime,
			EndTime:   endTime,
			Text:      text,
			Index:     uint16(intIndex),
		})
	}

	// Marshal that data
	jsonData, err := json.MarshalIndent(subtitles, "", " ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
