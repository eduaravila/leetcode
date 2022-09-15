/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeList(a,b *ListNode) *ListNode{
    merged_head := &ListNode{}
    current := merged_head    
    for a != nil && b != nil {
        if a.Val < b.Val {
            current.Next = a
            a = a.Next
        }else{        
            current.Next = b
            b = b.Next
        }    
        current = current.Next
    }
    if a != nil{
        current.Next = a
    }
    if b != nil {
        current.Next = b
    }
    return merged_head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) < 1{
        return nil
    }
    res := lists[0]    
    for _, list := range lists[1:]{        
        res = mergeList(res,list)
    }
    
    return res
}