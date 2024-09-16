package util

import (
	"regexp"
)

// TODO: Add more formats later on, like .smi (SAMI)

// Precompiled regular expressions for different subtitle formats
var (
	SRT = regexp.MustCompile(`(?m)^(\d+)\n(\d{2}:\d{2}:\d{2},\d{3})\s+-->\s+(\d{2}:\d{2}:\d{2},\d{3})\n((?:[^\n]+\n?)*)`)
	ASS = regexp.MustCompile(`(?m)^Dialogue:\s*(?:|\w+\W)(?:Marked=\d+,\s*)?\d+,\s*(\d{2}:\d{2}:\d{2}\.\d{2}),\s*(\d{2}:\d{2}:\d{2}\.\d{2}),[^,]*,[^,]*,\d+,\d+,\d+,[^,]*,\s*(.+)$`)
	VTT = regexp.MustCompile(`(?m)(?:^(\d+)\n)?(\d{2}:\d{2}:\d{2}\.\d{3})\s*-->\s*(\d{2}:\d{2}:\d{2}\.\d{3})\n((?:[^\n]+\n?)*)\n?(?:\n\n|\n\d+|\s*$)`)
)
