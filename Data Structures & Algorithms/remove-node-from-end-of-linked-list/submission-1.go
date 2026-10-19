/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    first := head
	for _ = range n {
		first = first.Next
	}

	second := head

	if first == nil {
		head = head.Next
		return head
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return head

}
