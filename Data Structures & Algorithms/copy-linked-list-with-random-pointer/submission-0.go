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
    list := new(Node)
    current2 := list

    for current != nil {
        current2.Val = current.Val
        m[current] = current2
        current = current.Next
        if current != nil {
            current2.Next = new(Node)
            current2 = current2.Next
        }
    }

    current = head
    current2 = list
    for current != nil {
        current2.Random = m[current.Random]
        current = current.Next
        current2 = current2.Next
    }

    return list
}
