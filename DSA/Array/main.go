package main

import "fmt"

func addValue(myArray []int, value ...int) []int {
	return append(myArray, value...)
}
func main() {
	myArray := []int{}
	myArray = addValue(myArray, 112, 113, 334)
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("value : %d ", myArray[i])
	}
}
