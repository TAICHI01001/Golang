package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value int) {
	var newNode = new(Node)
	newNode.data = value
	newNode.next = l.head
	l.head = newNode
}
func (l *LinkedList) Print() {
	var temp = l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

	fmt.Println()
}
func main() {
	var l = new(LinkedList)
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	l.Print()
}
