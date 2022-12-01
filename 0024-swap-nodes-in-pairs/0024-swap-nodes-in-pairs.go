/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    return solution(head,0)
}


func solution(head *ListNode, index int) *ListNode{
    if head == nil || head.Next == nil {
        return head
    }
    prev := solution(head.Next,index+1)
    
    if index %2 == 0 {
        head.Next = prev.Next    
        prev.Next= head
        return prev
    }else{
        head.Next = prev        
    }
    return head
}
