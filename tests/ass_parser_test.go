package parsers

import (
	"encoding/json"
	"reflect"
	"testing"

	"main.go/pkg/parsers"
)

func TestParseASS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name: "Basic ASS",
			input: `[Script Info]
Title: Example ASS

[Events]
Dialogue: 0,0:00:01.00,0:00:05.00,Default,,0,0,0,,This is a line of dialogue.
Dialogue: 0,0:00:06.00,0:00:10.00,Default,,0,0,0,,Another line of dialogue.`,
			expected: `[{"StartTime":"0:00:01.00","EndTime":"0:00:05.00","Text":"This is a line of dialogue.","Index":1},{"StartTime":"0:00:06.00","EndTime":"0:00:10.00","Text":"Another line of dialogue.","Index":2}]`,
			wantErr:  false,
		},
		{
			name:     "Empty ASS",
			input:    ``,
			expected: `Error: The content is empty. Please provide valid content.`,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsers.ParseASS(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseASS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				// If an error is expected, we don't compare the output
				return
			}

			// Convert got and expected to JSON objects for comparison
			var gotJSON, expectedJSON interface{}
			if err := json.Unmarshal([]byte(got), &gotJSON); err != nil {
				t.Fatalf("Failed to unmarshal got JSON: %v", err)
			}
			if err := json.Unmarshal([]byte(tt.expected), &expectedJSON); err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}

			// Compare the two JSON objects
			if !reflect.DeepEqual(gotJSON, expectedJSON) {
				t.Errorf("ParseASS() = %v, expected %v", gotJSON, expectedJSON)
			}
		})
	}
}
