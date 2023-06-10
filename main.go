package main

import (
	"fmt"
	"strings"
)

type Value struct {
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

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Prepend(data *Value) {
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

func (l *LinkedList) Append(data *Value) {
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

func (l *LinkedList) Insert(index int, data *Value) error {
	if (index < 0 || index >= l.len) && l.len > 0 {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		l.Prepend(data)
	} else if index == l.len-1 {
		l.Append(data)
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

func (l *LinkedList) Pop() (*Value, error) {
	if l.head == nil {
		return nil, fmt.Errorf("cannot pop from empty list")
	}

	removedNode := l.tail

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.len = 0
		return removedNode.value, nil
	}

	currentNode := l.head
	for currentNode.next != l.tail {
		currentNode = currentNode.next
	}

	currentNode.next = nil
	l.tail = currentNode
	l.len--

	return removedNode.value, nil
}

func (l *LinkedList) Shift() (*Value, error) {
	if l.head == nil {
		return nil, fmt.Errorf("cannot shift from empty list")
	}

	removedNode := l.head
	l.head = l.head.next
	l.len--

	if l.head == nil {
		l.tail = nil
	}

	return removedNode.value, nil
}

func (l *LinkedList) Delete(index int) (*Value, error) {
	if index < 0 || index >= l.len {
		return nil, fmt.Errorf("index out of range")
	}

	if index == 0 {
		return l.Shift()
	}

	currentNode := l.head
	currentIndex := 0

	// Traverse to the node before the index
	for currentIndex < index {
		currentNode = currentNode.next
		currentIndex++
	}

	removedNode := currentNode.next
	currentNode.next = removedNode.next

	if removedNode == l.tail {
		l.tail = currentNode
	}

	l.len--

	return removedNode.value, nil
}

func (l *LinkedList) FirstIndexOf(data *Value) (int, error) {
	if l.head == nil {
		return -1, fmt.Errorf("cannot search empty list")
	}

	currentNode := l.head
	currentIndex := 0

	for currentNode != nil {
		if currentNode.value.data == data.data {
			return currentIndex, nil
		}

		currentNode = currentNode.next
		currentIndex++
	}

	return -1, fmt.Errorf("value not found")
}

func (l *LinkedList) Get(index int) (*Value, error) {
	if index < 0 || index > l.len-1 {
		return nil, fmt.Errorf("index out of range")
	}

	currentNode := l.head
	currentIndex := 0

	for currentIndex < index {
		currentNode = currentNode.next
		currentIndex++
	}

	return currentNode.value, nil
}

func (l *LinkedList) Contains(data *Value) bool {
	_, err := l.FirstIndexOf(data)
	return err == nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *LinkedList) ToArray() []*Value {
	values := make([]*Value, l.len)
	currentNode := l.head

	for i := 0; i < l.len; i++ {
		values[i] = currentNode.value
		currentNode = currentNode.next
	}

	return values
}

func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil {
		// Empty list or single node, no need to reverse
		return
	}

	var prevNode *Node
	currentNode := l.head
	l.tail = l.head

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	l.head = prevNode
}

func (l *LinkedList) String() string {
	var builder strings.Builder
	currentNode := l.head

	for currentNode != nil {
		fmt.Fprintf(&builder, "%v ", currentNode.value.data)
		currentNode = currentNode.next
	}

	return builder.String()
}

func main() {
	list := NewLinkedList()

	fmt.Printf("Is list empty? %t\n", list.IsEmpty())

	idxToPrepend := []int{5, 10, 15, 20, 25, 30, 35, 40}
	for _, index := range idxToPrepend {
		list.Prepend(&Value{data: index})
	}

	fmt.Println("...Prepended data")

	fmt.Printf("Is list empty? %t\n", list.IsEmpty())

	idxToAppend := []int{60, 65, 70}
	for _, index := range idxToAppend {
		list.Append(&Value{data: index})
	}

	fmt.Println()
	fmt.Printf("Data: %s\n", list)

	fmt.Println()
	idxToSearch := []int{4, 7, 9, 10, 20}
	for _, index := range idxToSearch {
		value, err := list.Get(index)
		if err != nil {
			fmt.Printf("Not Found index %d: %s\n", index, err)
		} else {
			fmt.Printf("Found index %d: %d\n", index, value.data)
		}
	}

	fmt.Println()
	fmt.Printf("Data: %s\n", list)

	// Insert a new value at index 3
	err := list.Insert(3, &Value{data: 99})
	if err != nil {
		fmt.Println("Insert Error:", err)
	}

	fmt.Println()
	fmt.Printf("Data: %s\n", list)

	// Pop the last value
	value, err := list.Pop()
	fmt.Println()
	if err != nil {
		fmt.Println("Pop Error:", err)
	} else {
		fmt.Printf("Popped Value: %d\n", value.data)
	}
	fmt.Printf("Data: %s\n", list)

	// Shift the first value
	value, err = list.Shift()
	fmt.Println()
	if err != nil {
		fmt.Println("Shift Error:", err)
	} else {
		fmt.Printf("Shifted Value: %d\n", value.data)
	}
	fmt.Printf("Data: %s\n", list)

	// Search for a value
	valueToSearch := 99
	index, err := list.FirstIndexOf(&Value{data: valueToSearch})
	fmt.Println()
	if err != nil {
		fmt.Printf("Search Error: %s\n", err)
	} else {
		fmt.Printf("Value %d found at index %d\n", valueToSearch, index)
	}

	// Search for a value that does not exist
	valueToSearch = 100
	index, err = list.FirstIndexOf(&Value{data: valueToSearch})
	if err != nil {
		fmt.Printf("Error on search index %d: %s\n", valueToSearch, err)
	} else {
		fmt.Printf("Value %d found at index %d\n", valueToSearch, index)
	}

	// Delete at index 3
	fmt.Println()
	fmt.Printf("Data: %s\n", list)
	value, err = list.Delete(3)
	if err != nil {
		fmt.Println("Delete Error:", err)
	} else {
		fmt.Printf("Deleted Value: %d\n", value.data)
	}
	fmt.Printf("Data: %s\n", list)

	// Contains 100?
	fmt.Println()
	fmt.Printf("Contains 100? %t\n", list.Contains(&Value{data: 100}))

	// clear the list
	fmt.Println()
	list.Clear()
	fmt.Println()
	fmt.Printf("Data: %s\n", list)

	// Contains 100?
	fmt.Println()
	fmt.Printf("Contains 100? %t\n", list.Contains(&Value{data: 100}))

	// Add some values to the list 1 to 10
	fmt.Println()
	fmt.Println("...Adding values 1 to 10")
	for i := 1; i <= 10; i++ {
		list.Append(&Value{data: i})
	}
	fmt.Printf("Data: %s\n", list)

	// Reverse the list
	fmt.Println()
	list.Reverse()
	fmt.Printf("Data reversed: %s\n", list)
}
