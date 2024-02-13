package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	t.Run("equalLength", func(t *testing.T) {
		first := TransformToList(2, 4, 3)
		second := TransformToList(5, 6, 4)
		answer := TransformToList(7, 0, 8)

		result := addTwoNumbers(first, second)

		for answer != nil && result != nil {
			if answer.Val != result.Val {
				t.Errorf("%d != %d", answer.Val, result.Val)
			}
			answer = answer.Next
			result = result.Next
		}
	})

	t.Run("zeroLength", func(t *testing.T) {
		first := TransformToList(0)
		second := TransformToList(0)
		answer := TransformToList(0)

		result := addTwoNumbers(first, second)

		for answer != nil && result != nil {
			if answer.Val != result.Val {
				t.Errorf("%d != %d", answer.Val, result.Val)
			}
			answer = answer.Next
			result = result.Next
		}
	})

	t.Run("differentLength", func(t *testing.T) {
		first := TransformToList(9, 9, 9, 9, 9, 9, 9)
		second := TransformToList(9, 9, 9, 9)
		answer := TransformToList(8, 9, 9, 9, 0, 0, 0, 1)

		result := addTwoNumbers(first, second)

		for answer != nil && result != nil {
			if answer.Val != result.Val {
				t.Errorf("%d != %d", answer.Val, result.Val)
			}
			answer = answer.Next
			result = result.Next
		}
	})
}
