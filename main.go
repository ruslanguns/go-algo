package main

import (
	"fmt"
)

type Value struct {
	data int
}

type Node struct {
	value *Value
	next  *Node
	prev  *Node
}

type Deque struct {
	front *Node
	rear  *Node
	len   int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) insertFront(data *Value) {
	newNode := &Node{value: data, next: d.front, prev: nil}

	if d.front == nil {
		d.front = newNode
		d.rear = newNode
	} else {
		d.front.prev = newNode
		d.front = newNode
	}

	d.len++
}

func (d *Deque) insertLast(data *Value) {
	newNode := &Node{value: data, next: nil, prev: d.rear}

	if d.rear == nil {
		d.front = newNode
		d.rear = newNode
	} else {
		d.rear.next = newNode
		d.rear = newNode
	}

	d.len++
}

func (d *Deque) deleteFront() (*Value, error) {
	if d.front == nil {
		return nil, fmt.Errorf("cannot delete from empty deque")
	}

	value := d.front.value
	d.front = d.front.next

	if d.front == nil {
		d.rear = nil
	} else {
		d.front.prev = nil
	}

	d.len--

	return value, nil
}

func (d *Deque) deleteLast() (*Value, error) {
	if d.rear == nil {
		return nil, fmt.Errorf("cannot delete from empty deque")
	}

	value := d.rear.value
	d.rear = d.rear.prev

	if d.rear == nil {
		d.front = nil
	} else {
		d.rear.next = nil
	}

	d.len--

	return value, nil
}

func (d *Deque) isEmpty() bool {
	return d.front == nil
}

func (d *Deque) clear() {
	d.front = nil
	d.rear = nil
	d.len = 0
}

func (d *Deque) String() string {
	if d.front == nil {
		return "Deque is empty"
	}

	currentNode := d.front
	var result string

	for currentNode != nil {
		result += fmt.Sprintf("%v ", currentNode.value.data)
		currentNode = currentNode.next
	}

	return result
}

func main() {
	deque := NewDeque()

	fmt.Printf("Is deque empty? %t\n", deque.isEmpty())

	values := []int{5, 10, 15, 20, 25, 30, 35, 40}

	for _, value := range values {
		deque.insertLast(&Value{data: value})
	}

	fmt.Println("...Inserted values into the deque")

	fmt.Printf("Is deque empty? %t\n", deque.isEmpty())

	fmt.Printf("Deque: %s\n", deque)

	// Delete from the front
	value, err := deque.deleteFront()
	fmt.Println()
	if err != nil {
		fmt.Println("Delete Front Error:", err)
	} else {
		fmt.Printf("Deleted Value from Front: %d\n", value.data)
	}
	fmt.Printf("Deque: %s\n", deque)

	// Delete from the last
	value, err = deque.deleteLast()
	fmt.Println()
	if err != nil {
		fmt.Println("Delete Last Error:", err)
	} else {
		fmt.Printf("Deleted Value from Last: %d\n", value.data)
	}
	fmt.Printf("Deque: %s\n", deque)

	// Clear the deque
	fmt.Println()
	deque.clear()
	fmt.Printf("Deque: %s\n", deque)

	fmt.Printf("Is deque empty? %t\n", deque.isEmpty())
}
