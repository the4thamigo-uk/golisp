package main

import (
	"fmt"
)

type env map[string]val

func stdEnv() env {
	return env{
		"+": intPlus,
		"*": intMultiply,
		"-": intSubtract,
		"/": intDivide,
	}
}

func (e env) lookup(s *symexpr) (val, error) {
	if e != nil {
		if val, ok := e[s.expr]; ok {
			if val == nil {
				return nil, fmt.Errorf("Value found in environment is nil : %v : %v", s.expr, e)
			}
			return val, nil
		}
	}

	return nil, fmt.Errorf("Value not found in environment : %v : %v", s.expr, e)
}
