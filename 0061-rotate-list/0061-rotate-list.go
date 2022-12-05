/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getLen(head *ListNode)int{
    var res int
    current := head
    for current != nil{
        res++
        current = current.Next
    }
    return res    
}

func rotateRight(head *ListNode, k int) *ListNode {
    nodes := make([]*ListNode,getLen(head))    
    current := head
    var i int
    for current != nil{        
        nodes[(i+k)%len(nodes)] = current
        current = current.Next
        i++
    }
    
    nhead := &ListNode{}
    current = nhead
    for _, node := range nodes{
        current.Next = node
        node.Next = nil
        current =current.Next
    }    
    return nhead.Next    
}

