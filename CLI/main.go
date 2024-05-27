package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("", "What's I can help you", "This first test step golang cli")
	flag.Parse()
	fmt.Printf("%s!\n", *name)
}
