package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l *LinkedList) prepend(data int) {
	newNode := &Node{value: data, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}

	l.len++
}

func (l *LinkedList) append(data int) {
	newNode := &Node{value: data, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.len++
}

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.len != 0 {
		fmt.Printf("%d ", toPrint.value)
		toPrint = toPrint.next
		l.len--
	}
	fmt.Println()
}

func main() {
	myLinkedList := LinkedList{}
	myLinkedList.prepend(5)
	myLinkedList.prepend(10)
	myLinkedList.prepend(15)
	myLinkedList.prepend(20)
	myLinkedList.prepend(25)
	myLinkedList.prepend(30)
	myLinkedList.prepend(35)
	myLinkedList.prepend(40)

	myLinkedList.append(45)
	myLinkedList.append(65)
	myLinkedList.append(15)

	myLinkedList.printListData()
	fmt.Printf("Length: %d\n", myLinkedList.len)
}
