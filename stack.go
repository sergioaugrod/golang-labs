package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Enqueue(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Dequeue() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Stack empty!")
	}

	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}

func main() {
	stack := Stack{}
	fmt.Println("Stack created with size:", stack.Size())
	fmt.Println("Stack empty?", stack.Empty())

	stack.Enqueue("Go")
	stack.Enqueue(20)
	stack.Enqueue(2.23)
	stack.Enqueue("End")

	fmt.Println("Size after enqueue 4 items:", stack.Size())
	fmt.Println("Stack empty?", stack.Empty())

	for !stack.Empty() {
		v, _ := stack.Dequeue()
		fmt.Println("Dequeue:", v)
		fmt.Println("Size:", stack.Size())
		fmt.Println("Empty?:", stack.Empty())
	}

	_, err := stack.Dequeue()

	if err != nil {
		fmt.Println("Error:", err)
	}
}
