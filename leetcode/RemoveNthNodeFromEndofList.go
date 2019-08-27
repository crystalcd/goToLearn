package leetcode

/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeMap := make(map[int]*ListNode)
	i := 0
	for head != nil {
		nodeMap[i] = head
		head = head.Next
		i++
	}
	actN := i - n
	if actN == 0 {
		return nodeMap[1]
	} else if actN < 0 {
		return nodeMap[0]
	}
	nodeMap[actN-1].Next = nodeMap[actN+1]
	return nodeMap[0]
}
