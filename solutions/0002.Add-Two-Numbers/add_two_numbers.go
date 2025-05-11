package leetcode

import (
	"leetcode/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{Val: 0}

	n1, n2, carry, current := 0, 0, 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: (n1 + n2 + carry) % 10,
		}

		current = current.Next
		carry = (n1 + n2 + carry) / 10

	}

	return head.Next

	// fmt.Printf("\n\n\n ------- \n \n \n")
	// fmt.Printf("head: %d", head.Val)
	// fmt.Printf("\n\n\n ------- \n \n \n")

	// fmt.Printf("L2 VAL IS: %d", l2.Val)

	// fmt.Printf("\n\n\n ------- \n \n \n")

	// for l1 != nil {
	// 	l1 = l1.Next

	// }

	// fmt.Printf("\n\n\n ------- \n \n \n")

	// return head
}
