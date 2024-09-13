package models

// Caption represents a single subtitle entry within a subtitle file.
type Caption struct {
	StartTime string
	EndTime   string
	Text      string
	index     uint16
}
