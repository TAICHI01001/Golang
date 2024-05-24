package main

import "fmt"

func main() {
	myArray1 := []int{1, 2, 3}
	for i := 0; i < len(myArray1); i++ {
		fmt.Print("", myArray1[i])
	}
	var myArray2 = []int{}
	fmt.Println("", len(myArray2))
}
