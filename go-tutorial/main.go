package main

import "fmt"

type Pair[test1 int, test2 int] struct {
	data1 test1
	data2 test2
}

func add[number int | float64 | string](x, y number) number {
	return x + y
}
func main() {
	data1 := Pair[int, int]{data1: 1, data2: 2}
	fmt.Println("Pair : ", data1)
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
	fmt.Println(add("112 ", " 112"))
}
