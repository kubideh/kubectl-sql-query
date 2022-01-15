package main

import (
	"flag"
	"fmt"
)

func Run() {
	if len(flag.Args()) == 0 {
		flag.Usage()
	}

	fmt.Println(flag.Args())

	fmt.Println("Hello")
}
