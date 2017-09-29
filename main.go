package main

import (
	"fmt"
)

func main() {
	l := sexpr{
		&symexpr{"+"},
		sexpr{
			&symexpr{"/"},
			&intexpr{10},
			&intexpr{5},
		},
		sexpr{
			&symexpr{"*"},
			&symexpr{"x"},
			&intexpr{2},
			sexpr{
				&symexpr{"+"},
				&intexpr{1},
				&intexpr{1},
				&intexpr{1},
			},
		},
	}

	l = sexpr{
		sexpr{
			&symexpr{"lambda"},
			sexpr{
				&symexpr{"a"},
				&symexpr{"b"},
			},
			sexpr{
				&symexpr{"+"},
				&symexpr{"a"},
				&symexpr{"b"},
			},
		},
		&intexpr{3},
		&intexpr{2},
	}

	e := stdEnv()
	e.add("x", &intval{2})
	v, err := l.eval(e)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = v.print()
	if err != nil {
		fmt.Println(err)
		return
	}
}
