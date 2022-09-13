/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    current := head
    var sum, carry,digit int
    for l1 != nil || l2 != nil {
        sum = carry
        if l1 != nil{
            sum += l1.Val
            l1 = l1.Next
        
        }
        if l2 != nil{
            sum += l2.Val
            l2 = l2.Next
        }
        
        digit = sum % 10 
        carry = sum / 10                 
        current.Next = &ListNode{ Val: digit }
        current = current.Next
        
    }
    if carry > 0 {
        current.Next = &ListNode{ Val:carry }
    }
    
    return head.Next
}