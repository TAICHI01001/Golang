package main

import (
	"fmt"

	help "github.com/TAICHI01001/Golang-basic/func"
	math1 "github.com/TAICHI01001/Golang-basic/math"

	"rsc.io/quote"
)

func add(n int) int {
	return n + n
}

// const age int = 19

var a, b, c, d int = 1, 2, 3, 4
var (
	x, y, z, v string = "a", "b", "c", "d"
)

func main() {
	//! test absolute error numerical method
	fmt.Println("absolute error :", math1.ErrorAbsolute(22/21, 0.10476*(10^1)))

	// fmt.Println("error relative : ", math1.ErrorRelative(func() float64 { return math1.ErrorAbsolute(5.0, 6.0) }, 8.0))
	result := math1.ErrorAbsolute(1.04761904762, 0.10476*(10^1))
	fmt.Println("error relative : ", math1.ErrorRelative(func() float64 { return result }, 1.04761904762))

	fmt.Println(x, y, z, v)
	var e, f = 8, "Hi"
	g, h := "hey", 112
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)
	var (
		x int
		b int    = 2
		z string = "Hello"
	)
	fmt.Println(x, b, z)
	var person1 string = "TaiChi"
	fmt.Println("Name : ", person1)
	fmt.Println("addition : ", add(5))
	help.Message()
	help.Server()
	fmt.Println(quote.Go())
}
