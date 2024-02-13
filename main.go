package main

import "fmt"

func TransformToList(values ...int) *ListNode {
	var list *ListNode
	var last *ListNode

	for _, value := range values {
		element := &ListNode{Val: value}
		if list == nil {
			list = element
			last = element
		} else {
			last.Next = element
			last = element
		}
	}
	return list
}

func main() {
	first := TransformToList(9, 9, 9, 9, 9, 9, 9)
	second := TransformToList(9, 9, 9, 9)

	list := addTwoNumbers(first, second)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
