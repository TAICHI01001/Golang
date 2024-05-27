package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	name := flag.String("var", "What's I can help you!", "This first test golang cli")
	flag.Parse()
	if flag.NArg() > 0 {
		*name = strings.Join(flag.Args(), " ")
	}
	fmt.Printf("%s\n", *name)
}
