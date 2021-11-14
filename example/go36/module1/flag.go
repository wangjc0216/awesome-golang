package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	flag.StringVar(&name, "name", "everyone", "The greeting object.")

}

func main() {
	flag.Parse()
	hello(name)
}
func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
