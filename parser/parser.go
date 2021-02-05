package parser

import (
	"fmt"
	"strings"
	"text/scanner"
)

type parserToken int

const (
	pIdentifier parserToken = 0
)

func Parse(input string) string {
	var s scanner.Scanner
	out := ""
	s.Init(strings.NewReader(input))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		out += fmt.Sprintf("%q: %s\n", scanner.TokenString(tok), s.TokenText())
	}
	return out
}
