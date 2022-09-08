/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    stack := []*ListNode{}
    
    current := head
    for current != nil{
        stack = append(stack,current)
        current = current.Next
    }
    
    pop_count := 1
    if len(stack) == 1{
        return nil
    }
    if n == 1 {
        stack[len(stack)-2].Next = nil   
        return head
    }
    if n == len(stack){
        head = head.Next
        return head
    }
    current = stack[len(stack)-1]
    for pop_count < n {
        pop_count++
        stack = stack[:len(stack)-1]
        current = stack[len(stack)-1]
    }
    stack[len(stack)-2].Next = current.Next
    return head
    
}