package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("./plugin.so")
	if err != nil {
		panic(err)
	}
	sym, err := p.Lookup("TheInt")
	if err != nil {
		panic(err)
	}

	i, ok := sym.(*int)
	if !ok {
		panic("Wrong type")
	}

	fmt.Printf("The int is: %d\n", *i)
}
