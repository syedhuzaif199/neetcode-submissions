/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
	head := ListNode{}
	current := &head
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		current.Val = sum % 10
		fmt.Println("12, sum =", current.Val, "carry =", carry)
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil && l2 != nil {
			current.Next = new(ListNode)
			current = current.Next
		}
	}

	for l1 != nil {
		current.Next = new(ListNode)
		current = current.Next
		sum := l1.Val + carry
		carry = sum / 10
		current.Val = sum % 10
		fmt.Println("1, sum =", current.Val, "carry =", carry)
		l1 = l1.Next
	}

	for l2 != nil {
		current.Next = new(ListNode)
		current = current.Next
		sum := l2.Val + carry
		carry = sum / 10
		current.Val = sum % 10
		fmt.Println("2, sum =", current.Val, "carry =", carry)
		l2 = l2.Next
		if l2 != nil {
			current.Next = new(ListNode)
			current = current.Next
		}
	}

	if carry != 0 {
		current.Next = new(ListNode)
		current.Next.Val = carry
	}

	return &head
}
