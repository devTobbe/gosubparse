package util

import (
	"regexp"
)

// TODO: Add more formats later on, like .smi (SAMI)

// Precompiled regular expressions for different subtitle formats
var (
	SRT = regexp.MustCompile(`(?m)^(?P<Line>\d+)\n(?P<StartTime>\d{2}:\d{2}:\d{2},\d{3})\s+-->\s+(?P<EndTime>\d{2}:\d{2}:\d{2},\d{3})\n(?P<Content>(?:[^\n]+\n?)+)(?=\n\n|\n\d+\n|$)`)
	ASS = regexp.MustCompile(`(?m)^Dialogue:\s*(?:|\w+\W)(?:Marked=\d+,\s*)?\d+,(?P<StartTime>\d{2}:\d{2}:\d{2}\.\d{2}),(?P<EndTime>\d{2}:\d{2}:\d{2}\.\d{2}),[^,]*,[^,]*,\d+,\d+,\d+,[^,]*,(?P<Content>.+)$`)
	VTT = regexp.MustCompile(`(?m)(?:^(?P<Line>\d+)\n)?(?P<StartTime>\d{2}:\d{2}:\d{2}\.\d{3})\s*-->\s*(?P<EndTime>\d{2}:\d{2}:\d{2}\.\d{3})\n(?P<Content>(?:[^\n]+\n?)*)\n?(?=\n\n|\n\d+|\s*$)`)
)
