/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{}
    current := &head
    for list1 != nil && list2 != nil {
        current.Next = new(ListNode)
        current = current.Next
        if list1.Val < list2.Val {
            current.Val = list1.Val
            list1 = list1.Next
        } else {
            current.Val = list2.Val
            list2 = list2.Next
        }
    }

    for list1 != nil {
        current.Next = new(ListNode)
        current = current.Next
        current.Val = list1.Val
        list1 = list1.Next
    }
    for list2 != nil {
        current.Next = new(ListNode)
        current = current.Next
        current.Val = list2.Val
        list2 = list2.Next
    }
    return head.Next
}
