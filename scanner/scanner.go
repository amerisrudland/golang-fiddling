package scanner

import (
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

func ScanIntoTokens(input string) []string {
	var s scanner.Scanner
	var out []string
	s.Init(strings.NewReader(input))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		out = append(out, scanner.TokenString(tok))
	}
	return out
}
