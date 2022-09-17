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
            currentReversed.Next = reverseList(l,r)
            currentReversed = l 
            r = l
            l = l.Next
        }
        r = r.Next
        current++
    }
    
    
    return newHead.Next
}
// O(k) 
func reverseList(a,b *ListNode) *ListNode{
    newHead := b.Next
    l,r := a, b.Next
    
    for l != newHead {
        temp := l.Next
        l.Next = r
        r= l 
        l = temp                 
    }    
    return r
}