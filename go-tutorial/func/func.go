package functions

import "fmt"

func greeting() {
	fmt.Println("Hello Golang")
}
func add[number int | float64 | string](x, y number) number {
	return x + y
}
