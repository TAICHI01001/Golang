package main

import (
	"fmt"

	"github.com/TaiChi112/Golang/mikelopster/other"
	"github.com/google/uuid"
)

func Greeting() {
	fmt.Println("Hi")
}
func main() {
	id := uuid.New()
	fmt.Printf("UUID : %s", id)
	fmt.Println()
	fmt.Println("Hello world!")
	Greeting()
	Greeting()
	Greeting()
	other.Greeting()
	other.Re_Test()
}
