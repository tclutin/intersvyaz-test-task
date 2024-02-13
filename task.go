package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToEndList(list *ListNode, element *ListNode) *ListNode {
	if list == nil {
		list = element
	} else {
		current := list
		for current.Next != nil {
			current = current.Next
		}
		current.Next = element
	}
	return list
}

func addTwoNumbersSlow(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, digit int
	var result, element *ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += digit % 10

		if sum >= 10 {
			digit = 1
			element = &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			result = AddToEndList(result, element)
		} else {
			digit = 0
			element = &ListNode{
				Val:  sum,
				Next: nil,
			}
			result = AddToEndList(result, element)
		}
		sum = 0
	}

	if digit == 1 {
		element = &ListNode{
			Val:  1,
			Next: nil,
		}
		result = AddToEndList(result, element)
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, digit int
	var element *ListNode

	var last = &ListNode{}
	var result = last

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += digit

		if sum >= 10 {
			digit = 1
			element = &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			last.Next = element
			last = element

		} else {
			digit = 0
			element = &ListNode{
				Val:  sum,
				Next: nil,
			}
			last.Next = element
			last = element
		}
		sum = 0
	}

	if digit == 1 {
		element = &ListNode{
			Val:  1,
			Next: nil,
		}
		last.Next = element
		last = element
	}
	return result.Next
}
