package main

import (
	"fmt"
)

type expr interface {
	eval(e *env) (val, error)
}

type symexpr struct {
	expr string
}

func (s *symexpr) eval(e *env) (val, error) {
	if s == nil {
		return nil, fmt.Errorf("symexpr is nil")
	}
	return e.lookup(s)
}

type intexpr struct {
	expr int
}

func (i *intexpr) eval(e *env) (val, error) {
	if i == nil {
		return nil, fmt.Errorf("intexpr is nil")
	}
	return &intval{i.expr}, nil
}

type strexpr struct {
	expr string
}

func (s *strexpr) eval(e *env) (val, error) {
	if s == nil {
		return nil, fmt.Errorf("strexpr is nil")
	}
	return &strval{s.expr}, nil
}

type sexpr []expr

func (se sexpr) eval(e *env) (val, error) {

	if len(se) > 0 {
		if s, ok := (se[0]).(*symexpr); ok && s.expr == "lambda" {
			if len(se[1:]) != 2 {
				return nil, fmt.Errorf("lambda syntax error")
			}
			params, ok := interface{}(se[1]).(sexpr)
			if !ok {
				return nil, fmt.Errorf("lambda params not a sexpr")
			}
			body, ok := interface{}(se[2]).(sexpr)
			if !ok {
				return nil, fmt.Errorf("lambda body not a sexpr")
			}
			return lambda(params, body, e)
		}
	}

	var vals listval

	for _, expr := range se {
		val, err := expr.eval(e)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}

	if f, ok := (vals[0]).(*funcval); ok {
		// apply
		return f.val(vals[1:])
	}
	return &vals, nil
}
