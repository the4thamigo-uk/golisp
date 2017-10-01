package main

import (
	"fmt"
	"reflect"
)

func intOp(op func(int, int) int) *funcval {
	return &funcval{
		func(args []val) (val, error) {
			sum := &intval{}
			for i, arg := range args {
				val, ok := arg.(*intval)
				if !ok {
					return nil, fmt.Errorf("intadd arg is not an intval : %v : %v : %v", i, arg, reflect.TypeOf(arg))
				}
				if val == nil {
					return nil, fmt.Errorf("intadd arg is nil : %v : %v", i, arg)
				}
				if i == 0 {
					sum.val = val.val
				} else {
					sum.val = op(sum.val, val.val)
				}
			}

			return sum, nil
		}}
}

var (
	intPlus = intOp(func(x int, y int) int {
		return x + y
	})
	intMultiply = intOp(func(x int, y int) int {
		return x * y
	})
	intSubtract = intOp(func(x int, y int) int {
		return x - y
	})
	intDivide = intOp(func(x int, y int) int {
		return x / y
	})
)
