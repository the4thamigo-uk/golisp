package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	e := stdEnv()
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\n\nlisp>  ")
		text, _ := reader.ReadString('\n')

		l, err := readForm(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		v, err := l.eval(e)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if v == nil {
			fmt.Println("nil")
		} else {
			err = v.print()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
