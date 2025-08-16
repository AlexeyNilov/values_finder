package util

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

func ParseTemplate(templateStr string, data any) string {
	// Parse the template
	tpl, err := template.New("New").Parse(templateStr)
	if err != nil {
		panic(err)
	}

	// Use a bytes.Buffer to capture the output
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// Read from text file and return content as a string
func ReadTemplate(filepath string) string {
	// os.ReadFile reads the entire file and returns its contents as a byte slice.
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string and return it.
	return string(content)
}
