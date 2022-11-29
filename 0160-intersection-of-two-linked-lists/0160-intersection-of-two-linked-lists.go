/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    nodes := make(map[*ListNode]int)    
    currentA := headA
    var i int
    var res *ListNode  
    ires := int(^uint(0)>>1)
    for currentA != nil{
        nodes[currentA] = i
        i++
        currentA = currentA.Next
    }
    currentB := headB
    for currentB != nil{
        if node,e := nodes[currentB];e && node < ires{
            ires = node
            res = currentB
        }
        currentB = currentB.Next
    }
    
    return res
}