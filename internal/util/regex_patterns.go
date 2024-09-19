package util

import (
	"regexp"
)

// TODO: Add more formats later on, like .smi (SAMI) and .sub(MircoDVD)
// Precompiled regular expressions for different subtitle formats
var (
	SRT            = regexp.MustCompile(`(?m)^\d+?\n(\d{2}:\d{2}:\d{2}[.,]\d{3})\s+-->\s+(\d{2}:\d{2}:\d{2}[.,]\d{3})(?:[^\n])*\n((?:[^\n]+\n?)+)`)
	ASS            = regexp.MustCompile(`(?m)^Dialogue:\s*(?:|\w+\W)(?:Marked=\d+,\s*)?\d+,(?P<StartTime>\d+:\d+:\d+\.\d+),(?P<EndTime>\d+:\d+:\d+\.\d+),[^,]*,[^,]*,\d+,\d+,\d+,[^,]*,(?P<Content>.+)$`)
	VTT            = regexp.MustCompile(`(?m)^(?:WEBVTT\n)?(?:\d+\n)?(\d{2}:\d{2}:\d{2}\.\d{3})\s*-->\s*(\d{2}:\d{2}:\d{2}\.\d{3})\n((?:[^\n]+\n?)*)\n`)
	TagPattern     = regexp.MustCompile(`<[^>]+>`)
	BracketPattern = regexp.MustCompile(`{[^}]+}`)
)
