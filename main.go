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

type Stack struct {
	top *Node
	len int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(data *Value) {
	newNode := &Node{value: data, next: s.top}
	s.top = newNode
	s.len++
}

func (s *Stack) pop() (*Value, error) {
	if s.top == nil {
		return nil, fmt.Errorf("cannot pop from empty stack")
	}

	value := s.top.value
	s.top = s.top.next
	s.len--

	return value, nil
}

func (s *Stack) peek() (*Value, error) {
	if s.top == nil {
		return nil, fmt.Errorf("cannot peek an empty stack")
	}

	return s.top.value, nil
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) clear() {
	s.top = nil
	s.len = 0
}

func (s *Stack) String() string {
	if s.top == nil {
		return "Stack is empty"
	}

	var builder strings.Builder
	currentNode := s.top

	for currentNode != nil {
		fmt.Fprintf(&builder, "%v ", currentNode.value.data)
		currentNode = currentNode.next
	}

	return builder.String()
}

func main() {
	stack := NewStack()

	fmt.Printf("Is stack empty? %t\n", stack.isEmpty())

	values := []int{5, 10, 15, 20, 25, 30, 35, 40}

	for _, value := range values {
		stack.push(&Value{data: value})
	}

	fmt.Println("...Pushed values to the stack")

	fmt.Printf("Is stack empty? %t\n", stack.isEmpty())

	fmt.Printf("Stack: %s\n", stack)

	// Pop the top value
	value, err := stack.pop()
	fmt.Println()
	if err != nil {
		fmt.Println("Pop Error:", err)
	} else {
		fmt.Printf("Popped Value: %d\n", value.data)
	}
	fmt.Printf("Stack: %s\n", stack)

	// Peek the top value
	value, err = stack.peek()
	fmt.Println()
	if err != nil {
		fmt.Println("Peek Error:", err)
	} else {
		fmt.Printf("Peeked Value: %d\n", value.data)
	}
	fmt.Printf("Stack: %s\n", stack)

	// Clear the stack
	fmt.Println()
	stack.clear()
	fmt.Println()
	fmt.Printf("Stack: %s\n", stack)
}
