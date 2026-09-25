/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    if head.Next == nil || head.Next.Next == nil {
        return false
    }

    a, b := head.Next, head.Next.Next

    for a != b {
        if a.Next == nil || b.Next == nil || b.Next.Next == nil {
            return false
        }
        b = b.Next.Next

        a = a.Next
    }

    return true

}
