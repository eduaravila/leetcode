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
    postInf := int(^uint(0)>>1)
    
    res := &ListNode{Val: -postInf}
    
    for _, list := range lists{        
        res = mergeList(res,list)
    }
    
    return res.Next
}