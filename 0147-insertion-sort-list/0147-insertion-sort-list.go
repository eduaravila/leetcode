/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    inf := int(^uint(0)>>1)
    nhead := &ListNode{Val:-inf}
    nhead.Next = head    
    current := head.Next
    head.Next =nil
    
    for current != nil{
        temp := current.Next
        current.Next = nil
        
        prev := nhead
        currenth := nhead.Next        
        for currenth != nil{
            if current.Val <= currenth.Val{                                                               
                break
            }
            prev = prev.Next
            currenth = currenth.Next
        }
        prev.Next = current
        current.Next = currenth 
        current = temp
    }
    
    return nhead.Next
}

