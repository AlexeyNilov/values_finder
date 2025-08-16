package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	// Define a sample template string
	templateStr := "Hello, {{.Name}}! Welcome to {{.Place}}."

	// Define the data to be used in the template
	data := struct {
		Name  string
		Place string
	}{
		Name:  "Alice",
		Place: "Wonderland",
	}

	// Call the ParseTemplate function
	result := ParseTemplate(templateStr, data)

	// Expected output
	expected := "Hello, Alice! Welcome to Wonderland."

	// Assert that the result matches the expected output
	assert.Equal(t, expected, result)
}
