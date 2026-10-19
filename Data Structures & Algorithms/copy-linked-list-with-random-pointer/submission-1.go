/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    m := map[*Node]*Node{}

	list := new(Node)
	current := list

	for head != nil {
		node, exists := m[head]
		if !exists {
			node = &Node{
				Val: head.Val,
			}
			m[head] = node
		}

		if head.Random != nil {
			node.Random, exists = m[head.Random]
			if !exists {
				node.Random = &Node {
					Val: head.Random.Val,
				}
				m[head.Random] = node.Random
			}
		}

		current.Next = node
		current = current.Next

		head = head.Next
	}

	return list.Next
}
