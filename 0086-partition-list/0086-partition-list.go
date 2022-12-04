/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    partition :=&ListNode{}
    nhead := partition
    current := head
    
    for current != nil{
        if current.Val < x{
            nhead.Next = &ListNode{current.Val,nil}
            nhead = nhead.Next
        }
        current = current.Next
    }
    
    current = head
    for current != nil{
        if current.Val >= x{
            nhead.Next = &ListNode{current.Val,nil}
            nhead = nhead.Next
        }
        current = current.Next
    }
    return partition.Next
}
