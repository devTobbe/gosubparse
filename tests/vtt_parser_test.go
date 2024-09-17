package parsers

import (
	"encoding/json"
	"reflect"
	"testing"

	"main.go/pkg/parsers"
)

func TestParseVTT(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name: "Basic VTT",
			input: `WEBVTT

1
00:00:00.000 --> 00:00:02.000
Subtitle text here

2
00:00:03.000 --> 00:00:05.000
Another subtitle`,
			expected: `[{"StartTime":"00:00:00.000","EndTime":"00:00:02.000","Text":"Subtitle text here","Index":1},{"StartTime":"00:00:03.000","EndTime":"00:00:05.000","Text":"Another subtitle","Index":2}]`,
			wantErr:  false,
		},
		{
			name:     "Empty VTT",
			input:    ``,
			expected: `Error: The content is empty. Please provide valid content.`,
			wantErr:  true,
		},
		//		{
		//			name: "Invalid VTT",
		//			input: `WEBVTT
		//
		//1
		//00:00:00.000 --> 00:00:02.000
		//Subtitle text here
		//
		//2
		//00:00:03.000 --> 00:00:05.000`, // Invalid because it lacks text for the second subtitle block
		//			expected: `Error: The input string does not match the required format`, // Expect an error since the format is invalid
		//			wantErr:  true,
		//		},

		{
			name:  "Complex WebVTT",
			input: "WEBVTT\n\nSTYLE\n::cue {\n    background: rgba(0, 0, 0, 0.56);\n    text-shadow: 0 0 7px #000;\n    color: #fff;\n    font-family: Arial, sans-serif;\n    font-weight: bold;\n}\n::cue(v[voice=\"Speaker1\"]) { color: #00FF00; }\n::cue(v[voice=\"Speaker2\"]) { color: #0000FF; }\n\n1\n00:00:01.785 --> 00:00:03.804 align:center line:90% position:50%\n<v Speaker1>Hello, I am Speaker1</v>\n\n2\n00:00:04.123 --> 00:00:10.602 align:center line:90% position:50%\n<v Speaker2>and I am the almighty Speaker2!</v>\n\n3\n00:00:11.102 --> 00:00:15.105 align:center line:10% position:50%\n<v Speaker1>and again this is Speaker1 speaking</v>\n\n4\n00:00:16.000 --> 00:00:18.000 align:left line:50% position:10%\n<v Speaker2>This is a subtitle with different alignment and positioning.</v>\n\n5\n00:00:19.000 --> 00:00:22.000 align:right line:20% position:90%\n<v Speaker1>Here is another subtitle with right alignment.</v>\n\n6\n00:00:23.000 --> 00:00:25.000\nThis subtitle has no voice tag and should use default styling.\n\n7\n00:00:26.000 --> 00:00:30.000 align:center line:80% position:50%\n<v Speaker1>Special characters: <i>italic</i>, <b>bold</b>, &amp;.</v>",
			expected: `[
        {"StartTime": "00:00:01.785", "EndTime": "00:00:03.804", "Text": "Hello, I am Speaker1", "Index": 1},
        {"StartTime": "00:00:04.123", "EndTime": "00:00:10.602", "Text": "and I am the almighty Speaker2!", "Index": 2},
        {"StartTime": "00:00:11.102", "EndTime": "00:00:15.105", "Text": "and again this is Speaker1 speaking", "Index": 3},
        {"StartTime": "00:00:16.000", "EndTime": "00:00:18.000", "Text": "This is a subtitle with different alignment and positioning.", "Index": 4},
        {"StartTime": "00:00:19.000", "EndTime": "00:00:22.000", "Text": "Here is another subtitle with right alignment.", "Index": 5},
        {"StartTime": "00:00:23.000", "EndTime": "00:00:25.000", "Text": "This subtitle has no voice tag and should use default styling.", "Index": 6},
        {"StartTime": "00:00:26.000", "EndTime": "00:00:30.000", "Text": "Special characters: italic, bold, &.", "Index": 7}
    ]`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsers.ParseVTT(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseVTT() error = %v, wantErr %v", err, tt.wantErr)
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
				t.Errorf("ParseVTT() = %v, expected %v", gotJSON, expectedJSON)
			}
		})
	}
}
