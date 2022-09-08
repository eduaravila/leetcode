/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    temp_head := &ListNode{Val:0,Next:head}
    r := head
    for n > 0{
        r = r.Next
        n--
    }
    
    l:= temp_head
    for r != nil{
        r = r.Next
        l = l.Next
    }
    
    l.Next = l.Next.Next
    return temp_head.Next   
}