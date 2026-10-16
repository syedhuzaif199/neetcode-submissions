/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    size := 0
    current := head
    for current != nil {
        size += 1
        current = current.Next
    }

    current = head
    var prev *ListNode
    for _ = range size - n {
        prev = current
        current = current.Next
    }

    if prev != nil {
        prev.Next = current.Next
    } else {
        head = current.Next
    }

    return head
}
