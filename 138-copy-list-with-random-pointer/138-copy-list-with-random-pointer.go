/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil{
        return head
    }
    current := head
    head_copy := &Node{}
    current_copy := head_copy
    for current != nil {        
        current_copy.Val= current.Val
        current = current.Next
        if current != nil {
            current_copy.Next = &Node{}
            current_copy = current_copy.Next
        }        
    }
    
    current = head
    current_copy = head_copy
    
    
    for current!= nil{
        
        current_copy_temp := current_copy.Next
        
        current_copy.Next = current.Next
        current.Next= current_copy        
        current_copy.Random = current.Random
        current = current_copy.Next
        current_copy = current_copy_temp
        
    }    
    
    current_copy = head_copy
    for current_copy != nil {
        
        if current_copy.Random != nil{
            current_copy.Random = current_copy.Random.Next
        }
        if current_copy.Next == nil{
            break
        }
        current_copy = current_copy.Next.Next
    }
    
    current = head
    current_copy = head.Next
    for current != nil && current_copy != nil {
        
    
        current.Next = current_copy.Next
        if current_copy.Next ==nil{
            break
        }
        current_copy.Next = current.Next.Next
        current = current.Next
        current_copy= current_copy.Next
    }
    return head_copy
}