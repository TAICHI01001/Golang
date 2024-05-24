package main

import "fmt"

func addValue(myArray []int, value ...int) []int {
	return append(myArray, value...)
}
func updateValue(myArray []int, value int, newValue int) []int {
	for i := 0; i < len(myArray); i++ {
		if myArray[i] == value {
			myArray[i] = newValue
		}
	}
	return myArray
}
func displayValue(myArray []int) {
	for i := 0; i < len(myArray); i++ {
		fmt.Println(" Value : ", myArray[i])
	}
}
func main() {
	myArray := []int{}
	myArray = addValue(myArray, 112, 113, 334)
	displayValue(myArray)
	updateValue(myArray, 113, 118)
	displayValue(myArray)
}
