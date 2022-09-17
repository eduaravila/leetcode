/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    
    newHead := &ListNode{}
    currentReversed := newHead
    
    l,r := head,head
    current := 1 
    for r != nil{
        if current % k == 0{
            
            next := r.Next
            r.Next = nil
            tempHead := &ListNode{}
            reverseList(l,nil,tempHead)
            
            l.Next = next
            currentReversed.Next = tempHead.Next
            currentReversed = l
            
            r=l
            l = l.Next
            
        }        
        r = r.Next
        current++
    }
    
    
    return newHead.Next
}

func reverseList(current, prev , head *ListNode) {
    if current == nil{        
        head.Next = prev
        return 
    }    
    reverseList(current.Next,current,head)    
    current.Next = prev    
}