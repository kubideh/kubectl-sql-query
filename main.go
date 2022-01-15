package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		usageString := `kubectl-sql-query is the kubectl plugin to query the Kubernetes API server using SQL.

Usage:
  kubectl sql query <query-string>
`
		fmt.Fprintf(flag.CommandLine.Output(), usageString)
	}
	
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(flag.Args())

	fmt.Println("Hello")
}
