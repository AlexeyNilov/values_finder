package gemini

import (
	"strings"
	"testing"
)

func TestGenText(t *testing.T) {
	c := &Client{Model: "gemini-2.0-flash-lite"}
	got := c.GenText("Answer Yes if you read this text")
	if strings.TrimSpace(got) != "Yes" {
		t.Errorf("Expected Yes, got %s", got)
	}
}

// func TestGenerateOptions(t *testing.T) {
// 	c := &Client{Model: "gemini-2.0-flash-lite"}
// 	got, _ := c.GenerateOptions([]core.Choice{})
// 	if got[0] != "" {
// 		t.Errorf("%s", got[0])
// 	}
// }
