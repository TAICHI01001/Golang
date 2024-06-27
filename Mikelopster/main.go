package main

import (
	"fmt"
)

//	func Greeting() {
//		fmt.Println("Hi")
//	}
//
//	func Hi(name string) {
//		fmt.Printf("My name %s", name)
//	}
// type Student struct {
// 	name   string
// 	weight int
// 	height int
// 	grade  string
// }

// type Person struct {
// 	name    string
// 	age     int
// 	address Address
// }
// type Address struct {
// 	street  string
// 	city    string
// 	zipCode int
// }

//	func sayHello(word string, round int) {
//		for i := 0; i < round; i++ {
//			fmt.Println("Hello : ", word)
//		}
//	}
//
//	func add(a int, b int) int {
//		sumNum := a + b
//		return sumNum
//	}
// type Student struct {
// 	firstName string
// 	lastName  string
// }

// func (student Student) fullName() string {
// 	return student.firstName + " " + student.lastName
// }

// type Speaker interface {
// 	Speak() string
// }
// type Dog struct {
// 	name string
// }
// type Person struct {
// 	name string
// }

// func (dog Dog) Speak() string {
// 	return "Wookf!"
// }
// func (person Person) Speak() string {
// 	return "Hello"
// }
// func makeSound(s Speaker) {
// 	fmt.Println("", s.Speak())
// }

type Person struct {
	name string
	age  int
}

func main() {

	var x int = 10
	fmt.Println("", x)

	myArray := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(myArray); i++ {
		fmt.Print(" ", myArray[i])
	}

	mySlice := []int{}
	mySlice = append(mySlice, 112, 113, 11, 55)
	fmt.Println("\n", mySlice)
	fmt.Println("", mySlice[1])

	myMap := make(map[string]int)
	myMap["01"] = 10
	myMap["02"] = 10
	fmt.Println("", myMap)
	fmt.Println("", myMap["01"])

	person := Person{
		name: "TaiChi",
		age:  20,
	}
	fmt.Println("", person.name)
	fmt.Println("", person.age)

	person2 := make(map[string]Person)

	person2["0"] = Person{
		name: "Neo",
		age:  99,
	}
	person2["1"] = Person{
		name: "Neo1",
		age:  991,
	}
	fmt.Println("", person2)
	fmt.Println("", len(person2))
	fmt.Println("", person2["1"])
	// dog := Dog{
	// 	name: "Poppy",
	// }
	// person := Person{
	// 	name: "TaiChi",
	// }
	// makeSound(dog)
	// makeSound(person)
	// student := Student{
	// 	firstName: "TaiChi",
	// 	lastName:  "112",
	// }
	// fmt.Println("", student)
	// fullName := student.fullName()
	// fmt.Println("", fullName)

	// num1 := 5
	// num2 := 8
	// sumNum := add(num1, num2)
	// fmt.Println("sum : ", sumNum)
	// sum := add(1, 5)
	// fmt.Println("", sum)
	// sayHello("TaiChi", 2)
	// id := uuid.New()
	// fmt.Println("UUID : ", id)
	// fmt.Println()
	// fmt.Println("Hello world!")
	// Greeting()
	// Greeting()
	// Greeting()
	// other.Greeting()
	// other.Re_Test()
	// Hi("TaiChi\n")
	// a := 10
	// b := 2
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// Node1 := new(int)
	// if Node1 != nil {
	// 	fmt.Println("isn't null")
	// }
	// var score int = 112
	// var grade string

	// if score >= 80 {
	// 	grade = "A"
	// } else if score >= 70 {
	// 	grade = "B"
	// } else {
	// 	grade = "D"
	// }
	// fmt.Println("Your are grade : ", grade)

	// dayOfWeek := 5
	// switch dayOfWeek {
	// case 1:
	// 	fmt.Println("Sunday")
	// case 2:
	// 	fmt.Println("Monday")
	// case 3:
	// 	fmt.Println("Tuesday")
	// case 4:
	// 	fmt.Println("Wednesday")
	// case 5:
	// 	fmt.Println("Thursday")
	// case 6:
	// 	fmt.Println("Firday")
	// case 7:
	// 	fmt.Println("Satuday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// num1 := 10
	// num2 := 10
	// sumNum := num1 + num2
	// if sumNum >= 10 {
	// 	fmt.Println("Sum equal : True")
	// }

	// //Pre process
	// if sumNum := num1 + num2; sumNum >= 10 {
	// 	fmt.Println("Sum equal : duble true")
	// }

	// for loop
	// for i := 1; i < 5; i++ {
	// 	fmt.Println("Now var i : ", i, "And diaplay value i : ", i)
	// }

	// do while loop
	// i := 1
	// for {
	// 	fmt.Println("number : ", i)
	// 	i++
	// 	if i > 18 {
	// 		break
	// 	}
	// }
	// while loop
	// i := 2
	// for i < 10 {
	// 	fmt.Println("index : ", i)
	// 	i++
	// }
	// myArray := [4]int{1, 2, 3, 4}
	// fmt.Println("", myArray)

	// for i := 0; i < len(myArray); i++ {
	// 	fmt.Println("", myArray[i])
	// 	fmt.Println("", myArray)
	// }

	// mySlice1 := []int{1, 2, 3, 5}
	// mySlice1 := myArray[:]
	// fmt.Println("", mySlice1)
	// mySlice1 = append(mySlice1, 1, 3, 8, 7, 6)
	// fmt.Println("", mySlice1)

	// mySlice2 := []int{1}

	// for i := 0; i < len(mySlice1); i++ {
	// 	fmt.Println("", mySlice1[i])
	// }
	// for i := 0; i < len(mySlice2); i++ {
	// 	fmt.Println("", mySlice2[i])
	// }

	// fmt.Println("", len(mySlice1))
	// fmt.Println("", cap(mySlice1))
	// fmt.Println("", len(mySlice2))
	// fmt.Println("", cap(mySlice2))

	// subSlice := mySlice1[1:3] //start index 1 to index 2
	// fmt.Println("", subSlice)
	// fmt.Println("", len(subSlice))
	// fmt.Println("", cap(subSlice))

	// mySlice1 = append(mySlice1, 112, 112, 112)
	// fmt.Println("", mySlice1)

	// use map DSA

	// myMap := make(map[string]int)
	// myMap["apple"] = 5
	// myMap["banana"] = 112
	// myMap["orange"] = 8

	// fmt.Println("", myMap)
	// fmt.Println("", myMap["apple"])
	// fmt.Println("", myMap["banana"])
	// fmt.Println("", myMap["orange"])

	// for key, value := range myMap {
	// 	fmt.Println("", key, value)
	// }
	// // delete specifit map
	// delete(myMap, "apple")
	// for key, value := range myMap {
	// 	fmt.Println("", key, value)
	// }

	// val, ok := myMap["pear"]
	// if ok {
	// 	fmt.Println("", val)
	// } else {
	// 	fmt.Println(" Pear isn't found in map")
	// }
	// var student1 [1]Student
	// student1[0].name = "TaiChi"
	// student1[0].weight = 54
	// student1[0].height = 165
	// student1[0].grade = "A"
	// fmt.Println("", student1)
	// fmt.Println("", student1[0].name)
	// fmt.Println("", student1[0].weight)
	// fmt.Println("", student1[0].height)
	// fmt.Println("", student1[0].grade)

	// for i := 0; i < len(student1); i++ {
	// 	fmt.Println("", i+1)
	// 	fmt.Println("", student1[i].name)
	// 	fmt.Println("", student1[i].weight)
	// 	fmt.Println("", student1[i].height)
	// 	fmt.Println("", student1[i].grade)
	// }

	// student := make(map[string]Student)
	// student["01"] = Student{name: "TaiChi", weight: 54, height: 165, grade: "F"}
	// student["02"] = Student{name: "Art", weight: 99, height: 180, grade: "A"}
	// student["03"] = Student{name: "Tiger", weight: 60, height: 199, grade: "C"}

	// fmt.Println("", student)
	// fmt.Println("", student["01"])

	// for key, value := range student {
	// 	fmt.Println("", key, value)
	// }
	// var person1 Person

	// person1.name = "Hero"
	// person1.age = 112
	// person1.address = Address{
	// 	street:  "112 main gay",
	// 	city:    "bangkok",
	// 	zipCode: 112112,
	// }
	// fmt.Println("", person1)

	// myMap := make(map[string]Person)
	// myMap["01"] = Person{name: "TaiChi", age: 20, address: Address{street: "112 Women geek", city: "Tak", zipCode: 112112}}
	// fmt.Println("", myMap)

	// Alice := Person{
	// 	name: "Alic",
	// 	age:  18,
	// 	address: Address{
	// 		street:  "Left to Right in the moon",
	// 		city:    "ailen",
	// 		zipCode: 0000,
	// 	},
	// }
	// fmt.Println("", Alice)
}
