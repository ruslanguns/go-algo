package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	len  int
}

func (l *linkedList) prepend(data int) {
	newNode := node{data: data}
	newNode.next = l.head
	l.head = &newNode
	l.len++
}

func main() {
	myLinkedList := linkedList{}
	myLinkedList.prepend(5)
	myLinkedList.prepend(10)
	myLinkedList.prepend(15)
	myLinkedList.prepend(20)
	myLinkedList.prepend(25)
	myLinkedList.prepend(30)
	myLinkedList.prepend(35)
	myLinkedList.prepend(40)

	fmt.Println(myLinkedList)
}
