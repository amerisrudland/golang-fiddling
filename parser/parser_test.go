package parser

import "testing"

func TestTokenReturn(t *testing.T) {
	got := Parse(`
		// This is scanned code.
		if a > 10 {
			someParsable = text
		}`)
	want := "hi"

	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}

}
