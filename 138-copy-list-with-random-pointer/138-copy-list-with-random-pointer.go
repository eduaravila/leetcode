/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    copy_list := make(map[*Node]*Node)
    
    current := head
    for current != nil {
        copy_list[current] = &Node{Val:current.Val}
        current = current.Next
    }
    current = head
    for current != nil{
        copy_list[current].Next= copy_list[current.Next]
        copy_list[current].Random= copy_list[current.Random]
        current = current.Next
    }
    return copy_list[head]
}