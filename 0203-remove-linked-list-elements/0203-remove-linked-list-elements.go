/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {    
    nhead := &ListNode{}
    prev := nhead
    current := head 
    for current != nil{
        if current.Val == val {
            prev.Next = current.Next            
        }else{
            prev.Next = current
            prev = current
        }
        current = current.Next
    }    
    
    return nhead.Next
}


