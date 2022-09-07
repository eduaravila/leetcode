/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    
    current := head
    fast := current.Next
    faster := fast.Next
    for fast != nil{
        fast.Next = current
        current = fast                
        fast = faster
        
        if faster == nil{
            break
        }
        faster = faster.Next        
    }
    head.Next = nil
    return current
}