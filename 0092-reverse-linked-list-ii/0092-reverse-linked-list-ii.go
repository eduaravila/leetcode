/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(head *ListNode) *ListNode{
    if head == nil{
        return head
    }
    var prev *ListNode
    current := head
    for current!= nil{
        temp := current.Next
        current.Next = prev
        prev = current
        current = temp
    }
    return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{}
    lprev, current := dummy,head
    lprev.Next = head
    for i := 0 ; i < left-1;i++{
        if i < left {
            lprev,current = current, current.Next
        }
    }
    var prev *ListNode
    temp := current.Next
    for i := 0 ; i<  (right - left) + 1 ;i++{
        temp = current.Next
        current.Next = prev
        prev = current
        current = temp
    }
    lprev.Next.Next= current
    lprev.Next = prev
    return dummy.Next
}