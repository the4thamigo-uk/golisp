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
			&intexpr{2},
			&intexpr{2},
			sexpr{
				&symexpr{"+"},
				&intexpr{1},
				&intexpr{1},
				&intexpr{1},
			},
		},
	}

	e := stdEnv()
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
