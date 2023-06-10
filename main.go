package main

import "fmt"

type Value struct {
	// Define fields based on your requirements for the value object
	// For example, if you want to store an integer value:
	// data int
	//
	// Or, if you want to store a string value:
	// data string
	//
	// You can customize the fields based on the specific data type you need.
	data int
}

type Node struct {
	value *Value
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l *LinkedList) prepend(data *Value) {
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

func (l *LinkedList) append(data *Value) {
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

func (l *LinkedList) insert(index int, data *Value) error {
	if (index < 0 || index >= l.len) && l.len > 0 {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		l.prepend(data)
	} else if index == l.len-1 {
		l.append(data)
	} else {
		newNode := &Node{value: data, next: nil}

		currentNode := l.head
		currentIndex := 0

		// Traverse to the node before the index
		for currentIndex < index-1 {
			currentNode = currentNode.next
			currentIndex++
		}

		newNode.next = currentNode.next
		currentNode.next = newNode

		l.len++
	}

	return nil
}

func (l *LinkedList) traverseIndex(index int) (*Value, error) {
	if index < 0 || index >= l.len {
		return nil, fmt.Errorf("index out of range")
	}

	currentNode := l.head
	currentIndex := 0

	for currentNode != nil {
		if currentIndex == index {
			return currentNode.value, nil
		}

		currentNode = currentNode.next
		currentIndex++
	}

	return nil, fmt.Errorf("index out of range")
}

func (ll *LinkedList) printData() {
	if ll.head == nil {
		fmt.Println("LinkedList is empty")
		return
	}

	currentNode := ll.head

	fmt.Print("Data: ")
	for currentNode != nil {
		fmt.Printf("%v ", currentNode.value.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	idxToPrepend := []int{5, 10, 15, 20, 25, 30, 35, 40}
	for _, index := range idxToPrepend {
		list.prepend(&Value{data: index})
	}

	idxToAppend := []int{60, 65, 70}
	for _, index := range idxToAppend {
		list.append(&Value{data: index})
	}

	idxToSearch := []int{4, 7, 10, 20}
	for _, index := range idxToSearch {
		value, err := list.traverseIndex(index)
		if err != nil {
			fmt.Printf("Value at index %d: %s\n", index, err)
		} else {
			fmt.Printf("Value at index %d: %d\n", index, value.data)
		}
	}

	fmt.Println()
	list.printData()
	fmt.Printf("Length: %d\n", list.len)

	// Insert a new value at index 3
	err := list.insert(3, &Value{data: 99})
	if err != nil {
		fmt.Println("Insert Error:", err)
	}

	fmt.Println()
	list.printData()
	fmt.Printf("Length: %d\n", list.len)
}
