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
		fmt.Println(" Display : ", myArray[i])
	}
}
func main() {
	fmt.Println("Now Value in array is null and Add value in array")
	myArray := []int{}
	myArray = addValue(myArray, 112, 113, 114)
	displayValue(myArray)
	fmt.Println("Updated value")
	myArray = updateValue(myArray, 113, 118)
	displayValue(myArray)
	fmt.Println("Delete value")
	myArray = deleteValue(myArray, 112)
	displayValue(myArray)
}
