package main

import "fmt"

type GenericList []interface{}

func (list *GenericList) RemoveIndex(index int) interface{} {
	l := *list
	removed := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return removed
}

func (list *GenericList) RemoveFirst() interface{} {
	return list.RemoveIndex(0)
}

func (list *GenericList) Pop() interface{} {
	return list.RemoveIndex(len(*list) - 1)
}

func main() {
	list := GenericList{1, "Coffee", 42, true, 23, "Book", 3.14, false}

	fmt.Printf("Original List:\n%v\n\n", list)

	fmt.Printf("Removing from the beginning: %v, after removal: %v\n\n", list.RemoveFirst(), list)
	fmt.Printf("Removing from the end: %v, after removal: %v", list.Pop(), list)
}
