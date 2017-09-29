package main

import (
	"fmt"
)

type env struct {
	prior *env
	m     map[string]val
}

func stdEnv() *env {
	return &env{
		m: map[string]val{
			"+": intPlus,
			"*": intMultiply,
			"-": intSubtract,
			"/": intDivide,
		},
	}
}

func (e *env) add(s string, val val) error {
	if e != nil && e.m != nil {
		e.m[s] = val
		return nil
	}

	return fmt.Errorf("Invalid environment")
}

func (e *env) lookup(s *symexpr) (val, error) {
	if e != nil && e.m != nil {
		if val, ok := e.m[s.expr]; ok {
			if val == nil {
				return nil, fmt.Errorf("Value found in environment is nil : %v : %v", s.expr, e)
			}
			return val, nil
		}
	}

	if e.prior != nil {
		return e.prior.lookup(s)
	}

	return nil, fmt.Errorf("Value not found in environment : %v : %v", s.expr, e)
}

func (e *env) push() *env {
	return &env{
		prior: e,
		m:     map[string]val{},
	}
}
