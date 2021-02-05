package scanner

import "testing"

func TestTokenReturn(t *testing.T) {
	got := ScanIntoTokens(`
		// This is scanned code.
		if a > 10 {
			someParsable = text
		}`)
	want := `"Ident": if
	"Ident": a
	"\">\"": >
	"Int": 10
	"\"{\"": {
	"Ident": someParsable
	"\"=\"": =
	"Ident": text
	"\"}\"": }
	`

	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}

}
