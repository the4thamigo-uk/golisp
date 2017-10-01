package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

func parse(s string) (string, error) {
	return s, nil
}

func preprocess(s string) (string, error) {
	return s, nil
}

func readForm(text string) (sexpr, error) {
	s := &scanner.Scanner{}
	s.Init(strings.NewReader(text))
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch != scanner.EOF && !unicode.IsSpace(ch) && ch != '(' && ch != ')' && (i > 0 || !unicode.IsNumber(ch))
	}
	return scanForm(s)
}

func scanForm(s *scanner.Scanner) (sexpr, error) {
	var se sexpr

	if s.TokenText() != "(" {
		if c := s.Scan(); c != '(' {
			return nil, fmt.Errorf("open brace expected 2")
		}
	}

	for c := s.Scan(); c != scanner.EOF; c = s.Scan() {
		switch c {
		case '(':
			sub, err := scanForm(s)
			if err != nil {
				return nil, err
			}
			se = append(se, sub)
		case ')':
			return se, nil
		case scanner.Int:
			i, err := strconv.ParseInt(s.TokenText(), 0, 0)
			if err != nil {
				return nil, fmt.Errorf("failed to parse int : %v", s.TokenText())
			}
			se = append(se, &intexpr{int(i)})
		case scanner.String:
			se = append(se, &strexpr{s.TokenText()})
		case scanner.Ident:
			se = append(se, &symexpr{s.TokenText()})
		}
	}
	return nil, fmt.Errorf("close brace missing")
}
