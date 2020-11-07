package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sensimevanidus/comet"
)

const usage = `Usage: comet [options...] <yamlFile>

Options:
  -v  Verbose output.
`

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	if 1 > len(flag.Args()) {
		printUsageAndExit()
	}

	if err := comet.RunTestSuite(flag.Args()[0], *verbose); err != nil {
		log.Fatal(err)
	}
}
