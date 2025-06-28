package leetcode

import (
	"leetcode/structures"
	"sort"
)

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var listOfInt []int

	for list1 != nil {
		value := list1.Val
		list1 = list1.Next

		listOfInt = append(listOfInt, value)
	}

	for list2 != nil {
		value := list2.Val
		list2 = list2.Next

		listOfInt = append(listOfInt, value)
	}

	sort.IntSlice(listOfInt).Sort()

	return Ints2List(listOfInt)
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
