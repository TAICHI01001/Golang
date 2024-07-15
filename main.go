package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node // test1 int, test2 Node, test3 int, test4 *Node
}

func main() {
	var test1 Node
	test1.data = 1
	test1.next = nil

	fmt.Println("test1.data: ", test1.data, "test1.next: ", test1.next)
}
