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
func deleteValue(myArray []int, value int) []int {
	result := []int{}
	for i := 0; i < len(myArray); i++ {
		if myArray[i] != value {
			result = append(result, myArray[i])
		}
	}
	return result
}
func displayValue(myArray []int) {
	for i := 0; i < len(myArray); i++ {
		fmt.Println(" Value : ", myArray[i])
	}
}
func main() {
	myArray := []int{}
	myArray = addValue(myArray, 112, 113, 114)
	displayValue(myArray)
	updateValue(myArray, 113, 118)
	fmt.Println("......")
	myArray = deleteValue(myArray, 112)
	displayValue(myArray)
}
