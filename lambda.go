package main

import (
	"fmt"
)

func lambda(args []expr, e *env) (val, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("lambda needs two arguments")
	}
	params, ok := interface{}(args[0]).(sexpr)
	if !ok {
		return nil, fmt.Errorf("lambda params not a sexpr")
	}
	body, ok := interface{}(args[1]).(sexpr)
	if !ok {
		return nil, fmt.Errorf("lambda body not a sexpr")
	}
	return &funcval{
			func(vals []val) (val, error) {
				if len(params) != len(vals) {
					return nil, fmt.Errorf("invalid number of arguments")
				}

				e2 := e.push()

				for i, p := range params {
					if s, ok := p.(*symexpr); ok {
						e2.add(s, vals[i])
					} else {
						return nil, fmt.Errorf("parameter is not a symbol : %v", p)
					}
				}
				return body.eval(e2)
			},
		},
		nil
}

func defVar(args []expr, e *env) (val, error) {

	if len(args) != 2 {
		return nil, fmt.Errorf("defvar need two arguments")
	}
	name, ok := interface{}(args[0]).(*symexpr)
	if !ok {
		return nil, fmt.Errorf("defvar name is not a symexpr")
	}
	body := args[1]

	val, err := body.eval(e)
	if err != nil {
		return nil, err
	}

	e.add(name, val)
	return val, nil
}
