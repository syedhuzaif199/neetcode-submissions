/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

    m := map[*Node]*Node{}

	current := head
	for current != nil {
		m[current] = &Node{
			Val: current.Val,
		}
		current = current.Next
	}

	list := m[head]
	current = head
	for current != nil {
		list.Next = m[current.Next]
		list.Random = m[current.Random]
		list = list.Next
		current = current.Next
	}

	return m[head]
}
