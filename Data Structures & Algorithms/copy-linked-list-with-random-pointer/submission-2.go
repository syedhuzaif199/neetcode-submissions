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

	current = head
	for current != nil {
		m[current].Next = m[current.Next]
		m[current].Random = m[current.Random]
		current = current.Next
	}

	return m[head]
}
