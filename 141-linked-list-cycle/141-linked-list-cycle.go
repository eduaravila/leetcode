/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    f,ff := head,head.Next
    
    for f != nil && ff != nil{
        
        if f == ff {
            return true
        }
        
        f= f.Next
        if ff.Next == nil{
            break            
        }
        ff=ff.Next.Next
    }  
    return false
}