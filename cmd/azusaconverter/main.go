package main

import (
	"fmt"
	"os"

	"github.com/alitaso345/azusaconverter"
)

const name = "azusaconverter"

func main() {
	os.Exit(run(os.Args[1]))
}

const (
	exitCodeOK = iota
	exitCodeErr
)

func run(input string) int {
	if len(input) == 0 {
		fmt.Fprintf(os.Stderr, "usage: #{name} text \n")
		return exitCodeErr
	}
	convert(input)
	return exitCodeOK
}

func convert(input string) {
	output := azusaconverter.Convert(input)
	fmt.Fprint(os.Stdout, output)
}
