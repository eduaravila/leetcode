/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getHalf(head *ListNode) *ListNode{
    if head == nil{
        return head
    }
    slow,fast := head, head
    for fast != nil && fast.Next != nil {        
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    return slow
}

func merge(a,b *ListNode)*ListNode{
    if a == nil {
        return b
    }
    if b == nil {
        return a
    }
    
    if a.Val < b.Val {
        a.Next = merge(a.Next,b)
        return a
    }else{
        b.Next = merge(a,b.Next)
        return b
    }
    
    return nil
    
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }    
    mid := getHalf(head)    
    temp := mid.Next
    if head.Next == mid {
        head.Next = nil
        temp = mid        
    }else{
        mid.Next = nil    
    }
    l:=sortList(head)
    r:=sortList(temp)    
    res := merge(l,r)
    return res
}

