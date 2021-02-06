package scanner

import "testing"

func TestTokenReturn(t *testing.T) {
	got := ScanIntoTokens(`
		// This is scanned code.
		if a > 10 {
			someParsable = text
		}`)
	want := []string{
		"Ident",
		"Ident",
		"\">\"",
		"Int",
		"\"{\"",
		"Ident",
		"\"=\"",
		"Ident",
		"\"}\""}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("expected %v but got %v", want[i], got[i])
		}
	}

}
