/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToNew := map[*Node]*Node{}

    cur := head
    for cur != nil {
        new := &Node{Val:cur.Val}
        oldToNew[cur] = new
        cur = cur.Next
    }   

    cur = head 
    for cur != nil {
        oldToNew[cur].Next = oldToNew[cur.Next]
        oldToNew[cur].Random = oldToNew[cur.Random]
        cur = cur.Next
    }

    return oldToNew[head]
}