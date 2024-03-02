package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array: ")
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	bubbleSort(arr)
	fmt.Println("\nSorted array: ")
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
