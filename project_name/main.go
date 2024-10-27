package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println("i=", i, "j=", j, "k=", k, "c=", c, "python=", python, "java=", java)

	fmt.Println("My favorite number is", rand.Intn(10))
}
