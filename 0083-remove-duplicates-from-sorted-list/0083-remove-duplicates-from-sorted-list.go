/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    nHead := &ListNode{}
    unique := nHead
    current := head
    for current != nil{        
        if unique == nHead || unique.Val != current.Val {
            unique.Next = current
            unique = unique.Next
            temp := unique.Next
            unique.Next = nil
            current = temp
        }else{
            current = current.Next    
        }        
    }
    
    return nHead.Next
}