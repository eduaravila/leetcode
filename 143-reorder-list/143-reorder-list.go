/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    stack := []*ListNode{}
    
    r := head
    for r != nil{
        stack = append(stack, r)
        r= r.Next
    }

    r = stack[len(stack)-1]
    for head != r && r != head.Next{
        r.Next = head.Next
        head.Next = r        
        stack = stack[:len(stack)-1]
        head = r.Next
        r = stack[len(stack)-1]            
    }
    r.Next= nil
}