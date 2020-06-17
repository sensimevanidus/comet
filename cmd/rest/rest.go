package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sensimevanidus/rest"
)

const usage = `Usage: rest <yaml-file>
`

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if 1 >= len(os.Args) {
		printUsageAndExit()
	}

	if err := rest.RunTestSuite(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
