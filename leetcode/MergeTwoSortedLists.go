package leetcode

import "fmt"

/**
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := new(ListNode)
	rs := tmp
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
		fmt.Println(tmp.Val)
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return rs.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := new(ListNode)
	rs := tmp
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmp.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			tmp.Next = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
		fmt.Println(tmp.Val)
	}
	return rs.Next
}
