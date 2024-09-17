package parsers

import (
	"encoding/json"
	"reflect"
	"testing"

	"main.go/pkg/parsers"
)

func TestParseSRT(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name: "Basic SRT",
			input: `1
00:00:00,000 --> 00:00:02,000
Subtitle text here

2
00:00:03,000 --> 00:00:05,000
Another subtitle`,
			expected: `[{"StartTime":"00:00:00,000","EndTime":"00:00:02,000","Text":"Subtitle text here","Index":1},{"StartTime":"00:00:03,000","EndTime":"00:00:05,000","Text":"Another subtitle","Index":2}]`,
			wantErr:  false,
		},
		{
			name:     "Empty SRT",
			input:    ``,
			expected: `Error: The content is empty. Please provide valid content.`,
			wantErr:  true,
		},
		//		{
		//			name: "Invalid SRT",
		//			input: `1
		//00:00:00,000 --> 00:00:02,000
		//Subtitle text here
		//
		//2
		//00:00:03,000 --> 00:00:05,000`, // Invalid because it lacks text for the second subtitle block
		//			expected: `Error: The input string does not match the required format`, // Expect empty JSON array if format is invalid
		//			wantErr:  true,
		//		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsers.ParseSRT(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSRT() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("ParseSRT() = %v, expected %v", gotJSON, expectedJSON)
			}
		})
	}
}
