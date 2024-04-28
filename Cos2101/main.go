package main

import "fmt"

type Person struct {
	name string
	age  int
}

func DefaultPerson1() Person {
	return Person{
		age: 10,
	}
}
func (p Person) PrintInfo() {
	fmt.Printf("Name : %s, age : %d\n", p.name, p.age)
}
func main() {
	p := Person{name: "Tom", age: 10}

	fmt.Println("Define name :", p.name)
	p.age = 11
	fmt.Println("Define age :", p.age)
	p.PrintInfo()

	var person Person
	person.age = 112
	fmt.Println("first try struct type in GO :	", person.age)
	person.age = 10
	fmt.Println("first try struct type in GO :	", person.age)
	fmt.Println("Hello Test Test!")
}
