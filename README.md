# 🎬 GoSubParse (⚠️ WIP ⚠️)

A Go library for parsing subtitle files in `.srt`, `.ass`, `.ssa`, `.vtt` formats into a JSON format. This project is a Go rewrite of the TypeScript library [**SubParse**](https://github.com/devTobbe/subparse), offering a simple and efficient way to work with subtitle data programmatically.

## 🚀 Features

- **Supports Multiple Subtitle Formats**: Convert subtitle files in `.srt`, `.ass`, `.ssa` and `.vtt` formats into JSON.
- **Automatic File Detection**: Automatically identifies and parses subtitle files based on their extension.
- **Customizable Output**: Tailor the JSON output to your needs with various format options.

## 📦 Installation

To install **GoSubParse**, you need to use Go modules. First, initialize your module, then get the library:

```bash
go mod init your-module-name
go get github.com/your-username/subparse-go
```

## 📜 Usage Instructions

To use **GoSubParse** in your project, follow these steps:

1. **Import the Library:**
   <!--  TODO: Proper import once the package is finalized -->

   ```go
   import "github.com/your-username/subparse-go/pkg/parser"
   ```

   <!-- TODO: Change usage instructions once finished. -->

2. **Call `ParseSubtitleFile`:**

   Use the `ParseSubtitleFile` function to parse your subtitle files. Provide the path to the file you want to parse.

   The function automatically detects the subtitle format (e.g., SRT, ASS) based on the file content.

   Example usage:

   ```go
   package main

   import (
       "fmt"
       "log"
       "github.com/your-username/subparse-go/pkg/parser"
   )

   func main() {
       filePath := "path/to/your/subtitle/file.srt"
       subtitleData, err := parser.ParseSubtitleFile(filePath)
       if err != nil {
           log.Fatalf("Error parsing file: %v", err)
       }

       fmt.Printf("Parsed data: %+v\n", subtitleData)
   }
   ```

   Example output for an SRT file with the following content:

   ```srt
   1
   00:00:01,000 --> 00:00:04,000
   Hello, world!
   ```

   Would be:

   ```json
   [
     {
       "line": 1,
       "start": "00:00:01,000",
       "end": "00:00:04,000",
       "text": "Hello, world!"
     }
   ]
   ```

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
