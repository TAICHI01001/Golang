package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Product struct {
	name  string
	price float64
	count int
	info  string
}

func (pd Product) Display() {
	fmt.Printf("Name : %s, price : %f, count : %d, info : %s\n", pd.name, pd.price, pd.count, pd.info)
}
func (p Person) PrintInfo() {
	fmt.Printf("Name : %s, age : %d\n", p.name, p.age)
}
func main() {
	pd := Product{
		name:  "Apple",
		price: 10.5,
		count: 10,
		info:  "good",
	}
	pd.Display()
	var pd2 Product
	pd2.name = "Banana"
	pd2.price = 20.5
	pd2.count = 20
	pd2.info = "bad"
	pd2.Display()

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
