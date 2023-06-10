package main

import "fmt"

type Value struct {
	data int
}

type Node struct {
	value *Value
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (dl *DoublyLinkedList) Prepend(data *Value) {
	newNode := &Node{value: data, next: dl.head, prev: nil}

	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.head.prev = newNode
		dl.head = newNode
	}

	dl.len++
}

func (dl *DoublyLinkedList) Append(data *Value) {
	newNode := &Node{value: data, next: nil, prev: dl.tail}

	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		dl.tail.next = newNode
		dl.tail = newNode
	}

	dl.len++
}

func (dl *DoublyLinkedList) Insert(index int, data *Value) error {
	if (index < 0 || index >= dl.len) && dl.len > 0 {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		dl.Prepend(data)
	} else if index == dl.len-1 {
		dl.Append(data)
	} else {
		newNode := &Node{value: data}

		currentNode := dl.head
		currentIndex := 0

		// Traverse to the node before the index
		for currentIndex < index-1 {
			currentNode = currentNode.next
			currentIndex++
		}

		newNode.next = currentNode.next
		newNode.prev = currentNode
		currentNode.next.prev = newNode
		currentNode.next = newNode

		dl.len++
	}

	return nil
}

func (dl *DoublyLinkedList) Delete(index int) (*Value, error) {
	if index < 0 || index >= dl.len {
		return nil, fmt.Errorf("index out of range")
	}

	if index == 0 {
		return dl.DeleteHead()
	} else if index == dl.len-1 {
		return dl.DeleteTail()
	} else {
		currentNode := dl.head
		currentIndex := 0

		// Traverse to the node at the given index
		for currentIndex < index {
			currentNode = currentNode.next
			currentIndex++
		}

		currentNode.prev.next = currentNode.next
		currentNode.next.prev = currentNode.prev
		dl.len--

		return currentNode.value, nil
	}
}

func (dl *DoublyLinkedList) DeleteHead() (*Value, error) {
	if dl.head == nil {
		return nil, fmt.Errorf("cannot delete from empty list")
	}

	value := dl.head.value
	dl.head = dl.head.next

	if dl.head == nil {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}

	dl.len--

	return value, nil
}

func (dl *DoublyLinkedList) DeleteTail() (*Value, error) {
	if dl.tail == nil {
		return nil, fmt.Errorf("cannot delete from empty list")
	}

	value := dl.tail.value
	dl.tail = dl.tail.prev

	if dl.tail == nil {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}

	dl.len--

	return value, nil
}

func (dl *DoublyLinkedList) Get(index int) (*Value, error) {
	if index < 0 || index >= dl.len {
		return nil, fmt.Errorf("index out of range")
	}

	currentNode := dl.head
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

func (dl *DoublyLinkedList) Reverse() {
	if dl.head == nil || dl.head.next == nil {
		// Empty list or single node, no need to reverse
		return
	}

	currentNode := dl.head
	var prevNode *Node

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = prevNode
		currentNode.prev = nextNode
		prevNode = currentNode
		currentNode = nextNode
	}

	dl.head, dl.tail = dl.tail, dl.head
}

func (dl *DoublyLinkedList) ToArray() []*Value {
	values := make([]*Value, 0, dl.len)
	currentNode := dl.head

	for currentNode != nil {
		values = append(values, currentNode.value)
		currentNode = currentNode.next
	}

	return values
}

func (dll *DoublyLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.len = 0
}

func (dl *DoublyLinkedList) PrintData() {
	if dl.head == nil {
		fmt.Println("DoublyLinkedList is empty")
		return
	}

	values := dl.ToArray()

	fmt.Print("Data: ")
	for _, value := range values {
		fmt.Printf("%v ", value.data)
	}
	fmt.Println()
	fmt.Printf("Length: %d\n", dl.len)
}

func main() {
	dl := &DoublyLinkedList{}

	fmt.Printf("Is doubly linked list empty? %t\n", dl.head == nil)

	values := []int{5, 10, 15, 20, 25, 30, 35, 40}

	for _, value := range values {
		dl.Append(&Value{data: value})
	}

	fmt.Println("...Appended values to the doubly linked list")

	fmt.Printf("Is doubly linked list empty? %t\n", dl.head == nil)

	fmt.Println()
	dl.PrintData()

	// Insert a new value at index 3
	err := dl.Insert(3, &Value{data: 99})
	if err != nil {
		fmt.Println("Insert Error:", err)
	}

	fmt.Println()
	dl.PrintData()

	// Delete at index 2
	value, err := dl.Delete(2)
	if err != nil {
		fmt.Println("Delete Error:", err)
	} else {
		fmt.Printf("Deleted Value: %d\n", value.data)
	}
	fmt.Println()
	dl.PrintData()

	// Get value at index 1
	index := 1
	value, err = dl.Get(index)
	if err != nil {
		fmt.Printf("Get Error at index %d: %s\n", index, err)
	} else {
		fmt.Printf("Value at index %d: %d\n", index, value.data)
	}
	fmt.Println()

	// Convert the doubly linked list to an array
	array := dl.ToArray()
	fmt.Println("Doubly Linked List as Array:")
	for i, value := range array {
		fmt.Printf("Index %d: %d\n", i, value.data)
	}
	fmt.Println()

	// Reverse the doubly linked list
	dl.Reverse()
	fmt.Println()
	dl.PrintData()
}
