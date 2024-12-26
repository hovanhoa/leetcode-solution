/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := map[*Node]*Node{}
    cur := head

    for cur != nil {
        copy := &Node{Val: cur.Val}
        oldToCopy[cur] = copy
        cur = cur.Next
    }

    cur = head
    for cur != nil {
        copy := oldToCopy[cur]
        copy.Next = oldToCopy[cur.Next]
        copy.Random = oldToCopy[cur.Random]
        cur = cur.Next
    }

    return oldToCopy[head]
}