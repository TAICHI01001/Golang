package main

import (
	"fmt"
	"os"
)

func main() {
	// Declare variable of type os.File and Create file test.txt
	file, err := os.Create("test.txt")
	//handle checking error
	if err != nil {
		//display if having error
		fmt.Println("Error creating file : ", err)
	}
	// closing file after use
	defer file.Close()

	// write to file
	_, err = file.WriteString("Testing I can write to this file")
	//handle checking error
	if err != nil {
		fmt.Println("Error writing to file : ", err)
		return
	}

	//define permissions
	err = os.Chmod("test.txt", 0644)
	//handle checking error
	if err != nil {
		fmt.Println("Error changing file permissions :", err)
		return
	}

	// When everything is ok
	fmt.Println("File created and permissions successfully")
}
