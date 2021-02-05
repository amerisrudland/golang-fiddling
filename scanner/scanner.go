package scanner

import (
	"fmt"
	"strings"
	"text/scanner"
)


const (
	EOF = -(iota + 1)
	Equals
	GreaterThan
	LessThan
	OpenBracket
	CloseBracket
)

func ScanIntoTokens(input string) string {
	var s scanner.Scanner
	out := ""
	s.Init(strings.NewReader(input))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		out += fmt.Sprintf("%q: %s\n", scanner.TokenString(tok), s.TokenText())
	}
	return out
}