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

type Queue struct {
	front *Node
	rear  *Node
	len   int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(data *Value) {
	newNode := &Node{value: data, next: nil}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

	q.len++
}

func (q *Queue) dequeue() (*Value, error) {
	if q.front == nil {
		return nil, fmt.Errorf("cannot dequeue from empty queue")
	}

	value := q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.len--

	return value, nil
}

func (q *Queue) isEmpty() bool {
	return q.front == nil
}

func (q *Queue) clear() {
	q.front = nil
	q.rear = nil
	q.len = 0
}

func (q *Queue) String() string {
	if q.front == nil {
		return "Queue is empty"
	}

	var builder strings.Builder
	currentNode := q.front

	for currentNode != nil {
		fmt.Fprintf(&builder, "%v ", currentNode.value.data)
		currentNode = currentNode.next
	}

	return builder.String()
}

func main() {
	queue := NewQueue()

	fmt.Printf("Is queue empty? %t\n", queue.isEmpty())

	values := []int{5, 10, 15, 20, 25, 30, 35, 40}

	for _, value := range values {
		queue.enqueue(&Value{data: value})
	}

	fmt.Println("...Enqueued values to the queue")

	fmt.Printf("Is queue empty? %t\n", queue.isEmpty())

	fmt.Printf("Queue: %s\n", queue)

	// Dequeue until queue is empty
	for !queue.isEmpty() {
		value, _ := queue.dequeue()
		fmt.Printf("Dequeued: %d\n", value.data)
		fmt.Printf("Queue: %s\n", queue)
	}

	// Clear the queue
	fmt.Println()
	queue.clear()
	fmt.Println()
	fmt.Printf("Queue: %s\n", queue)
}
