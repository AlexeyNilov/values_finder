package gemini

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	c := &Client{Model: "gemini-2.0-flash-lite"}
	got := c.GenText("Answer Yes if you read this text")
	if strings.TrimSpace(got) != "Yes" {
		t.Errorf("Expected Yes, got %s", got)
	}
}
