/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    current := head
    size := 0
    for current != nil {
        current = current.Next
        size += 1
    }

    stack := make([]*ListNode, size)
    top := 0

    current = head
    for _ = range size/2 {
        current = current.Next
    }

    for current != nil {
        stack[top] = current
        current = current.Next
        top += 1
    }

    current = head

    for top > 1 {
        top -= 1
        topEle := stack[top]
        temp := current.Next
        current.Next = topEle
        topEle.Next = temp
        current = temp
    }

    current.Next = stack[top-1]
    current.Next.Next = nil
}
