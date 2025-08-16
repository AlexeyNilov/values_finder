package gemini

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	got := GenText("I'm testing Gemini API call, does it work? Answer Yes if you read this text")
	if strings.TrimSpace(got) != "Yes" {
		t.Errorf("Expected Yes, got %s", got)
	}
}
