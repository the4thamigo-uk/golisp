package main

import (
	"fmt"
)

type val interface {
	print() error
}

type intval struct {
	val int
}

func (i *intval) print() error {
	if i == nil {
		return fmt.Errorf("intval is nil")
	}
	fmt.Printf("%v", i.val)
	return nil
}

type strval struct {
	val string
}

func (s *strval) print() error {
	if s == nil {
		return fmt.Errorf("strval is nil")
	}
	fmt.Printf("'%v'", s.val)
	return nil
}

type funcval struct {
	val func(args []val) (val, error)
}

func (f *funcval) print() error {
	if f == nil {
		return fmt.Errorf("funcval is nil")
	}
	fmt.Printf("%v", f.val)
	return nil
}

type listval []val

func (l *listval) print() error {
	if l == nil {
		return fmt.Errorf("listval is nil")
	}
	fmt.Print("(")
	for i, v := range *l {
		fmt.Printf("%v", v)
		if i < len(*l)-1 {
			fmt.Print(",")
		}
	}
	fmt.Print(")")
	return nil
}
