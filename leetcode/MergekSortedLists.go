package leetcode

/**
Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	start, end := 0, len(lists)-1
	for start < end {
		mid := (end + start - 1) / 2
		for i := 0; i <= mid; i++ {
			lists[i] = Merge(lists[i], lists[end-i])
		}
		end = (end + start) / 2
	}
	return lists[0]
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = Merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = Merge(l1, l2.Next)
		return l2
	}
}
