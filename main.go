package main

import "fmt"

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

func (l *LinkedList) pop() (*Value, error) {
	if l.head == nil {
		return nil, fmt.Errorf("cannot pop from empty list")
	}

	currentNode := l.head
	previousNode := l.head

	for currentNode.next != nil {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	previousNode.next = nil
	l.tail = previousNode
	l.len--

	return currentNode.value, nil
}

func (l *LinkedList) shift() (*Value, error) {
	if l.head == nil {
		return nil, fmt.Errorf("cannot shift from empty list")
	}

	removedNode := l.head
	l.head = l.head.next
	l.len--

	return removedNode.value, nil
}

func (l *LinkedList) delete(index int) (*Value, error) {
	if index < 0 || index >= l.len {
		return nil, fmt.Errorf("index out of range")
	}

	if index == 0 {
		return l.shift()
	} else if index == l.len-1 {
		return l.pop()
	} else {
		currentNode := l.head
		currentIndex := 0

		// Traverse to the node before the index
		for currentIndex < index-1 {
			currentNode = currentNode.next
			currentIndex++
		}

		removedNode := currentNode.next
		currentNode.next = removedNode.next
		l.len--

		return removedNode.value, nil
	}
}

func (l *LinkedList) firstIndexOf(data *Value) (int, error) {
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

func (l *LinkedList) get(index int) (*Value, error) {
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

func (l *LinkedList) contains(data *Value) bool {
	if l.head == nil {
		return false
	}

	currentNode := l.head

	for currentNode != nil {
		if currentNode.value.data == data.data {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *LinkedList) toArray() []*Value {
	values := make([]*Value, 0, l.len)
	currentNode := l.head

	for currentNode != nil {
		values = append(values, currentNode.value)
		currentNode = currentNode.next
	}

	return values
}

func (l *LinkedList) printData() {
	if l.head == nil {
		fmt.Println("LinkedList is empty")
		return
	}

	values := l.toArray()

	fmt.Print("Data: ")
	for _, value := range values {
		fmt.Printf("%v ", value.data)
	}
	fmt.Println()
	fmt.Printf("Length: %d\n", l.len)
}

func main() {
	list := LinkedList{}

	fmt.Printf("Is list empty? %t\n", list.isEmpty())

	idxToPrepend := []int{5, 10, 15, 20, 25, 30, 35, 40}
	for _, index := range idxToPrepend {
		list.prepend(&Value{data: index})
	}

	fmt.Println("...Prepended data")

	fmt.Printf("Is list empty? %t\n", list.isEmpty())

	idxToAppend := []int{60, 65, 70}
	for _, index := range idxToAppend {
		list.append(&Value{data: index})
	}

	fmt.Println()
	idxToSearch := []int{4, 7, 10, 20}
	for _, index := range idxToSearch {
		value, err := list.get(index)
		if err != nil {
			fmt.Printf("Not Found index %d: %s\n", index, err)
		} else {
			fmt.Printf("Found index %d: %d\n", index, value.data)
		}
	}

	fmt.Println()
	list.printData()

	// Insert a new value at index 3
	err := list.insert(3, &Value{data: 99})
	if err != nil {
		fmt.Println("Insert Error:", err)
	}

	fmt.Println()
	list.printData()

	// Pop the last value
	value, err := list.pop()
	fmt.Println()
	if err != nil {
		fmt.Println("Pop Error:", err)
	} else {
		fmt.Printf("Popped Value: %d\n", value.data)
	}
	list.printData()

	// Shift the first value
	value, err = list.shift()
	fmt.Println()
	if err != nil {
		fmt.Println("Shift Error:", err)
	} else {
		fmt.Printf("Shifted Value: %d\n", value.data)
	}
	list.printData()

	// Search for a value
	valueToSearch := 99
	index, err := list.firstIndexOf(&Value{data: valueToSearch})
	fmt.Println()
	if err != nil {
		fmt.Printf("Search Error: %s\n", err)
	} else {
		fmt.Printf("Value %d found at index %d\n", valueToSearch, index)
	}

	// Search for a value that does not exist
	valueToSearch = 100
	index, err = list.firstIndexOf(&Value{data: valueToSearch})
	if err != nil {
		fmt.Printf("Error on search index %d: %s\n", valueToSearch, err)
	} else {
		fmt.Printf("Value %d found at index %d\n", valueToSearch, index)
	}

	// Delete at index 3
	fmt.Println()
	list.printData()
	value, err = list.delete(3)
	if err != nil {
		fmt.Println("Delete Error:", err)
	} else {
		fmt.Printf("Deleted Value: %d\n", value.data)
	}
	list.printData()

	// Contains 100?
	fmt.Println()
	fmt.Printf("Contains 100? %t\n", list.contains(&Value{data: 100}))

	// clear the list
	fmt.Println()
	list.clear()
	fmt.Println()
	list.printData()

	// Contains 100?
	fmt.Println()
	fmt.Printf("Contains 100? %t\n", list.contains(&Value{data: 100}))
}
