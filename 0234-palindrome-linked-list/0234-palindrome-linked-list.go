/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var reverse *ListNode
    
    current := head
    
    for current != nil{            
        newNode := &ListNode{Val:current.Val}
        newNode.Next = reverse
        reverse = newNode
        current = current.Next        
    }
    
    for head != nil{
        if head.Val != reverse.Val{
           return false 
        }
        reverse = reverse.Next
        head = head.Next
    }
    return true
}